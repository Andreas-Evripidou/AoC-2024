package main

import (
	"fmt"
	"os"

	"github.com/Andreas-Evripidou/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Day 10 input:", input)

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
	// TODO: Implement Part 1
	return "Not implemented"
}

func solvePart2(input string) string {
	// TODO: Implement Part 2
	return "Not implemented"
}
