package ducks

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	tooManyArgs = "There are too many arguments.\n"
	notEnoughArgs = "There are not enough arguments.\n"
	mustBeInteger = "Argument(s) must be positive integers.\n"
	wrongAFormat = "Wrong real constant (A) format. The format must be a one decimal digit.\n"
	argMustBeBetween = "The real constant A must be between 0.0 and 2.5\n"
)

// A - Real constant provided as argument
var A float64 = 0.0

// CheckHelp arg -h
func CheckHelp() bool {
	argsWithoutProg := os.Args[1:]

	for _, arg := range argsWithoutProg {
		if (arg == "-h") {
			return true
		}
	}
	return false
}

// checkAFormat - check A (real constant) provided as argument with a regular expression
func checkAFormat(input string) bool {
	var re = regexp.MustCompile("[-]?[0-9]{1,9}\\.[0-9]{1,1}")

	match := re.FindString(input)
	if len(match) != len(input) {
		fmt.Println(wrongAFormat)
		return false
	}
	return true;
}

func checkArg(arg string) bool {
	if !checkAFormat(arg) {
		return false
	}
	resFloat, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		panic(err)
	}
	if (resFloat < 0 || resFloat > 2.5) {
		fmt.Println(argMustBeBetween)
		return false
	}
	A = resFloat
	return true
}

// CheckArgs check user input's args
func CheckArgs() bool {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println(notEnoughArgs)
		return false
	}
	if len(argsWithoutProg) > 1 {
		fmt.Println(tooManyArgs)
		return false
	}
	if !checkArg(argsWithoutProg[0]) {
		return false
	}
	return true
}