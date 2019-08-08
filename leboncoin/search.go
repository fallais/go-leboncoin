package leboncoin

import (
	"fmt"
	"net/url"
	"strconv"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// ZipCode ...
type ZipCode struct {
	ZipCode string `json:"zipcode"`
}

// CategoryFilter ...
type CategoryFilter struct {
	ID string `json:"id"`
}

// KeywordsFilter ...
type KeywordsFilter struct {
	Text string `json:"text"`
}

// Filters holds the filters for the search.
type Filters struct {
	Category *CategoryFilter           `json:"category"`
	Location map[string]interface{}    `json:"location"`
	Keywords *KeywordsFilter           `json:"keywords"`
	Ranges   map[string]map[string]int `json:"ranges"`
	Enums    map[string][]string       `json:"enums"`
}

// Search is used to search.
type Search struct {
	Limit   int      `json:"limit"`
	Filters *Filters `json:"filters"`
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewSearch returns a new Search
func NewSearch() *Search {
	filter := &Filters{
		Category: nil,
		Location: nil,
		Keywords: nil,
		Ranges:   make(map[string]map[string]int),
		Enums:    make(map[string][]string),
	}

	return &Search{
		Limit:   100,
		Filters: filter,
	}
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

	// Create the search
	search := NewSearch()

	// Set the limit
	search.SetLimit(100)

	// Process the parameters
	for k, v := range params {
		switch k {
		case "category":
			// Parse the string to int
			i, err := strconv.Atoi(v[0])
			if err != nil {
				return nil, fmt.Errorf("Error while parsing the category : %s", err)
			}

			// Set the category
			search.SetCategory(i)
			break
		case "locations":
			//
			break
		default:
			// Check in Enums
			_, ok := Enums[k]
			if ok {
				//

				continue
			}

			// Check in Range
			if contains(Ranges, k) {
				//

				continue
			}
		}
	}

	return search, nil
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// SetLimit sets the limit.
func (s *Search) SetLimit(limit int) {
	s.Limit = limit
}

// SetCategory sets the category.
func (s *Search) SetCategory(categoryID int) error {
	_, ok := Categories[categoryID]
	if !ok {
		return fmt.Errorf("category does not exists")
	}

	s.Filters.Category = &CategoryFilter{
		ID: fmt.Sprint(categoryID),
	}

	return nil
}

// SetKeywords sets the keywords.
func (s *Search) SetKeywords(keywords string) {
	s.Filters.Keywords = &KeywordsFilter{
		Text: keywords,
	}
}

// SetLocationWithDepartment sets the location with department number.
func (s *Search) SetLocationWithDepartment(department int) {
	location := make(map[string]interface{})
	location["department"] = fmt.Sprint(department)
	s.Filters.Location = location
}

// SetLocationWithZipcodes sets the location with zipcodes.
func (s *Search) SetLocationWithZipcodes(zipcodes []ZipCode) {
	location := make(map[string]interface{})
	location["zipcodes"] = zipcodes
	s.Filters.Location = location
}

// AddRange adds a range filter.
func (s *Search) AddRange(name string, value map[string]int) error {
	if !contains(Ranges, name) {
		return fmt.Errorf("range does not exists")
	}

	s.Filters.Ranges[name] = value

	return nil
}

// AddEnum adds an enumeration filter.
func (s *Search) AddEnum(name string, value string) error {
	_, ok := Enums[name]
	if !ok {
		return fmt.Errorf("enum does not exists")
	}

	var enum []string
	enum = append(enum, value)

	s.Filters.Enums[name] = enum

	return nil
}
