package main

import (
	"day1/utils"
	"fmt"
	"sort"
	"strings"
)

func TransformToLists(data string) ([]int, []int) {
	arr := strings.Split(data, "\n")
	list1 := make([]int, len(arr))
	list2 := make([]int, len(arr))

	for index, elem := range arr {
		tmp := strings.Split(elem, "   ")
		list1[index] = utils.StringToInt(tmp[0])
		list2[index] = utils.StringToInt(tmp[1])
	}
	return list1, list2
}

func main() {
	fileString := utils.ReadFileToString("./input.txt")

	list1, list2 := TransformToLists(fileString)

	sort.Ints(list1)
	sort.Ints(list2)

	distance := 0
	for i := 0; i < len(list1); i++ {
		tmp := list1[i] - list2[i]
		if tmp < 0 {
			distance += -tmp
		} else {
			distance += tmp
		}
	}
	fmt.Println(distance)
}
