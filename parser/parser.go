package parser

import (
	"encoding/json"
	"fmt"
	helper "goDevs/helper"
	kafka "goDevs/streaming"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)


func Parser(kafkaTopic string) []interface{} {
	var s helper.Sheet
	var p helper.Person
	var data []interface{}


	resp, err := http.Get(fmt.Sprintf(helper.URL))
	if err != nil {
		fmt.Println("Error: ",err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	from := strings.Index(string(respBody), "{")
	to := strings.LastIndex(string(respBody), "}") + 1
	result := respBody[from:to]

	if jsErr := json.Unmarshal(result, &s); jsErr != nil {
		fmt.Println(jsErr)
		os.Exit(0)
	}

	for _, item := range s.Table.Rows {
		name, nameOk := item.C[1].(map[string]interface{})
		if nameOk {
			p.Name = name["v"].(string)
		} else {
			p.Name = "-"
		}

		company, companyOk := item.C[2].(map[string]interface{})
		if companyOk {
			p.Company = company["v"].(string)
		} else {
			p.Company = "-"
		}

		socialMedia, socialMediaOk := item.C[3].(map[string]interface{})
		if socialMediaOk {
			p.SocialMedia = socialMedia["v"].(string)
		} else {
			p.SocialMedia = "-"
		}

		pStructToBytes, _ := json.Marshal(p)

		kafka.Producer(kafkaTopic, pStructToBytes, []byte(p.Name))

		data = append(data,p)
	}

	return data
}
