package ducks

import (
	"continuous"
	"fmt"
	"math"
)

func formatMinutesToHuman(minutes float64) string {
	var res string

	hh := continuous.ComputeExpectedValue(minutes) / 60
	_, hhRHS := math.Modf(hh);
	mm := hhRHS * 60
	mmLHS, mmRHS := math.Modf(mm);
	res += fmt.Sprintf("%.0f", mmLHS)
	res += "m "
	
	ss := mmRHS * 60
	ssLHS, _ := math.Modf(ss);
	res += fmt.Sprintf("%.0f", ssLHS + 1)
	res += "s"

	return res
}

// Ducks - main function
func Ducks() {
	fmt.Println("Average return time:", formatMinutesToHuman(A))
	fmt.Println("Standard deviation:", continuous.ComputeStandardDeviation())
	fmt.Println("Time after which 50% of the ducks are back:", continuous.ComputeTimeDucksBack(50))
	fmt.Println("Time after which 99% of the ducks are back:", continuous.ComputeTimeDucksBack(99))
	fmt.Println("Percentage of ducks back after 1 minute:", continuous.ComputePercentageDucksBack(1))
	fmt.Println("Percentage of ducks back after 2 minutes:", continuous.ComputePercentageDucksBack(2))
}