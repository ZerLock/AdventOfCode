package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	var directories []int
	var total int

	for _, line := range lines {
		cmd := line[:4]
		switch cmd {
		case "$ cd":
			path := line[5:]
			if path != ".." { // go into new directory
				directories = append(directories, 0)
				continue
			}

			// "cd .." return out of current directory so add it size to the total size if it's less than 100000
			directorySize := directories[len(directories)-1]
			if directorySize <= 100000 {
				total += directorySize
			}

			// If there is not no directory add to the new directory the precedent directory size
			directories = directories[:len(directories)-1]
			if len(directories) > 0 {
				directories[len(directories)-1] += directorySize
			}
		case "$ ls", "dir ": // Do nothing
		default: // file
			file := strings.Fields(line)           // Split line
			fileSize, err := strconv.Atoi(file[0]) // Get (int) file size with splitted line
			if err != nil {
				panic(err)
			}
			directories[len(directories)-1] += fileSize // Increment the size of the last directory (current directory when processing)
		}
	}

	fmt.Println(total)
}
