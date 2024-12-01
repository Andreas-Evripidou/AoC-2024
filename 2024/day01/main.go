package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/Andreas-Evripidou/AoC-2024/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Day 01 input:", input)

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
	list1, list2 := utils.ParseTwoNumberLists(input, "   ")

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0
	for index, _ := range list1 {
		sum += utils.GetIntAbsoluteValue(list1[index] - list2[index])
	}

	return strconv.Itoa(sum)
}

func solvePart2(input string) string {
	countMap1, countMap2 := utils.ParsTwoCountMaps(input, "   ")

	sum := 0
	for key := range countMap1 {
		num, err := strconv.Atoi(key)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		sum += num * countMap1[key] * countMap2[key]
		delete(countMap2, key)
	}
	for key := range countMap2 {
		num, err := strconv.Atoi(key)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		sum += num * countMap1[key] * countMap2[key]
	}

	return strconv.Itoa(sum)
}
