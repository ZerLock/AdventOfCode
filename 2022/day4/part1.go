package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInteger(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return res
}

func createElf(elf string) []int {
	result := make([]int, 0)
	arr := strings.Split(elf, "-")
	for _, section := range arr {
		result = append(result, toInteger(section))
	}
	return result
}

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	pairs := strings.Split(fileString, "\n")
	count := 0

	for _, pair := range pairs {
		elfs := strings.Split(pair, ",")

		firstElf := createElf(elfs[0])
		secondElf := createElf(elfs[1])

		if (firstElf[0] <= secondElf[0] && firstElf[1] >= secondElf[1]) || (secondElf[0] <= firstElf[0] && secondElf[1] >= firstElf[1]) {
			count++
		}
	}
	fmt.Println(count)
}
