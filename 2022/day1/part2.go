package main;

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	elfs := strings.Split(fileString, "\n\n")
	sortedElfs := make([]int, 0)
	result := 0

	for _, elf := range elfs {
		foods := strings.Split(elf, "\n")
		total := 0
		for _, food := range foods {
			res, err := strconv.Atoi(food)
			if err != nil {
				panic(err)
			}
			total += res
		}
		sortedElfs = append(sortedElfs, total)
	}
	sort.Ints(sortedElfs)
	for _, elf := range sortedElfs[len(sortedElfs)-3:] {
		result += elf
	}
	fmt.Println(result)
}