package alavara

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Client ...
type Client struct {
	Key string
}

// Rate ...
type Rate struct {
	Rate float64 `json:"rate"`
	Name string  `json:"name"`
	Type string  `json:"type"`
}

// Response ...
type Response struct {
	TotalRate float64 `json:"totalRate"`
	Rates     []Rate  `json:"rates"`
}

// RequestTaxRate ...
func (c Client) RequestTaxRate(zipCode string) (float64, error) {
	url := "taxrates.api.avalara.com"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	q := req.URL.Query()
	q.Add("country", "USA")
	q.Add("postal", zipCode)
	req.URL.RawQuery = q.Encode()
	//set auth header
	req.Header.Set("Authorization", "AvalaraApiKey "+c.Key)
	// create http client and send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	response := new(Response)
	// check if body missing
	if resp.Body == nil {
		return 0, errors.New("missing response body")
	}
	// throw error if json not decoded improperly
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}
	return response.TotalRate, nil

}
