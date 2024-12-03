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

func RemoveElemAtIndex[T comparable](list []T, index int) []T {
	new_arr := make([]T, len(list)-1)
	k := 0

	for i := 0; i < len(list); i++ {
		if i != index {
			new_arr[k] = list[i]
			k++
		}
	}
	return new_arr
}
