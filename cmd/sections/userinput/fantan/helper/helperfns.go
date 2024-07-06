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
