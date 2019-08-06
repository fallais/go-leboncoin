# Go-LeBonCoin

![Coop](https://github.com/fallais/go-leboncoin/blob/master/gopher.png)

**go-leboncoin** is a Golang client for the **Le Bon Coin REST API**.

## Usage

```go
package main

import (
	"fmt"

	"github.com/fallais/go-leboncoin/leboncoin"
)

func main() {
	// Create a new client
	lbc := leboncoin.New()

	// Create the search
	search := leboncoin.NewSearch()
	search.SetLimit(100)
	search.SetCategory(leboncoin.MotorcycleCategory)
	search.SetLocationWithDepartment("31")
	search.SetKeywords("Honda CBF 600")
	search.AddRange(leboncoin.PriceRange, map[string]int{"min": 1500, "max": 3000})
	search.AddRange(leboncoin.CubicCapacityRange, map[string]int{"min": 500})
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

```