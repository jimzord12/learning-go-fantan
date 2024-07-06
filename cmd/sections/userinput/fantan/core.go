package fantan

import (
	"fmt"
	"hello-world/cmd/sections/userinput/fantan/helper"
	"math/rand"
	"strconv"
	"time"
)

func FanTanGame(balance float64) float64 {
	const REWARD_MULTI = 2.85

	fmt.Println("Welcome to the Fan Tan Game!")
	fmt.Println("The rules are simple: ")
	fmt.Println("1. You have to guess the number of beans left in the table.")
	fmt.Println("2. The number of beans left in the table is between 1 and 4.")
	fmt.Println("3. You basically have 25% to guess correctly, can only guess once.")
	fmt.Print("4. If you guess correctly, you win the game.\n\n")

	// BET Management
	var bet int
	for {
		bet = helper.GetAndParseUserInputInt("Please Enter your BET:")
		if float64(bet) > balance {
			fmt.Println("[PROBLEM]: => Your Balance is insuffecient.")
			continue
		} else {
			break
		}
	}

	// Game Management
	var randFanTanNumber int32 = getRandInt(1, 4)
	for {
		var userGuess string
		if guess, err := getUserInputNum(&userGuess, "Please enter a number between 1 - 4"); guess == int(randFanTanNumber) && err == nil {
			fmt.Println("****************************************")
			fmt.Println("Congratulations! You guessed correctly!")
			fmt.Println("You Won => ", float64(bet)*REWARD_MULTI, " <=")
			fmt.Println("****************************************")

			return float64(bet) * REWARD_MULTI
		} else {
			if err != nil || guess < 1 || guess > 4 {
				fmt.Println("****************************************")
				fmt.Println("Invalid Selection.")
				fmt.Println("****************************************")

				continue
			} else {
				fmt.Println("*****************************************************************************")
				fmt.Println("Sadly, the final Beans were (", randFanTanNumber, "). Better Luck Next Time.")
				fmt.Println("You Lost => ", bet, " <=")
				fmt.Println("*****************************************************************************")

				return float64(bet) * (-1)
			}
		}
	}
}

// Internal Helper Functions
func getUserInputNum(userGuess *string, msg string) (int, error) {
	// var guess string
	fmt.Println(msg)
	fmt.Scanln(userGuess)
	fmt.Println("You selected the number: ", *userGuess)

	if convertedGuess, err := strconv.Atoi(*userGuess); err == nil {
		return convertedGuess, nil
	} else {
		return 0, err
	}
}

func getRandInt(min int32, max int32) int32 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// diff := max - min
	result := (r.Int31n(max) + 1)
	counter := 0
	for {
		if result < min {
			result = (r.Int31n(max) + 1)
		} else {
			break
		}
		counter++
	}
	// fmt.Println("Random Counter: ", counter)
	return result
}
