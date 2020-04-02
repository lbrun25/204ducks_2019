package ducks

import (
	"fmt"
)

// Increment - for computing integrals
var Increment float64 = 0.001

// Ducks - main function
func Ducks() {
	expectedValue := ComputeExpectedValue()

	fmt.Println("Average return time:", FormatMinutesToHuman(expectedValue))
	fmt.Printf("Standard deviation: %.3f\n", ComputeStandardDeviation(expectedValue))
	fmt.Println("Time after which 50% of the ducks are back:", SecondsToMinutes(ComputeTimeDucksBack(50.0)))
	fmt.Println("Time after which 99% of the ducks are back:", SecondsToMinutes(ComputeTimeDucksBack(99.0)))
	fmt.Printf("Percentage of ducks back after 1 minute: %.1f\n", ComputePercentageDucksBack(1.0))
	fmt.Printf("Percentage of ducks back after 2 minutes: %.1f\n", ComputePercentageDucksBack(2.0))
}