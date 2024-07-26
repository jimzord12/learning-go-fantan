package inputhelpers

import (
	"fmt"
	"strconv"
)

func GetTerminalInput(msg string) string {
	fmt.Println(msg)
	var input string
	fmt.Scanln(&input)
	return input
}

func GetTerminalInputInt(msg string, passes func(int) bool) int {
	var input string
	var intInput int
	var err error

	for {
		fmt.Println(msg)
		fmt.Scanln(&input)
		intInput, err = strconv.Atoi(input)
		if err == nil && passes(intInput) {
			break
		}
		fmt.Println("Invalid Input, Please try again.")
	}

	return intInput
}
