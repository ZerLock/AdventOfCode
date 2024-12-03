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

func CalculateSimilarityScore(list1 []int, list2 []int) int {
	score := 0

	withoutDup := utils.RemoveDuplicates[int](list1)
	for _, elem := range withoutDup {
		score += elem * utils.CountElem[int](list2, elem)
	}
	return score
}

func main() {
	fileString := utils.ReadFileToString("./input.txt")

	list1, list2 := TransformToLists(fileString)

	sort.Ints(list1)
	sort.Ints(list2)

	similarity := CalculateSimilarityScore(list1, list2)

	fmt.Println(similarity)

}
