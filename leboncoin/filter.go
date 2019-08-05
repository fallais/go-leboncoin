package leboncoin

type Search struct {
	Limit   int     `json:"limit"`
	Filters Filters `json:"filters"`
}

type CategoryFilter struct {
	ID string `json:"id"`
}

type LocationFilter struct {
	Department string `json:"department"`
}

type KeywordsFilter struct {
	Text string `json:"text"`
}

type Filters struct {
	Category CategoryFilter `json:"category"`
	Location LocationFilter `json:"location"`
	Keywords KeywordsFilter `json:"keywords"`
}
