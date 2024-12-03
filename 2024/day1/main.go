package main

import (
	"fmt"
	"sort"
	"strings"
	"utils"
)

func ParseLists(data string) ([]int, []int) {
	lines := strings.Split(data, "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))

	for index, line := range lines {
		values := strings.Split(line, "   ")
		left[index] = utils.StringToInt(values[0])
		right[index] = utils.StringToInt(values[1])
	}
	return left, right
}

func Part1() int {
	fileContent := utils.ReadInput("./input.txt")

	left, right := ParseLists(fileContent)

	sort.Ints(left)
	sort.Ints(right)

	distance := 0
	for i := 0; i < len(left); i++ {
		tmpDist := left[i] - right[i]
		if tmpDist < 0 {
			distance += -tmpDist
		} else {
			distance += tmpDist
		}
	}
	return distance
}

func Part2() int {
	fileContent := utils.ReadInput("./input.txt")

	left, right := ParseLists(fileContent)

	sort.Ints(left)
	sort.Ints(right)

	similarityScore := 0
	leftWithoutDuplicates := utils.FilterDuplicates(left)
	for _, e := range leftWithoutDuplicates {
		similarityScore += e * utils.Count(right, e)
	}
	return similarityScore
}

func main() {
	fmt.Println("Day 01:\n-------")
	fmt.Println("Part 01:", Part1())
	fmt.Println("Part 02:", Part2())
	fmt.Println("-------")
}
