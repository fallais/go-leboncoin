package main

import (
	"fmt"

	"go-leboncoin/leboncoin"
)

func main() {
	// Set HTTP Client
	lbc := leboncoin.New()

	price := leboncoin.Range{
		Max: 3500,
	}

	search := leboncoin.NewSearch()
	search.SetLimit(100)
	search.SetCategory(leboncoin.MotorcycleCategory)
	search.SetLocationWithDepartment("31")
	search.SetKeywords("Honda CBF 600")
	search.AddRange("price", price)

	// Search the ads
	resp, err := lbc.Search(search)
	if err != nil {

	}

	// Display the ads
	for _, ad := range resp.Ads {
		fmt.Println(ad.Subject, ad.Price)
	}
}
