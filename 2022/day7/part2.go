package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func goOutOfDirectory(dirSizes *[]int, directories *[]int) {
	directorySize := (*directories)[len((*directories))-1]  // Get directory size
	(*dirSizes) = append((*dirSizes), directorySize)        // Add this directory size in the "total" array
	(*directories) = (*directories)[:len((*directories))-1] // delete current directory because we go outside it
	if len((*directories)) > 0 {                            // If there is not no directory haha
		(*directories)[len((*directories))-1] += directorySize // We add the directory size to the new current directory
	}
}

func main() {
	fileBytes, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	fileString := string(fileBytes)
	lines := strings.Split(fileString, "\n")

	var dirSizes []int
	var directories []int

	for _, line := range lines {
		cmd := line[:4] // Get first command
		switch cmd {
		case "$ cd":
			path := line[5:]  // Get command parameters
			if path != ".." { // go into new directory
				directories = append(directories, 0) // New directory of size 0 at the moment
			} else {
				goOutOfDirectory(&dirSizes, &directories) // Go out of the directory (cd ..)
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

	// We go back to the root to add directory size in the "total" array
	for len(directories) > 0 {
		goOutOfDirectory(&dirSizes, &directories)
	}

	// Sort to have the smaller directory first
	sort.Ints(dirSizes)

	// Get used and free memory
	used := dirSizes[len(dirSizes)-1]
	free := 70000000 - used

	// Search directory to delete
	i := sort.Search(len(dirSizes), func(i int) bool {
		return free+dirSizes[i] >= 30000000
	})

	fmt.Println(dirSizes[i])
}
