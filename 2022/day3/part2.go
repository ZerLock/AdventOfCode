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
	fileBytes, err := os.ReadFile("./input2.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	bags := strings.Split(fileString, "\n")
	groupSize := 3
	priority := 0

	for i := 0; i < len(bags); i += groupSize {
		group := bags[i : i+groupSize]
		for _, item := range group[0] {
			if strings.Contains(group[1], string(item)) && strings.Contains(group[2], string(item)) {
				priority += getPriority(byte(item))
				break
			}
		}
	}
	fmt.Println(priority)
}
