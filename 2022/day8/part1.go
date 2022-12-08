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
			haut := true
			for tmpHaut >= 0 {
				if lines[tmpHaut][x] >= treeValue {
					haut = false
					break
				}
				tmpHaut--
			}

			// droite
			tmpDroite := x + 1
			droite := true
			for tmpDroite < len(line) {
				if lines[y][tmpDroite] >= treeValue {
					droite = false
					break
				}
				tmpDroite++
			}

			// bas
			tmpBas := y + 1
			bas := true
			for tmpBas < len(lines) {
				if lines[tmpBas][x] >= treeValue {
					bas = false
					break
				}
				tmpBas++
			}

			// gauche
			tmpGauche := x - 1
			gauche := true
			for tmpGauche >= 0 {
				if lines[y][tmpGauche] >= treeValue {
					gauche = false
					break
				}
				tmpGauche--
			}

			if haut || droite || gauche || bas {
				total++
			}
		}
	}

	fmt.Println(total)
}
