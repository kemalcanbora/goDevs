package helper


type Sheet struct {
	Version string `json:"version"`
	ReqID   string `json:"reqId"`
	Status  string `json:"status"`
	Sig     string `json:"sig"`
	Table   struct {
		Cols []struct {
			ID      string `json:"id"`
			Label   string `json:"label"`
			Type    string `json:"type"`
			Pattern string `json:"pattern,omitempty"`
		} `json:"cols"`
		Rows []struct {
			C []interface{} `json:"c"`
		} `json:"rows"`
		ParsedNumHeaders int `json:"parsedNumHeaders"`
	} `json:"table"`
}

type Person struct {
	Name        string `json:"name"`
	Company     string `json:"company"`
	SocialMedia string `json:"social_media"`
}