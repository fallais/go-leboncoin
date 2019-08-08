package runner

import (
	"io/ioutil"

	"go-leboncoin/leboncoin"

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

// Run ...
func Run(filename string) {
	var filters Filters

	// Read the file
	file, err := ioutil.ReadFile(filename)
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
		}).Infof("New ads have been found !")
	}
}
