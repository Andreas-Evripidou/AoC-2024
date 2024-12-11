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
	fmt.Println("Day 11 input:", input)

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
	stringStones := strings.Split(input, " ")

	stones := []int{}
	for _, stone := range stringStones {
		intStone, _ := strconv.Atoi(stone)
		stones = append(stones, intStone)
	}
	for range 25 {
		stones = blink(stones)
	}

	return strconv.Itoa(len(stones))
}

func solvePart2(input string) string {
	stringStones := strings.Split(input, " ")

	stones := map[int]int{}
	for _, stone := range stringStones {
		intStone, _ := strconv.Atoi(stone)
		stones[intStone] = 1
	}
	for range 75 {
		stones = efficientBlink(stones)
	}
	count := countStones(stones)

	return strconv.Itoa(count)
}

func countStones(stones map[int]int) int {
	sum := 0
	for _, count := range stones {
		sum += count
	}
	return sum
}

func blink(stones []int) []int {
	newStones := []int{}

	for _, stone := range stones {
		stringStone := strconv.Itoa(stone)

		if stone == 0 {
			newStones = append(newStones, 1)
		} else if len(stringStone)%2 == 0 {
			newStones = append(newStones, splitStone(stringStone)...)
		} else {
			newStones = append(newStones, stone*2024)
		}
	}

	return newStones
}

func splitStone(stringStone string) []int {
	stoneLen := len(stringStone)
	leftStone, _ := strconv.Atoi(stringStone[:stoneLen/2])
	rightStone, _ := strconv.Atoi(stringStone[stoneLen/2:])
	return append([]int{}, leftStone, rightStone)
}

func efficientBlink(stones map[int]int) map[int]int {
	newStones := make(map[int]int)
	for stone, count := range stones {
		if stone == 0 {
			newStones[1] += count
			continue
		}

		stringStone := strconv.Itoa(stone)
		if len(stringStone)%2 == 0 {
			efficientSplitStone(stringStone, newStones, count)
		} else {
			newStones[stone*2024] += count
		}
	}

	return newStones
}

func efficientSplitStone(stringStone string, newStones map[int]int, count int) {
	mid := len(stringStone) / 2
	left, _ := strconv.Atoi(stringStone[:mid])
	right, _ := strconv.Atoi(stringStone[mid:])
	newStones[left] += count
	newStones[right] += count
}
