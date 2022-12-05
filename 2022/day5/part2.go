package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Convert a string to an intenger
func toInteger(str string) int {
	data, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return data
}

// Get datas from the line instruction
func getLineData(line string) (int, int, int) {
	data := strings.Split(line, " ")
	from := toInteger(data[3]) - 1
	to := toInteger(data[5]) - 1
	index := toInteger(data[1])
	return from, to, index
}

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	separatedParts := strings.Split(fileString, "\n\n")
	var stacks [9][]string

	// Create stacks
	for _, line := range strings.Split(separatedParts[0], "\n") {
		for index, item := range line {
			if index%4 == 1 && item != ' ' {
				stacks[index/4] = append([]string{string(item)}, stacks[index/4]...)
			}
		}
	}

	// Move boxes
	for _, move := range strings.Split(separatedParts[1], "\n") {
		from, to, index := getLineData(move)
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-index:]...)
		stacks[from] = stacks[from][:len(stacks[from])-index]
	}

	fmt.Println(stacks)
}
