package env

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

// The ENV variables should not be set inside the program
// They can be set using the OS's terminal or by using .env file
// When using the .env file approach, the godotenv 3rd pt lib must be used to load them to the program's env
// But to access(get) them, we need to use the os package.
func Main() {
	// Print the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	fmt.Printf("Current working directory: %s\n", cwd)

	// Get the directory of the executable
	execDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalf("Error getting executable directory: %v", err)
	}
	fmt.Println(execDir)

	// Determine the path to the .env file
	envPath := filepath.Join(execDir, ".env")

	// // Get the directory of the executable
	// execPath, err := os.Executable()
	// if err != nil {
	// 	log.Fatalf("Error getting executable path: %v", err)
	// }
	// execDir := filepath.Dir(execPath)
	// fmt.Printf("Executable path: %s\n", execPath)
	// fmt.Printf("Executable directory: %s\n", execDir)

	// Load the .env file
	err = godotenv.Load(envPath)
	// err = godotenv.Load(filepath.Join(cwd, "bin", ".env"))
	// err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file\n%s", err)
	}

	// Get the port number from the environment variable
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080 // default port if not set or invalid
	}

	// Get the environment (development, production, etc.)
	env := os.Getenv("ENV")
	if env == "" {
		env = "development" // default environment
	}

	fmt.Printf("Starting server on port %d in %s mode\n", port, env)
}
