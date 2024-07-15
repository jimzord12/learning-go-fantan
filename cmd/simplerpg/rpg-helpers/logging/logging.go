package logging

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var Logger *log.Logger

func LogInit() {
	Logger = log.New(os.Stdout, "[ERROR]: ", log.LstdFlags)
}

func LogError(logger *log.Logger, msg string) {
	// Retrieve the caller information
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		logger.Println("Failed to get caller information")
		return
	}

	// Get the function name
	funcName := runtime.FuncForPC(pc).Name()

	// Format the log message
	formattedMsg := fmt.Sprintf("[%s:%d %s] %s", file, line, funcName, msg)

	// Log the message
	logger.Println(formattedMsg)
}
