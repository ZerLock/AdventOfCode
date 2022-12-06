package main

import (
	"fmt"
	"os"
)

func isMarker(arr string) bool {
	m := make(map[rune]bool)
	for _, i := range arr {
		_, ok := m[i]
		if ok {
			return false
		}
		m[i] = true
	}
	return true
}

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)

	for i := 0; i < len(fileString)-4; i++ {
		if isMarker(fileString[i : i+4]) {
			fmt.Println("First marker position : ", i+4)
			break
		}
	}
}
