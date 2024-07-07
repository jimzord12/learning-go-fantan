package main

import (
	"fmt"
	"hello-world/cmd/sections/userinput/fantan"
	"hello-world/cmd/sections/userinput/fantan/helper"
)

func main() {
	fmt.Print("Running main.go\n\n")

	userBalance := helper.GetAndParseUserInputInt("How much money do you wish to DEPOSIT?")
	fantan.PlayFantan(float64(userBalance))
}
