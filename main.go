package main

import (
	"fmt"

	"go-leboncoin/leboncoin"
)

func main() {
	// Set HTTP Client
	lbc := leboncoin.New()

	search := &leboncoin.Search{
		Limit: 100,
		Filters: leboncoin.Filters{
			Category: leboncoin.CategoryFilter{
				ID: "3",
			},
			Location: leboncoin.LocationFilter{
				Department: "31",
			},
			Keywords: leboncoin.KeywordsFilter{
				Text: "Honda CBF 600",
			},
		},
	}

	//
	resp, err := lbc.Search(search)
	if err != nil {

	}

	fmt.Println(resp.Total)
}
