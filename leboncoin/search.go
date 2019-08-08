package leboncoin

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Filters are the filters of the search.
type Filters struct {
	Ranges map[string]map[string]int
	Enums  map[string][]string
}

// Location is the location for the search.
type Location struct {
	Type       string
	ZipCodes   []int
	Department int
	Region     int
	Area       string
}

// Search is the searching structure.
type Search struct {
	categoryID int
	keywords   string
	location   Location
	filters    *Filters
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewSearch returns a new Search.
func NewSearch(categoryID int, keywords string, location Location, filters *Filters) (*Search, error) {
	// Check the category
	_, ok := Categories[categoryID]
	if !ok {
		return nil, fmt.Errorf("category does not exists")
	}

	search := &Search{
		categoryID: categoryID,
		keywords:   keywords,
		location:   location,
		filters:    filters,
	}

	return search, nil
}

// NewSearchFromURL returns a new Search from the given URL.
func NewSearchFromURL(u string) (*Search, error) {
	// Prepare the URL
	parsedURL, err := url.Parse(u)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}

	// Parse the parameters
	params, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the parameters : %s", err)
	}

	// Check the category (required)
	categoryStr, ok := params["category"]
	if !ok {
		return nil, fmt.Errorf("category is not the parameters")
	}
	if len(categoryStr) <= 0 {
		return nil, fmt.Errorf("category must not be empty")
	}
	// Parse the category
	category, err := strconv.Atoi(categoryStr[0])
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the category : %s", err)
	}

	// Check the keywords (not required)
	keywords := ""
	if len(params["text"]) > 0 || len(strings.TrimSpace(params["text"][0])) > 0 {
		keywords = params["text"][0]
	}

	// Location
	location := Location{
		Type:       "department",
		Department: 31,
	}

	// Create the search
	search, err := NewSearch(category, keywords, location, nil)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the search : %s", err)
	}

	return search, nil
}

//------------------------------------------------------------------------------
// Helpers
//------------------------------------------------------------------------------

func parseLocation(l string) (*Location, error) {
	var location *Location

	// Region
	r, err := regexp.Compile("r_(\\d{2})")
	if err != nil {
		return nil, fmt.Errorf("Error while compiling the regex : %s", err)
	}
	if r.MatchString(l) {
		reg := r.FindString(l)

		// Parse the string to int
		i, err := strconv.Atoi(reg)
		if err != nil {
			return nil, fmt.Errorf("Error while converting the region : %s", err)
		}

		location.Type = "region"
		location.Region = i
	}

	// Departement
	r, err = regexp.Compile("d_(\\d{2})")
	if err != nil {
		return nil, fmt.Errorf("Error while compiling the regex : %s", err)
	}
	if r.MatchString(l) {
		dep := r.FindString(l)

		// Parse the string to int
		i, err := strconv.Atoi(dep)
		if err != nil {
			return nil, fmt.Errorf("Error while converting the department : %s", err)
		}

		location.Type = "department"
		location.Department = i
	}

	return location, nil
}
