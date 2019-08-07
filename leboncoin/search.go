package leboncoin

import (
	"fmt"
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
func (s *Search) AddRange(name string, value map[string]int) {
	s.Filters.Ranges[name] = value
}

// AddEnum adds an enumeration filter.
func (s *Search) AddEnum(name string, value string) {
	var enum []string
	enum = append(enum, value)

	s.Filters.Enums[name] = enum
}
