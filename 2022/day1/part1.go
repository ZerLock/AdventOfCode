package main

import (
	"fmt"
	"os"
	"strings"
)

var getPoints = map[byte]int{
	'A': 1,
	'B': 2,
	'C': 3,
}

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	rounds := strings.Split(fileString, "\n")
	points := 0

	for _, round := range rounds {
		enemy := round[0]
		me := round[2] - ('X' - 'A')
		points += getPoints[me]
		if enemy == me {
			points += 3
		} else if (me-enemy == 1) || (enemy == 'C' && me == 'A') {
			points += 6
		}
	}
	fmt.Println(points)
}
