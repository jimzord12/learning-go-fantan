package invpkg

import (
	"errors"
)

func (it Inventory) FindItemIndex(desiredItemName string) (int, error) {
	for i, v := range it.Items {
		if v.Name == desiredItemName {
			return i, nil
		}
	}

	return -1, errors.New("[ERROR]: Could not find Index")
}
