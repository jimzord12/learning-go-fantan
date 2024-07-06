package fantan

import (
	"fmt"
	"hello-world/cmd/sections/userinput/fantan/helper"
	"os"
)

func PlayFantan() {
	fmt.Print("Running main.go\n\n")
	var wins, losses int
	var profit float64
	var deposit int

	deposit = helper.GetAndParseUserInputInt("Plzzz Give me Money")
	fmt.Print("Thank you for your money! (", deposit, ")\n\n")

	// 1st Game
	result := FanTanGame(deposit)
	if result > 0 {
		wins++
	} else {
		losses++
	}
	profit += result
	deposit += int(result)

	fmt.Println("-----------------------------------------")
	fmt.Println("WINS: ", wins, " | LOSSES: ", losses)
	fmt.Println("-----------------------------------------")
	fmt.Println("Profit: (", profit, ")")
	fmt.Println("-----------------------------------------")
	fmt.Println("Balance: (", deposit, ")")
	fmt.Println("-----------------------------------------")

	if deposit <= 0 {
		fmt.Println("Thank you for giving us your Money!!!")
		fmt.Println("Have a good day and come again soon!")
		os.Exit(0)
	}

	fmt.Println("To Quit press 'q' and to Play Again press 'p'...")

	// Next Games
	for {
		fmt.Println("To Play Again press 'p/P' | To Quit press 'q/Q'")
		var userChoise string
		fmt.Scanln(&userChoise)
		if userChoise == "p" || userChoise == "P" {
			result = FanTanGame(deposit)
			if result > 0 {
				wins++
			} else {
				losses++
			}
			profit += result
			deposit += int(result)

			fmt.Println("-----------------------------------------")
			fmt.Println("WINS: ", wins, " | LOSSES: ", losses)
			fmt.Println("-----------------------------------------")
			fmt.Println("Profit: (", profit, ")")
			fmt.Println("-----------------------------------------")
			fmt.Println("Balance: (", deposit, ")")
			fmt.Println("-----------------------------------------")

			if deposit <= 0 {
				fmt.Println("Thank you for giving us your Money!!!")
				fmt.Println("Have a good day and come again soon!")
				os.Exit(0)
			}
		} else if userChoise == "q" || userChoise == "Q" {
			fmt.Println("Thank you for Playing!")
			os.Exit(0)
		} else {
			fmt.Println("Please enter either 'q' or 'y'...")
		}
	}
}
