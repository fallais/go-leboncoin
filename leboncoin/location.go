package leboncoin

import (
	"fmt"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Location is the location for the search.
type Location struct {
	t          string
	zipCodes   []int
	department int
	region     int
	area       string
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewDepartmentLocation returns a new Location from the given department.
func NewDepartmentLocation(department int) Location {
	return Location{
		t:          "department",
		department: department,
	}
}

// NewRegionLocation returns a new Location from the given region.
func NewRegionLocation(region int) Location {
	return Location{
		t:      "region",
		region: region,
	}
}

//------------------------------------------------------------------------------
// Function
//------------------------------------------------------------------------------

// String returns a string representation of the Location.
func (l *Location) String() string {
	switch l.t {
	case "department":
		return fmt.Sprintf("department: %d", l.department)
	case "region":
		return fmt.Sprintf("region: %d", l.region)
	default:
		return "unknown"
	}
}
