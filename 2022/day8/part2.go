package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")
	var total int

	for y, line := range lines {
		for x, tree := range line {
			tmpTotal := 0
			treeValue := byte(tree)
			if y == 0 || x == 0 || x == len(line)-1 || y == len(lines)-1 {
				tmpTotal++
			}

			// haut
			tmpHaut := y - 1
			haut := 0
			for tmpHaut >= 0 {
				haut++
				if lines[tmpHaut][x] >= treeValue {
					break
				}
				tmpHaut--
			}

			// droite
			tmpDroite := x + 1
			droite := 0
			for tmpDroite < len(line) {
				droite++
				if lines[y][tmpDroite] >= treeValue {
					break
				}
				tmpDroite++
			}

			// bas
			tmpBas := y + 1
			bas := 0
			for tmpBas < len(lines) {
				bas++
				if lines[tmpBas][x] >= treeValue {
					break
				}
				tmpBas++
			}

			// gauche
			tmpGauche := x - 1
			gauche := 0
			for tmpGauche >= 0 {
				gauche++
				if lines[y][tmpGauche] >= treeValue {
					break
				}
				tmpGauche--
			}

			tmp := haut * droite * bas * gauche
			if tmp > total {
				total = tmp
			}
		}
	}

	fmt.Println(total)
}
