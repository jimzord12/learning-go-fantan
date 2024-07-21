package tests

import "fmt"

func TextInfoWrapper(testType string, testFn func(...any), testFnArgs ...any) {
	fmt.Printf("[TESTING]: Begin of (%s) Testing", testType)
	testFn(testFnArgs...)
	fmt.Printf("[TESTING]: Begin of (%s) Testing", testType)
	fmt.Println()
}
