package inputhelpers

import (
	"fmt"
)

func GetTerminalInput(msg string) string {
	fmt.Println(msg)
	var input string
	fmt.Scanln(&input)
	return input
}
