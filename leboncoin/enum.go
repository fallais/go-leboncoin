package leboncoin

// Enum is an enumeration filter.
type Enum string

const (
	// MotoTypeEnum is the enumeration for the type of motorcycle.
	MotoTypeEnum Enum = "moto_type"

	// MotoBrandEnum is the enumeration for the brand of the moto.
	MotoBrandEnum Enum = "moto_brand"

	// FuelEnum is the enumeration for the fuel of the vehicule.
	FuelEnum Enum = "fuel"

	// GearboxEnum is the enumeration for the gearbox of the vehicule.
	GearboxEnum Enum = "gearbox"

	// BrandEnum is the enumeration for the brand of the vehicule.
	BrandEnum Enum = "brand"
)
