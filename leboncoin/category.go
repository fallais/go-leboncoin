package leboncoin

// Category is the product category.
type Category int

// Categories
const (
	CarCategory Category = iota + 2
	MotorcycleCategory
	CaravaningCategory
	CommercialVehicleCategory
	TrucksCategory
)

// Name returns the name of category.
func (c Category) Name() string {
	categories := []string{"", "", "voitures", "motos", "utilitaires", "camions"}
	return categories[c]
}
