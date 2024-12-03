package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Andreas-Evripidou/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Day 02 input:", input)

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
	reports := utils.ParseListOfListOfInts(input, " ")

	sum := calculateSumOfValidReports(reports, 0)

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {
	reports := utils.ParseListOfListOfInts(input, " ")

	sum := calculateSumOfValidReports(reports, 1)

	return strconv.Itoa(sum)
}

func calculateSumOfValidReports(reports [][]int, toleration int) int {
	sum := 0
	for _, levels := range reports {

		if !isSortedWithTolerance(levels, toleration) {
			continue
		}

		if areValidLevels(levels) {
			sum++
			continue
		}
		if toleration == 0 {
			continue
		}

		for itemToRemove := 0; itemToRemove < len(levels); itemToRemove++ {
			newLevels := append([]int{}, levels[:itemToRemove]...)    // Elements before the i-th
			newLevels = append(newLevels, levels[itemToRemove+1:]...) // Elements after the i-th
			if areValidLevels(newLevels) {
				sum++
				break
			}
		}

	}
	return sum
}

func areValidLevels(levels []int) bool {
	for index := 0; index < len(levels)-1; index++ {
		distance := utils.GetIntAbsoluteValue(levels[index] - levels[index+1])

		if distance < 1 || distance > 3 {
			return false
		}

	}
	return true
}

func isSortedWithTolerance(slice []int, toleration int) bool {
	if len(slice) < 2 {
		return true
	}

	ascViolations, descViolations := 0, 0

	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			ascViolations++
		} else if slice[i] > slice[i-1] {
			descViolations++
		}

		if ascViolations > toleration && descViolations > toleration {
			return false
		}
	}

	return true
}
