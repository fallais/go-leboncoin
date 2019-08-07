package leboncoin

// Category is the product category.
type Category int

const (
	// VehiculeCateggory is the main vehicule category (1).
	VehiculeCateggory Category = iota + 1
	// CarCategory for cars (2).
	CarCategory
	// MotorcycleCategory for motos (3).
	MotorcycleCategory
	// CaravaningCategory for caravans (4).
	CaravaningCategory
	// CommercialVehicleCategory for commercial vehicules (5).
	CommercialVehicleCategory
	// CarEquipmentCategory for cars equipements (6).
	CarEquipmentCategory
	// BoatCategory for boats and related (7).
	BoatCategory
	// RealEstateCategory is the main real estate category (8).
	RealEstateCategory
	// RealEstateSaleCategory for real estate sales (9).
	RealEstateSaleCategory
	// RealEstateRentCategory for real estate rent (9).
	RealEstateRentCategory
)

// Categories is the list of categories.
var Categories = map[int]string{
	1: "vehicules",
	2: "voitures",
	3: "motos",
}
