package utils

func ArrayContains[T comparable](array *[]T, objToFind T) bool {
	for _, item := range *array {
		if item == objToFind {
			return true
		}
	}

	return false
}
