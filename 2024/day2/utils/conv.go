package utils

import (
	"strconv"
	"strings"
)

func StringToInt(str string) int {
	n, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return n
}

func StringToIntArray(str string, delim string) []int {
	arr := strings.Split(str, delim)
	res := make([]int, len(arr))

	for index, elem := range arr {
		res[index] = StringToInt(elem)
	}
	return res
}
