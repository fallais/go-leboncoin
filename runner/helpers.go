package runner

import (
	"fmt"
	"strconv"
	"strings"
)

func parseRange(v string) (map[string]int, error) {
	// Explode the value and check the length
	minAndMax := strings.Split(v, "-")
	if len(minAndMax) != 2 {
		return nil, fmt.Errorf("range format is incorrect")
	}

	rangeMap := make(map[string]int)

	// Process the min
	if minAndMax[0] != "min" {
		parsedMin, err := strconv.Atoi(minAndMax[0])
		if err != nil {
			return nil, fmt.Errorf("error while converting the string to int : %s", err)
		}

		rangeMap["min"] = parsedMin
	}

	// Process the max
	if minAndMax[1] != "max" {
		parsedMax, err := strconv.Atoi(minAndMax[1])
		if err != nil {
			return nil, fmt.Errorf("error while converting the string to int : %s", err)
		}

		rangeMap["max"] = parsedMax
	}

	return rangeMap, nil
}
