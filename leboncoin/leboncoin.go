package leboncoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"go-leboncoin/models"
)

const baseURI = "https://api.leboncoin.fr/finder/search"

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
func (lbc *LeBonCoin) Search(search *Search) ([]*models.Ad, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(baseURI)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}

	// Create the data
	data, err := json.Marshal(search)
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
	var response *Response

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s", err)
	}

	// Create the ads
	var ads []*models.Ad
	for _, lbcAd := range response.Ads {
		// Parse the date
		publishedDate, err := time.Parse("2006-01-02 15:04:05", lbcAd.FirstPublicationDate)
		if err != nil {
			fmt.Println("Error")
			continue
		}

		// Sanitize the price
		price := 0
		if len(lbcAd.Price) > 0 {
			price = lbcAd.Price[0]
		}

		// Create the model
		ad := models.NewAd(lbcAd.Subject, lbcAd.Body, price, publishedDate)

		// Append
		ads = append(ads, ad)
	}

	return ads, nil
}
