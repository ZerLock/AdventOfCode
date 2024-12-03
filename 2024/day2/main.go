package main

import (
	"fmt"
	"sort"
	"strings"
	"utils"
)

func StringReportToIntArray(str string) []int {
	str_arr := strings.Split(str, " ")
	new_arr := make([]int, len(str_arr))

	for index, e := range str_arr {
		new_arr[index] = utils.StringToInt(e)
	}
	return new_arr
}

func IsReportSafe(report []int) bool {
	isSorted := sort.SliceIsSorted(report, func(a, b int) bool { return report[a] < report[b] })
	isReversedSorted := sort.SliceIsSorted(report, func(a, b int) bool { return report[a] > report[b] })

	if isReversedSorted {
		utils.Reverse(report)
	}

	isSafeTransition := true
	for i := 1; i < len(report); i++ {
		distance := report[i] - report[i-1]

		if distance > 3 || distance <= 0 {
			isSafeTransition = false
		}
	}

	return (isSorted || isReversedSorted) && isSafeTransition
}

func Part1() int {
	fileContent := utils.ReadInput("./input.txt")
	lines := strings.Split(fileContent, "\n")

	reports := make([][]int, len(lines))
	for index, line := range lines {
		reports[index] = StringReportToIntArray(line)
	}

	nbSafeReports := 0
	for _, report := range reports {
		if IsReportSafe(report) {
			nbSafeReports++
		}
	}
	return nbSafeReports
}

func IsFalsyReportSafe(report []int) bool {
	for i := 0; i < len(report); i++ {
		falsyReport := utils.RemoveAtIndex(report, i)
		if IsReportSafe(falsyReport) {
			return true
		}
	}
	return false
}

func Part2() int {
	fileContent := utils.ReadInput("./input.txt")
	lines := strings.Split(fileContent, "\n")

	reports := make([][]int, len(lines))
	for index, line := range lines {
		reports[index] = StringReportToIntArray(line)
	}

	nbSafeReports := 0
	for _, report := range reports {
		if IsReportSafe(report) || IsFalsyReportSafe(report) {
			nbSafeReports++
		}
	}
	return nbSafeReports
}

func main() {
	fmt.Println("Day 02:\n-------")
	fmt.Println("Part 1:", Part1())
	fmt.Println("Part 2:", Part2())
	fmt.Println("-------")
}
