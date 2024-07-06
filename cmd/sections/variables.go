package sections

import "fmt"

// Automatically inferred type
// var agency = "NASA" // package level scope, compiler inferred type
// agency2 := "NASA2" // At package level scope, the ":=" shorthand is NOT allowed

// Explicit type
// var mission string = "Cassini" // package level scope, explicit type
// var year int = 1997            // package level scope, explicit type

func Variables() {
	fmt.Print("1. Running variables.go\n\n")

	// Local scope
	var message1 string = "This is a" // local scope, explicit type
	message2 := "variable"            // local scope, compiler inferred type
	fmt.Println(message1, message2)

}
