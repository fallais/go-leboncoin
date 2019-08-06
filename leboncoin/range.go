package leboncoin

// Range is a range filter.
type Range string

const (
	// CubicCapacityRange is the cubic capacity of the motor.
	CubicCapacityRange Range = "cubic_capacity"

	// PriceRange is the price of the item.
	PriceRange Range = "price"

	// RegDateRange is the year of building.
	RegDateRange Range = "regdate"

	// MileAgeRange is the amount of kilometers.
	MileAgeRange Range = "mileage"
)
