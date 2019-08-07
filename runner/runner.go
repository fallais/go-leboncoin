package runner

import (
	//"fmt"
	"io/ioutil"
	"strconv"
	"strings"

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

	// Read configuration file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Errorf("Error while reading the filters file : %s", err)
		return
	}

	// Unmarshal
	err = yaml.Unmarshal(file, &filters)
	if err != nil {
		logrus.Errorf("Error while unmarshalling the configuration : %s", err)
		return
	}

	// Create a new client
	lbc := leboncoin.New()

	for _, filter := range filters.Filters {
		if !filter.IsEnabled {
			continue
		}

		logrus.WithFields(logrus.Fields{
			"filter_name": filter.Name,
		}).Infoln("Searching ads")

		// Create the search
		search := leboncoin.NewSearch()
		search.SetLimit(100)
		search.SetCategory(filter.Category)
		search.SetLocationWithDepartment(filter.Location)
		search.SetKeywords(filter.Keywords)

		// Ranges
		for k, v := range filter.Criterias.Ranges {
			rangeMap := make(map[string]int)

			// Explode the value
			minAndMax := strings.Split(v, "-")

			if len(minAndMax) != 2 {
				return
			}

			// Min
			if minAndMax[0] != "min" {
				parsedMin, err := strconv.Atoi(minAndMax[0])
				if err != nil {
					return
				}

				rangeMap["min"] = parsedMin
			}

			// Min
			if minAndMax[1] != "max" {
				parsedMax, err := strconv.Atoi(minAndMax[1])
				if err != nil {
					return
				}

				rangeMap["max"] = parsedMax
			}

			// Add the range
			search.AddRange(k, rangeMap)
		}

		// Enums
		/* 		for _, enum := range filter.Criterias.Enums {
		   		} */

		// Search the ads
		resp, err := lbc.Search(search)
		if err != nil {
			logrus.Errorf("Error while searching ads : %s", err)
			return
		}

		// Display the ads
		logrus.Infof("%d ads have been found !", len(resp.Ads))
		for _, ad := range resp.Ads {
			if len(ad.Price) > 0 {
				logrus.WithFields(logrus.Fields{
					"price": ad.Price[0],
				}).Infoln(ad.Subject)
			} else {
				logrus.Infof("%s : pas de prix\n", ad.Subject)
			}
		}
	}
}

/* search.SetKeywords("Honda CBF 600")
search.AddRange(leboncoin.PriceRange, map[string]int{"min": 4000, "max": 8000})
search.AddRange(leboncoin.CubicCapacityRange, map[string]int{"min": 500})
search.AddEnum(leboncoin.BrandEnum, "Ford")
search.AddEnum(leboncoin.FuelEnum, "1") */
