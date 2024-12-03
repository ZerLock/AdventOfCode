package main

import (
	"day2/utils"
	"fmt"
	"sort"
	"strings"
)

func IsReportSafe(report []int) bool {
	isSorted := sort.SliceIsSorted(report, func(a, b int) bool { return report[a] < report[b] })
	isReverseSorted := sort.SliceIsSorted(report, func(a, b int) bool { return report[a] > report[b] })

	if isReverseSorted {
		for i, j := 0, len(report)-1; i < j; i, j = i+1, j-1 {
			report[i], report[j] = report[j], report[i]
		}
	}

	isSafeTransition := true
	for i := 1; i < len(report); i++ {
		if report[i]-report[i-1] > 3 || report[i]-report[i-1] == 0 {
			isSafeTransition = false
		}
	}

	return (isSorted || isReverseSorted) && isSafeTransition
}

func main() {
	fileString := utils.ReadFileToString("./input.txt")
	rawReports := strings.Split(fileString, "\n")

	reports := make([][]int, len(rawReports))
	for index, rawReport := range rawReports {
		reports[index] = utils.StringToIntArray(rawReport, " ")
	}

	nbSafeReports := 0
	for _, report := range reports {
		if IsReportSafe(report) == true {
			nbSafeReports++
		}
	}

	fmt.Println(nbSafeReports)
}
