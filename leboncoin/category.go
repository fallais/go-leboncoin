package leboncoin

// Category is the product category.
type Category int

// Categories
const (
	Car Category = iota + 2
	Motorcycle
	Caravaning
	CommercialVehicle
	Trucks
)

// String returns the category as a string.
func (c Category) String() string {
	return string(c)
}

// Name returns the name of category.
func (c Category) Name() string {
	categories := []string{"", "", "voitures", "motos", "utilitaires", "camions"}
	return categories[c]
}