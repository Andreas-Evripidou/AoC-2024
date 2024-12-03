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

	sum := 0
	for _, report := range reports {

		if isSafeReport(report) {
			sum++
		}
	}

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {
	reports := utils.ParseListOfListOfInts(input, " ")

	sum := 0
	for _, report := range reports {

		if isSafeReport(report) {
			sum++
			continue
		}

		if isSafeReportWithoutOneItem(report) {
			sum++
		}

	}

	return strconv.Itoa(sum)
}

func isSafeReportWithoutOneItem(report []int) bool {
	for itemToRemove := 0; itemToRemove < len(report); itemToRemove++ {
		newReport := append([]int{}, report[0:itemToRemove]...)
		newReport = append(newReport, report[itemToRemove+1:]...)
		if isSafeReport(newReport) {
			return true
		}
	}
	return false
}

func isSafeReport(report []int) bool {
	ascViolations, descViolations := 0, 0

	for index := 0; index < len(report)-1; index++ {
		distance := utils.GetIntAbsoluteValue(report[index] - report[index+1])

		if distance < 1 || distance > 3 {
			return false
		}

		if report[index] > report[index+1] {
			ascViolations++
		} else if report[index] < report[index+1] {
			descViolations++
		}

		if ascViolations > 0 && descViolations > 0 {
			return false
		}

	}
	return true
}
