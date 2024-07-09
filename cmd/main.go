package main

import (
	"fmt"
	"hello-world/cmd/fantan"
	"hello-world/cmd/fantan/helper"
	"hello-world/cmd/sections/myinterfaces"
)

func main() {
	// fantanWrapper()

	fmt.Print("Running main.go\n\n")

	// functions.Main()
	// structs.Main()
	// mytypes.Main()
	myinterfaces.Main()
	// whatever.Main()

}

func fantanWrapper() {
	userBalance := helper.GetAndParseUserInputInt("How much money do you wish to DEPOSIT?")
	fantan.PlayFantan(float64(userBalance))

}
