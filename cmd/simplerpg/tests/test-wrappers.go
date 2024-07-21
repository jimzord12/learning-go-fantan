package tests

import "fmt"

type testFunc func()

func TextInfoWrapper(test testFunc, testType string) {
	fmt.Printf("[TESTING]: Begin of (%s) Testing", testType)
	test()
	fmt.Printf("[TESTING]: Begin of (%s) Testing", testType)
	fmt.Println()
}
