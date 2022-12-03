package main

import (
	"fmt"
	"os"
	"strings"
)

func getPriority(item byte) int {
	if item > 'a' {
		return int(item - 96) // 'a' - 1 because count begin at 1
	}
	return int(item - 64 + 26) // 'a' - 1 because count begin at 1 + 26 which is the number of letter in alphabet
}

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	bags := strings.Split(fileString, "\n")
	priority := 0

	for _, bag := range bags {
		firstPart := bag[:len(bag)/2]
		secondPart := bag[len(bag)/2:]
		for _, item := range firstPart {
			if strings.Contains(secondPart, string(item)) {
				priority += getPriority(byte(item))
				break
			}
		}
	}
	fmt.Println(priority)
}
