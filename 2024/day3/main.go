package main

import (
	"fmt"
	"regexp"
	"strings"
	"utils"
)

func ComputeMul(value string) int {
	rawValues := strings.Split(value, ",")
	left := utils.StringToInt(rawValues[0][4:])
	right := utils.StringToInt(rawValues[1][:len(rawValues[1])-1])

	return left * right
}

func Part1() int {
	fileContent := utils.ReadInput("./input.txt")
	pattern := `mul\([0-9]{1,3}\,[0-9]{1,3}\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	matches := re.FindAllString(fileContent, -1)
	if matches == nil {
		return 0
	}

	result := 0
	for _, match := range matches {
		result += ComputeMul(match)
	}
	return result
}

var processMul bool = true

func ComputeInstruction(value string) int {
	if strings.HasPrefix(value, "mul(") && processMul {
		return ComputeMul(value)
	}
	if value == "don't()" {
		processMul = false
	} else if value == "do()" {
		processMul = true
	}
	return 0
}

func Part2() int {
	fileContent := utils.ReadInput("./input.txt")
	pattern := `(mul\([0-9]{1,3}\,[0-9]{1,3}\))|(do\(\))|(don\'t\(\))`
	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	matches := re.FindAllString(fileContent, -1)
	if matches == nil {
		return 0
	}

	result := 0
	for _, match := range matches {
		result += ComputeInstruction(match)
	}
	return result
}

func main() {
	fmt.Println("Day 03:\n-------")
	fmt.Println("Part 1:", Part1())
	fmt.Println("Part 2:", Part2())
	fmt.Println("-------")
}
