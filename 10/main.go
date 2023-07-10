package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		// Here we use math.Floor/10 and * 10
		// For example at first with math.Floor/10 we get "-25.4/10=-2.54" and then the func math.Floor returns the int less than -2.54 so its -2,
		// then we multiply it by 10 and get -20...
		groupKey := int(math.Floor(temp)/10) * 10 // Calculate the group key

		groups[groupKey] = append(groups[groupKey], temp)
	}

	// Print the temperature groups with a step 10
	for groupKey, val := range groups {
		fmt.Printf("%d: %v\n", groupKey, val)
	}
}
