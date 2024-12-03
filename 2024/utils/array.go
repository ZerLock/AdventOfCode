package utils

func Count[T comparable](list []T, value T) int {
	count := 0

	for _, e := range list {
		if e == value {
			count++
		}
	}
	return count
}

func Contains[T comparable](list []T, value T) bool {
	for _, elem := range list {
		if elem == value {
			return true
		}
	}
	return false
}

func FilterDuplicates[T comparable](list []T) []T {
	new_arr := make([]T, 0)

	for _, e := range list {
		if Contains[T](new_arr, e) {
			continue
		}

		new_arr = append(new_arr, e)
	}
	return new_arr
}

func RemoveAtIndex[T comparable](list []T, index int) []T {
	new_arr := make([]T, (len(list) - 1))
	k := 0

	for i := 0; i < len(list); i++ {
		if i != index {
			new_arr[k] = list[i]
			k++
		}
	}
	return new_arr
}

func Reverse[T comparable](list []T) {
	for i, j := 0, (len(list) - 1); i < j; i, j = (i + 1), (j - 1) {
		list[i], list[j] = list[j], list[i]
	}
}
