# Go-LeBonCoin

**go-leboncoin** is a Golang client for the **Le Bon Coin REST API**.

## Usage

```go
package main

import (
	"fmt"

	"go-leboncoin/leboncoin"
)

func main() {
	// Set HTTP Client
	lbc := leboncoin.New()

	search := leboncoin.NewSearch()
	search.SetLimit(100)
	search.SetCategory(leboncoin.Motorcycle)
	search.SetLocationWithDepartment("31")
	search.SetKeywords("Honda CBF 600")

	// Search the ads
	resp, err := lbc.Search(search)
	if err != nil {

	}

	// Display the ads
	for _, ad := range resp.Ads {
		fmt.Println(ad.Subject, ad.Price)
	}
}
```