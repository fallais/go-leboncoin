package runner

import (
	"io/ioutil"
	"time"

	"go-leboncoin/leboncoin"
	"go-leboncoin/models"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Filter ...
type Filter struct {
	Name      string `json:"name" yaml:"name"`
	IsEnabled bool   `json:"is_enabled" yaml:"is_enabled"`
	URL       string `json:"url" yaml:"url"`
}

// Filters is the structure of the filters
type Filters struct {
	Filters []Filter `json:"filters" yaml:"filters"`
}

// Runner is a runner tool.
type Runner struct {
	filename string
	offset   time.Duration
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// New returns a new Runner.
func New(filename string) *Runner {
	return &Runner{
		filename: filename,
		offset:   5 * time.Minute,
	}
}

//------------------------------------------------------------------------------
// Function
//------------------------------------------------------------------------------

// Run runs the job.
func (r *Runner) Run() {
	var filters Filters

	// Read the file
	file, err := ioutil.ReadFile(r.filename)
	if err != nil {
		logrus.Errorf("Error while reading the filters file : %s", err)
		return
	}

	// Unmarshal the file
	err = yaml.Unmarshal(file, &filters)
	if err != nil {
		logrus.Errorf("Error while unmarshalling the configuration : %s", err)
		return
	}

	// Create a new client
	lbc := leboncoin.New()

	// Process the filters
	for _, filter := range filters.Filters {
		if !filter.IsEnabled {
			logrus.WithFields(logrus.Fields{
				"filter_name": filter.Name,
			}).Infoln("Filter is not enabled")
			continue
		}

		// Create the search
		search, err := leboncoin.NewSearchFromURL(filter.URL)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"filter_name": filter.Name,
			}).WithError(err).Errorln("Error while creating the search from the URL")
			continue
		}

		logrus.WithFields(logrus.Fields{
			"filter_name": filter.Name,
		}).Infoln("Searching the ads")

		// Search the ads
		response, err := lbc.Search(search)
		if err != nil {
			logrus.Errorf("Error while searching ads : %s", err)
			return
		}

		logrus.WithFields(logrus.Fields{
			"filter_name": filter.Name,
			"nb_ads":      len(response.Ads),
		}).Infoln("Successfully searched the ads")

		// Find the news ads (after the offset)
		var newAds []*models.Ad
		for _, lbcAd := range response.Ads {
			// Parse the date
			publishedDate, err := time.ParseInLocation("2006-01-02 15:04:05", lbcAd.FirstPublicationDate, time.Local)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"filter_name": filter.Name,
				}).WithError(err).Errorln("Error while parsing the date")
				continue
			}

			// Check the offset
			if !publishedDate.After(time.Now().Add(-r.offset)) {
				continue
			}

			// Sanitize the price
			price := 0
			if len(lbcAd.Price) > 0 {
				price = lbcAd.Price[0]
			}

			// Create the model
			ad := models.NewAd(lbcAd.Subject, lbcAd.Body, price, lbcAd.URL, publishedDate)

			// Append
			newAds = append(newAds, ad)
		}

		logrus.WithFields(logrus.Fields{
			"filter_name": filter.Name,
			"nb_ads":      len(newAds),
			"offset":      r.offset,
		}).Infoln("Displaying the new ads")
		for _, ad := range newAds {
			logrus.WithFields(logrus.Fields{
				"price": ad.Price,
				"date":  ad.PublishedDate,
			}).Infoln(ad.Name)
		}
	}
}
