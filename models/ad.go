package models

import (
	"time"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Ad is the structure of an ad.
type Ad struct {
	Name          string
	Description   string
	Price         int
	URL           string
	PublishedDate time.Time
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewAd returns a new Ad.
func NewAd(name, description string, price int, url string, pubDate time.Time) *Ad {
	return &Ad{
		Name:          name,
		Description:   description,
		Price:         price,
		URL:           url,
		PublishedDate: pubDate,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------
