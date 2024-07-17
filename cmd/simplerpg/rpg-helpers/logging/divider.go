package logging

import (
	"fmt"
	"strings"
)

func StdDivider(char string, amount int) {
	fmt.Println()
	fmt.Println(strings.Repeat(char, amount))
	fmt.Println()
}
