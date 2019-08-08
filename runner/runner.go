package runner

import (
	"io/ioutil"

	"go-leboncoin/leboncoin"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Criterias ...
type Criterias struct {
	Ranges map[string]string `json:"ranges" yaml:"ranges"`
	Enums  map[string]string `json:"enums" yaml:"enums"`
}

// Filter ...
type Filter struct {
	Name      string    `json:"name" yaml:"name"`
	IsEnabled bool      `json:"is_enabled" yaml:"is_enabled"`
	Keywords  string    `json:"keywords" yaml:"keywords"`
	Category  int       `json:"category" yaml:"category"`
	Location  int       `json:"location" yaml:"location"`
	Criterias Criterias `json:"criterias" yaml:"criterias"`
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

		logrus.WithFields(logrus.Fields{
			"filter_name": filter.Name,
		}).Infoln("Searching ads")

		// Create the location
		location := leboncoin.Location{
			Type:       "department",
			Department: 31,
		}

		// Create the filters
		filters := &leboncoin.Filters{}

		// Ranges
		ranges := make(map[string]map[string]int)
		for k, v := range filter.Criterias.Ranges {
			// Parse the range
			rangeMap, err := parseRange(v)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"filter_name": filter.Name,
					"range":       v,
				}).WithError(err).Errorln("Error while parsing the range")
				continue
			}

			// Add the range
			ranges[k] = rangeMap
		}
		filters.Ranges = ranges

		// Enums
		enums := make(map[string][]string)
		for k, v := range filter.Criterias.Enums {
			// Add the enums
			enums[k] = []string{v}
		}
		filters.Enums = enums

		// Create the search
		search, err := leboncoin.NewSearch(filter.Category, filter.Keywords, location, filters)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"filter_name": filter.Name,
			}).WithError(err).Errorln("Error while creating the search")
			continue
		}

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
