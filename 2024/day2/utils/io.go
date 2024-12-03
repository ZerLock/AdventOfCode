package utils

import "os"

func ReadFileToString(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}
