package leboncoin

// Category is the product category.
type Category int

const (
	Car Category = iota + 2
	Motorcycle
	Caravaning
	Utilitaires
	Trucks
)

// String returns the cateogry as a string.
func (c Category) String() string {
	categories := []string{"", "", "voitures", "motos"}
	return categories[c]
}
