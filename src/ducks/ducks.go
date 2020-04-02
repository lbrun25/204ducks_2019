package ducks

import (
	"fmt"
)

// Increment - for computing integrals
var Increment float64 = 0.001

// Ducks - main function
func Ducks() {
	expectedValue := ComputeExpectedValue(A)

	fmt.Println("Average return time:", FormatMinutesToHuman(expectedValue))
	fmt.Printf("Standard deviation: %.3f\n", ComputeStandardDeviation(A, expectedValue))
	fmt.Println("Time after which 50% of the ducks are back:", ComputeTimeDucksBack(50))
	fmt.Println("Time after which 99% of the ducks are back:", ComputeTimeDucksBack(99))
	fmt.Printf("Percentage of ducks back after 1 minute: %.1f\n", ComputePercentageDucksBack(1))
	fmt.Printf("Percentage of ducks back after 2 minutes: %.1f\n", ComputePercentageDucksBack(2))
}