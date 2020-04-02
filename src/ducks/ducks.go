package ducks

import (
	"continuous"
	"fmt"
	//"math"
)

// Ducks - main function
func Ducks() {
	// ceil := int(continuous.ComputeExpectedValue(A))
	// sec := math.Ceil(continuous.ComputeExpectedValue(A) - float64(ceil) * 60)
	fmt.Println("Average return time: ", continuous.ComputeExpectedValue(A))
	fmt.Println("Standard deviation: ", continuous.ComputeStandardDeviation())
	fmt.Println("Time after which 50% of the ducks are back: ", continuous.ComputeTimeDucksBack(50))
	fmt.Println("Time after which 99% of the ducks are back: ", continuous.ComputeTimeDucksBack(99))
	fmt.Println("Percentage of ducks back after 1 minute: ", continuous.ComputePercentageDucksBack(1))
	fmt.Println("Percentage of ducks back after 2 minutes: ", continuous.ComputePercentageDucksBack(2))
}