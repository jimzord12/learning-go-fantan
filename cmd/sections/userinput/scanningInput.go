package userinput

import "fmt"

func GettingInputFromUser(msgToDisplay string) {
	var userTextInput string
	fmt.Println(msgToDisplay)
	fmt.Scanln(&userTextInput)

	fmt.Println("\nYour input was: ", userTextInput)
}

func Main() {
	GettingInputFromUser("Tell me your name...")
}
