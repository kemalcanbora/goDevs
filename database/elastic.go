package elastic

import (
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	helpers "goDevs/helper"
	"os"
)

func esConnection() *elastic.Client {
	helpers.GetEnv()
	esHost := os.Getenv("DOCKER_ES_HOST")
	esUserName := os.Getenv("DOCKER_ES_USER_NAME")
	esPassword := os.Getenv("DOCKER_ES_PASSWORD")

	es, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(esHost),
		elastic.SetBasicAuth(esUserName, esPassword),
	)

	if err != nil {
		fmt.Println(err)
	}

	return es
}
