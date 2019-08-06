package leboncoin

// Range is a range filter.
type Range string

const (
	// CubicCapacity is the cubic capacity of the motor.
	CubicCapacity Range = "cubic_capacity"

	// Price is the price of the item.
	Price Range = "price"

	// RegDate is the year of building.
	RegDate Range = "regdate"

	// MileAge is the amount of kilometers.
	MileAge Range = "mileage"
)
