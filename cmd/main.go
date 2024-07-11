package main

import (
	"fmt"

	"github.com/jimzord12/learning-go-fantan/cmd/fantan"
	"github.com/jimzord12/learning-go-fantan/cmd/fantan/helper"
	"github.com/jimzord12/learning-go-fantan/cmd/sections/mytypes"
)

func main() {
	// fantanWrapper()

	fmt.Print("Running main.go\n\n")

	// functions.Main()
	// structs.Main()
	// mytypes.Main()
	mytypes.Main2() // Type Assertion
	// myinterfaces.Main()
	// whatever.Main()

}

func fantanWrapper() {
	userBalance := helper.GetAndParseUserInputInt("How much money do you wish to DEPOSIT?")
	fantan.PlayFantan(float64(userBalance))

}
