package utils

func CountElem[T comparable](list []T, value T) int {
	res := 0

	for _, elem := range list {
		if elem == value {
			res++
		}
	}
	return res
}

func Contains[T comparable](list []T, value T) bool {
	for _, elem := range list {
		if elem == value {
			return true
		}
	}
	return false
}

func RemoveDuplicates[T comparable](list []T) []T {
	result := make([]T, 0)

	for _, elem := range list {
		if Contains[T](result, elem) == true {
			continue
		}

		result = append(result, elem)
	}
	return result
}
