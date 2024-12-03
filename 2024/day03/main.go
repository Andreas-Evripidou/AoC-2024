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
	fmt.Println("Day 03 input:", input)

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
	sum := 0
	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	multiplicationList := r.FindAllStringSubmatch(input, -1)

	for _, multiplication := range multiplicationList {
		sum += utils.StringToInt(multiplication[1]) * utils.StringToInt(multiplication[2])
	}

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {
	sanitizedInput := ""
	for _, line := range strings.Split(input, "do()") {
		sanitizedInput += strings.Split(line, "don't()")[0]
	}

	sum := 0
	r, _ := regexp.Compile(`mul\(([0-9]+),([0-9]+)\)`)
	multiplicationList := r.FindAllStringSubmatch(sanitizedInput, -1)

	for _, multiplication := range multiplicationList {
		sum += utils.StringToInt(multiplication[1]) * utils.StringToInt(multiplication[2])
	}

	return strconv.Itoa(sum)
}
