package leboncoin

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

type CategoryFilter struct {
	ID string `json:"id"`
}

type LocationFilter struct {
	Department string `json:"department"`
}

type KeywordsFilter struct {
	Text string `json:"text"`
}

// Filters holds the filters for the search.
type Filters struct {
	Category *CategoryFilter  `json:"category"`
	Location *LocationFilter  `json:"location"`
	Keywords *KeywordsFilter  `json:"keywords"`
	Ranges   map[string]Range `json:"ranges"`
	Enums    map[string]Enum  `json:"enum"`
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
	categoryFilter := &CategoryFilter{
		ID: category.String(),
	}

	filter

	s.CategoryFilter = categoryFilter
}
