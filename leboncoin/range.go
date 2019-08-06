package leboncoin

// Range is a range filter with a minimum and a maximum.
type Range struct {
	Min int `json:"min,omitempty"`
	Max int `json:"max,omitempty"`
}

// CubicCapacity is the cubic capacity of the motor.
type CubicCapacity Range

// Price is the price of the item.
type Price Range

// RegDate ...
type RegDate Range

// MileAge is the amount of kilometers.
type MileAge Range
