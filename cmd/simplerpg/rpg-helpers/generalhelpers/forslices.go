package generalhelpers

import (
	"golang.org/x/exp/constraints"
)

// Function to remove an element from a slice
func RemoveFromSlice[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if the index is out of bounds
	}
	return append(slice[:index], slice[index+1:]...)
}

func ExistsInSlice[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func GetAvgFromSlice[T constraints.Integer | constraints.Float](slice []T) float64 {
	if len(slice) == 0 {
		return 0
	}

	var sum T

	for _, v := range slice {
		sum += v
	}

	return float64(sum) / float64(len(slice))
}

// func FindIndex[T comparable](slice []T, identifier string) (int, error) {
// 	for i, v := range slice {
// 		if v == identifier {
// 			return i, nil
// 		}
// 	}

// 	return -1, errors.New("[ERROR]: Could not find Index")
// }
