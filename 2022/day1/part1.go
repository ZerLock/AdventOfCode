package main;

import (
	"os"
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	elfs := strings.Split(fileString, "\n\n")
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
		if (total > result) {
			result = total
		}
	}
	fmt.Println(result)
}