package sections

import (
	"fmt"
	"strconv"
	"strings"
)

func Strings() {
	fmt.Print("2. Running strings.go\n\n")

	length := len("Hello, World!")
	fmt.Println("2.1 The String length is: ", length) // 13

	// String testing
	str1 := "Go Programming"
	str2 := "go programming"
	strings.EqualFold(str1, str2)            // true
	strings.Contains(str1, "Go")             // true
	strings.HasPrefix(str1, "Go")            // true
	strings.HasSuffix(str1, "ing")           // true
	strings.Index(str1, "o")                 // 1
	strings.LastIndex(str1, "m")             // 10
	strings.Replace(str1, "Go", "Golang", 1) // Golang Programming
	strings.Split(str1, " ")                 // [Go Programming]
	strings.ToLower(str1)                    // go programming
	strings.ToUpper(str1)                    // GO PROGRAMMING
	strings.TrimSpace("  Go Programming  ")  // Go Programming

	// Converting Number to String
	number := 10
	str := fmt.Sprintf("%d", number)
	// or
	str = fmt.Sprint(number)
	// or using strconv
	str = strconv.Itoa(number)

}
