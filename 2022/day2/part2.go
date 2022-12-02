package main

import (
	"fmt"
	"os"
	"strings"
)

var getLosePoints = map[byte]int{
	'A': 3,
	'B': 1,
	'C': 2,
}

var getWinPoints = map[byte]int{
	'A': 2,
	'B': 3,
	'C': 1,
}

var getWinLose = map[byte]int{
	'A': 0,
	'B': 3,
	'C': 6,
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
		finish := round[2] - ('X' - 'A')
		pointsWinLose := getWinLose[finish]
		pointsMove := 0
		if finish == 'A' {
			pointsMove = getLosePoints[enemy]
		} else if finish == 'C' {
			pointsMove = getWinPoints[enemy]
		} else {
			pointsMove = int(enemy) - 'A' + 1
		}
		points += pointsWinLose + pointsMove
	}
	fmt.Println(points)
}
