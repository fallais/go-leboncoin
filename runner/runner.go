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

		// Create the search
		search := leboncoin.NewSearch()

		// Set the limit
		search.SetLimit(100)

		// Set the category
		err := search.SetCategory(filter.Category)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"filter_name": filter.Name,
			}).WithError(err).Errorln("Error while setting the category")
			continue
		}

		// Set the location
		search.SetLocationWithDepartment(filter.Location)
		search.SetKeywords(filter.Keywords)

		// Ranges
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
			err = search.AddRange(k, rangeMap)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"filter_name": filter.Name,
					"range_name":  k,
				}).WithError(err).Errorln("Error while adding the range")
				continue
			}
		}

		// Enums
		for k, v := range filter.Criterias.Enums {
			// Add the enus
			err = search.AddEnum(k, v)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"filter_name": filter.Name,
					"enum_name":   k,
				}).WithError(err).Errorln("Error while adding the enum")
				continue
			}
		}

		// Search the ads
		ads, err := lbc.Search(search)
		if err != nil {
			logrus.Errorf("Error while searching ads : %s", err)
			return
		}

		// Display the ads
		logrus.WithFields(logrus.Fields{
			"filter_name": filter.Name,
			"nb_ads":      len(ads),
		}).Infof("New ads have been found !")
		for _, ad := range ads {
			logrus.WithFields(logrus.Fields{
				"price":          ad.Price,
				"published_date": ad.PublishedDate,
			}).Infoln(ad.Name)
		}
	}
}
