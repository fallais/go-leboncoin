package leboncoin

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// ErrCategoryNotExists is raised when the category does not exist.
var ErrCategoryNotExists = errors.New("category does not exist")

// ErrCategoryEmpty is raised when the category is empty.
var ErrCategoryEmpty = errors.New("category must not be empty")

// ErrDepartmentFormatIncorrect is raised when department format is incorrect.
var ErrDepartmentFormatIncorrect = errors.New("format of department is incorrect")

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Filters are the filters of the search.
type Filters struct {
	Ranges map[string]map[string]int
	Enums  map[string][]string
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
	// Check if the category exists
	if !categoryExists(categoryID) {
		return nil, ErrCategoryNotExists
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

	// Parse the category
	category, err := parseCategory(params)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the category : %s", err)
	}

	// Check the keywords (not required)
	keywords := parseKeywords(params)

	// Location
	location := NewDepartmentLocation(31)

	// Ranges
	ranges, err := parseRanges(params)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the ranges : %s", err)
	}
	// Enums
	enums, err := parseEnums(params)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the enums : %s", err)
	}

	// Filters
	filters := &Filters{
		Enums:  enums,
		Ranges: ranges,
	}

	// Create the search
	search, err := NewSearch(category, keywords, location, filters)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the search : %s", err)
	}

	return search, nil
}

//------------------------------------------------------------------------------
// Helpers
//------------------------------------------------------------------------------

// parseCategory parses the category from URL parameter.
func parseCategory(params url.Values) (int, error) {
	// Check the category (required)
	categoryStr, ok := params["category"]
	if !ok || len(categoryStr) != 1 {
		return 0, ErrCategoryEmpty
	}
	// Parse the category
	category, err := strconv.Atoi(categoryStr[0])
	if err != nil {
		return 0, fmt.Errorf("Error while parsing the category : %s", err)
	}
	// Check if the category exists
	if !categoryExists(category) {
		return 0, ErrCategoryNotExists
	}

	return category, nil
}

// parseKeywords parses the keywords from URL parameter.
func parseKeywords(params url.Values) string {
	if len(params["text"]) > 0 && len(strings.TrimSpace(params["text"][0])) > 0 {
		return params["text"][0]
	}

	return ""
}

// parseRanges parses the ranges from URL parameter.
func parseRanges(params url.Values) (map[string]map[string]int, error) {
	ranges := make(map[string]map[string]int)

	for k, v := range params {
		if contains(Ranges, k) {
			r, err := parseRange(v[0])
			if err != nil {
				return nil, fmt.Errorf("Error while parsing the range : %s", err)
			}

			ranges[k] = r
		}
	}

	return ranges, nil
}

// parseEnums parses the enums from URL parameter.
func parseEnums(params url.Values) (map[string][]string, error) {
	enums := make(map[string][]string)

	for k, v := range params {
		_, ok := Enums[k]
		if ok {
			enums[k] = v
		}
	}

	return enums, nil
}

// parseLocation parses the Location from URL parameter.
func parseLocation(l string) (*Location, error) {
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

		location := NewRegionLocation(i)
		return &location, nil
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

		location := NewDepartmentLocation(i)
		return &location, nil
	}

	return nil, fmt.Errorf("bad format")
}

func parseRange(v string) (map[string]int, error) {
	// Explode the value and check the length
	minAndMax := strings.Split(v, "-")
	if len(minAndMax) != 2 {
		return nil, fmt.Errorf("range format is incorrect")
	}

	rangeMap := make(map[string]int)

	// Process the min
	if minAndMax[0] != "min" {
		parsedMin, err := strconv.Atoi(minAndMax[0])
		if err != nil {
			return nil, fmt.Errorf("error while converting the string to int : %s", err)
		}

		rangeMap["min"] = parsedMin
	}

	// Process the max
	if minAndMax[1] != "max" {
		parsedMax, err := strconv.Atoi(minAndMax[1])
		if err != nil {
			return nil, fmt.Errorf("error while converting the string to int : %s", err)
		}

		rangeMap["max"] = parsedMax
	}

	return rangeMap, nil
}
