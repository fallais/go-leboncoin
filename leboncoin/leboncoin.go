package leboncoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"go-leboncoin/leboncoin/protocols"
)

// BaseURI is the base URI of the REST API.
const BaseURI = "https://api.leboncoin.fr/finder/search"

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// LeBonCoin is a client for LeBonCoin REST API.
type LeBonCoin struct {
	client *http.Client
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New returns a new client for LeBonCoin REST API.
func New() *LeBonCoin {
	client := &http.Client{}

	return &LeBonCoin{
		client: client,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Search ...
func (lbc *LeBonCoin) Search(search *Search) (*protocols.Response, error) {
	request := requestFromSearch(search)

	// Prepare the URL
	reqURL, err := url.Parse(BaseURI)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}

	// Create the data
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("Error while marshalling the values : %s", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Header
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := lbc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error while reading the request : %s", err)
	}

	// Prepare the response
	var response *protocols.Response

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s", err)
	}

	return response, nil
}

//------------------------------------------------------------------------------
// Helpers
//------------------------------------------------------------------------------

func requestFromSearch(s *Search) *protocols.Request {
	request := &protocols.Request{}

	location := make(map[string]interface{})
	location["department"] = fmt.Sprint(s.location.department)

	request.Limit = 100
	request.Filters = &protocols.Filters{
		Category: &protocols.CategoryFilter{
			ID: fmt.Sprint(s.categoryID),
		},
		Keywords: &protocols.KeywordsFilter{
			Text: s.keywords,
		},
		Location: location,
	}

	if s.filters != nil {
		request.Filters.Ranges = s.filters.Ranges
		request.Filters.Enums = s.filters.Enums
	}

	return request
}
