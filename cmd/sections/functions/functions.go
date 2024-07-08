package functions

import (
	"fmt"
)

func ReturningMultipleVars() (string, int, bool) {
	name := "Jimzord"
	age := 28
	isMan := true

	return name, age, isMan
}

func NamedReturnVars() (isNoob bool, totalWins int) {
	isNoob = true
	totalWins = -500

	return
}

func VariadicFunctions(numbers ...int) {
	var sum int
	for i, v := range numbers {
		fmt.Printf("[INDEX]: #%d\n", i)
		sum += v
		fmt.Printf("[SUM]: %v\n", sum)
		fmt.Println("----------------------")
	}
}

func Main() {

	fmt.Println("\n==== Named Return Vars ====")
	a, b := NamedReturnVars()
	fmt.Println("1. Using Named Retuen Vars: ", a)
	fmt.Println("2. Using Named Retuen Vars: ", b)

	fmt.Println("\n==== Error Handling ====")
	if answerMapping, err := FunctionWithErrorHandling(27); err != nil {
		// When you see [err != nil], err is NOT nil, so basically there is an Error
		// Mean the this block of code is HANDLING the ERROR
		fmt.Println("[error]: ", err)
	} else {
		// Here you can continue with the program's logic
		fmt.Println(answerMapping)
	}

	fmt.Println("\n==== Variadic Functions ====")
	VariadicFunctions(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("\n==== Anonymous Functions ====")
	anonymous := func(x int, y int) int {
		return x * y
	}
	fmt.Println("Anonymous Fn: ", anonymous(2, 3))

	fmt.Println("\n==== Closures ====")
	nextFibNum := FibonacciSequence(0, 1)
	fmt.Println("Fib #1:", nextFibNum())
	fmt.Println("Fib #2:", nextFibNum())
	fmt.Println("Fib #3:", nextFibNum())
	fmt.Println("Fib #4:", nextFibNum())
	fmt.Println("Fib #5:", nextFibNum())
	fmt.Println("Fib #6:", nextFibNum())
	fmt.Println("Fib #7:", nextFibNum())

}
