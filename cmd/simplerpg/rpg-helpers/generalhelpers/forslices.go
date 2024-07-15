package generalhelpers

// Function to remove an element from a slice
func RemoveFromSlice[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice // Return the original slice if the index is out of bounds
	}
	return append(slice[:index], slice[index+1:]...)
}

// func FindIndex[T comparable](slice []T, identifier string) (int, error) {
// 	for i, v := range slice {
// 		if v == identifier {
// 			return i, nil
// 		}
// 	}

// 	return -1, errors.New("[ERROR]: Could not find Index")
// }
