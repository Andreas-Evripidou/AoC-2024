package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/Andreas-Evripidou/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Day 04 input:", input)

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
	stringGrid := utils.ParseListOfListOfString(input)

	stringToSearch := input + "\n"

	columnsLength := len(stringGrid[0])
	rowsLength := len(stringGrid)
	for i := range columnsLength {
		verticalString := ""
		for j := range rowsLength {
			verticalString += stringGrid[j][i]
		}
		stringToSearch += verticalString + "\n"
	}

	diagonalsString := extractLargeDiagonals(stringGrid)
	stringToSearch += diagonalsString

	r, _ := regexp.Compile("XMAS")
	sum := len(r.FindAllString(stringToSearch, -1))
	r, _ = regexp.Compile("SAMX")
	sum += len(r.FindAllString(stringToSearch, -1))

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {
	stringGrid := utils.ParseListOfListOfString(input)

	occurrences := 0

	columnsLength := len(stringGrid[0])
	rowsLength := len(stringGrid)
	for i := range columnsLength - 2 {
		for j := range rowsLength - 2 {
			middle := stringGrid[j+1][i+1]
			if middle != "A" {
				continue
			}
			diagonal1 := stringGrid[j][i] + middle + stringGrid[j+2][i+2]
			diagonal2 := stringGrid[j+2][i] + middle + stringGrid[j][i+2]

			if (diagonal1 == "MAS" || diagonal1 == "SAM") && (diagonal2 == "MAS" || diagonal2 == "SAM") {
				occurrences++
			}

		}

	}

	return strconv.Itoa(occurrences)
}

func extractLargeDiagonals(grid [][]string) string {
	n := len(grid)
	if n == 0 {
		return ""
	}

	m := len(grid[0])
	var diagonals []string

	// Extract top-left to bottom-right diagonals
	for d := 0; d < n+m-1; d++ {
		var diagonal string
		for i := 0; i < n; i++ {
			j := d - i
			if j >= 0 && j < m {
				diagonal += grid[i][j]
			}
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}

	// Extract top-right to bottom-left diagonals
	for d := 0; d < n+m-1; d++ {
		var diagonal string
		for i := 0; i < n; i++ {
			j := i - d + (m - 1)
			if j >= 0 && j < m {
				diagonal += grid[i][j]
			}
		}
		if len(diagonal) >= 4 {
			diagonals = append(diagonals, diagonal)
		}
	}

	// Combine diagonals with "\n" separator
	return strings.Join(diagonals, "\n")
}
