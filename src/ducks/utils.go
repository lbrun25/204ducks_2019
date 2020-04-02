package ducks

import (
	"math"
	"fmt"
)

// SecondsToMinutes - Format -> "2m 21s"
func SecondsToMinutes(inSeconds float64) string {
	var res string
	minutes := inSeconds / 60
	seconds := math.Mod(inSeconds, 60)

	res += fmt.Sprintf("%.0f", minutes) + "m "
	res += fmt.Sprintf("0%.0f", seconds) + "s"
	return res
}

// FormatMinutesToHuman - Format -> "2m 21s"
func FormatMinutesToHuman(minutes float64) string {
	var res string

	hh := minutes / 60
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