package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Andreas-Evripidou/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Day 07 input:", input)

	if input == "" {
		fmt.Print("Please update the input.txt\n")
		os.Exit(1)
	}

	// Part 1
	fmt.Println("Part 1:", solvePart1(input))

	// Part 2
	fmt.Println("Part 2:", solvePart2(input))
}

func solvePart1(input string) string {
	lines := utils.ParseLines(input)
	validEquations := 0
	for _, line := range lines {
		splitString := strings.SplitN(line, ": ", 2)
		expectedSum, _ := strconv.Atoi(splitString[0])

		stringNumbers := strings.Split(splitString[1], " ")
		numbers := make([]int, len(stringNumbers))
		for i, n := range stringNumbers {
			numbers[i], _ = strconv.Atoi(n)
		}

		if checkRecursive(numbers, expectedSum, numbers[0], 1) {
			validEquations += expectedSum
		}

	}

	return strconv.Itoa(validEquations)
}

func solvePart2(input string) string {
	lines := utils.ParseLines(input)
	validEquations := 0
	for _, line := range lines {
		splitString := strings.SplitN(line, ": ", 2)
		expectedSum, _ := strconv.Atoi(splitString[0])

		stringNumbers := strings.Split(splitString[1], " ")
		numbers := make([]int, len(stringNumbers))
		for i, n := range stringNumbers {
			numbers[i], _ = strconv.Atoi(n)
		}

		if checkRecursive2(numbers, expectedSum, numbers[0], 1) {
			validEquations += expectedSum
		}

	}

	return strconv.Itoa(validEquations)
}

func checkRecursive(numbers []int, target, current int, index int) bool {
	if index == len(numbers) {
		return current == target
	}

	return checkRecursive(numbers, target, current+numbers[index], index+1) ||
		checkRecursive(numbers, target, current*numbers[index], index+1)
}

func checkRecursive2(numbers []int, target, current int, index int) bool {
	if index == len(numbers) {
		return current == target
	}

	concatenated := strconv.Itoa(current) + strconv.Itoa(numbers[index])
	newNum, _ := strconv.Atoi(concatenated)

	return checkRecursive2(numbers, target, current+numbers[index], index+1) ||
		checkRecursive2(numbers, target, current*numbers[index], index+1) ||
		checkRecursive2(numbers, target, newNum, index+1)
}
