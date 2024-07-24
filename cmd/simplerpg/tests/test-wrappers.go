package tests

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/simplerpg/rpg-helpers/logging"
)

func TextInfoWrapper(testType string, testFn func(...any), testFnArgs ...any) {
	fmt.Println("")
	logging.StdDivider("#", 75)
	fmt.Println("")
	fmt.Printf("[TESTING]: Begin of (%s) Testing", testType)
	fmt.Println("")

	testFn(testFnArgs...)
	fmt.Println("")

	fmt.Printf("[TESTING]: End of (%s) Testing", testType)
	fmt.Println("")
	logging.StdDivider("#", 75)
	fmt.Println("")
}
