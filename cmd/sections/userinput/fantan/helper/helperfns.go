package helper

import (
	"fmt"
	"strconv"
)

func GetAndParseUserInputInt(msg string) int {
	var userInput string
	for {
		fmt.Println(msg)
		fmt.Scanln(&userInput)

		if convInput, err := strconv.Atoi(userInput); err == nil {
			return convInput
		} else {
			fmt.Println("Invalid input, try again")
			continue
		}
	}
}

func GetUserInput(msg string) string {
	var userInput string
	fmt.Println(msg)
	fmt.Scanln(&userInput)

	return userInput
}

func GetAndValidateInput(scannerFn func(msg string) string, scannerMsg string, args ...string) string {
	// p, P, q, Q
	isFirstInteration := true
	for {
		var userInputStr string
		if isFirstInteration {
			userInputStr = scannerFn(scannerMsg)
		} else {
			userInputStr = scannerFn(fmt.Sprint("Invalid Input, please try again.\n Available options: ", args))
		}
		for _, str := range args {
			if str == userInputStr {
				return userInputStr
			}
		}
		isFirstInteration = false
	}
}
