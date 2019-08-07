# Go-LeBonCoin

![Coop](https://github.com/fallais/go-leboncoin/blob/master/gopher.png)

**go-leboncoin** is a Golang client for the **Le Bon Coin REST API**.

**DISCLAIMER** : This is currently in **beta**, I mean the development is in progress. It is not yet released.

## Why ?

The aim is to create a bot that informs you of new ads based on your criterias. For example :

```yaml
filters:
  - name: moto
    is_enabled: true
    category: 3
    location: 31
    keywords: Honda CBF 600
    criterias:
      ranges:
        price: 1000-2500
        cubic_capacity: 500-600
      enums:
        moto_brand: honda
  - name: voiture
    is_enabled: false
    category: 2
    location: 31
    criterias:
      ranges:
        price: 4000-5000
      enums:
        brand: Ford
```

## Use as a library

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
	search.SetCategory(3)
	search.SetLocationWithDepartment(31)
	search.SetKeywords("Honda CBF 600")
	search.AddRange("price", map[string]int{"min": 1500, "max": 3000})
	search.AddRange("cubic_capacity", map[string]int{"min": 500})
	search.AddEnum("moto_brand", "honda")

	// Search the ads
	resp, _err_ := lbc.Search(search)

	// Display the ads
	fmt.Printf("%d ads have been found !\n", len(resp.Ads))
	for _, ad := range resp.Ads {
		fmt.Printf("%s : %d€\n", ad.Subject, ad.Price[0])
	}
}

```