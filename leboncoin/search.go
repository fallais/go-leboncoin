package leboncoin

import (
	"fmt"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type ZipCode struct {
	ZipCode string `json:"zipcode"`
}

type CategoryFilter struct {
	ID string `json:"id"`
}

type KeywordsFilter struct {
	Text string `json:"text"`
}

// Filters holds the filters for the search.
type Filters struct {
	Category *CategoryFilter        `json:"category"`
	Location map[string]interface{} `json:"location"`
	Keywords *KeywordsFilter        `json:"keywords"`
	Ranges   map[string]Range       `json:"ranges"`
	Enums    map[string]Enum        `json:"enum"`
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
		Ranges:   nil,
		Enums:    nil,
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
func (s *Search) SetCategory(category Category) {
	s.Filters.Category = &CategoryFilter{
		ID: fmt.Sprint(category),
	}
}

// SetKeywords sets the keywords.
func (s *Search) SetKeywords(keywords string) {
	s.Filters.Keywords = &KeywordsFilter{
		Text: keywords,
	}
}

// SetLocationWithDepartment sets the location with department number.
func (s *Search) SetLocationWithDepartment(department string) {
	location := make(map[string]interface{})
	location["department"] = department
	s.Filters.Location = location
}

// SetLocationWithZipcodes sets the location with zipcodes.
func (s *Search) SetLocationWithZipcodes(zipcodes []ZipCode) {
	location := make(map[string]interface{})
	location["zipcodes"] = zipcodes
	s.Filters.Location = location
}
