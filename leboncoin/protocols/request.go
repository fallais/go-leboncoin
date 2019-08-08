package protocols

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

// Request is used to request.
type Request struct {
	Limit   int      `json:"limit"`
	Filters *Filters `json:"filters"`
}
