package functions

import "errors"

// In Go, FUNCTIONS should NOT perform Error handling, but simply return an ERROR
// the CALLER should use this ERROR to perform ERROR HANDLING
func FunctionWithErrorHandling(age int) (map[string]bool, error) {
	if age == 28 {
		return map[string]bool{
			"correct": true,
		}, nil
	}

	return map[string]bool{
		"correct": true,
	}, errors.New("the error msg")
}
