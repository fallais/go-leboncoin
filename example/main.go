package main

import (
	"fmt"

	"go-leboncoin/leboncoin"
)

func main() {
	// Set HTTP Client
	lbc := leboncoin.New()

	// Create the search
	search := leboncoin.NewSearch()
	search.SetLimit(100)
	search.SetCategory(leboncoin.MotorcycleCategory)
	search.SetLocationWithDepartment("31")
	search.SetKeywords("Honda CBF 600")
	search.AddRange("price", map[string]int{"min": 1500, "max": 3000})
	search.AddEnum(leboncoin.MotoBrandEnum, "honda")

	// Search the ads
	resp, err := lbc.Search(search)
	if err != nil {

	}

	// Display the ads
	fmt.Printf("%d ads have been found !\n", len(resp.Ads))
	for _, ad := range resp.Ads {
		fmt.Printf("%s : %d€\n", ad.Subject, ad.Price[0])
	}
}
