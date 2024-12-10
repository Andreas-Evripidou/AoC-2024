package main

import (
	"fmt"
	"os"
	"strconv"

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
	gridMap := utils.ParseListOfListOfInts(input, "")
	fmt.Print(gridMap)

	trailHeads := findAllTrailHeads(gridMap)

	sumScore := 0
	for _, trailHead := range trailHeads {
		sumScore += len(getTrailHeadFinishStates(trailHead, gridMap))
	}

	return strconv.Itoa(sumScore)
}

var directions = []utils.Point{
	{X: 0, Y: -1}, // Up (^)
	{X: 1, Y: 0},  // Right (>)
	{X: 0, Y: 1},  // Down (v)
	{X: -1, Y: 0}, // Left (<)
}

func getTrailHeadFinishStates(currentPosition utils.Point, gridMap [][]int) map[utils.Point]bool {
	finishStates := map[utils.Point]bool{}

	currentHeight := gridMap[currentPosition.Y][currentPosition.X]
	if currentHeight == 9 {
		finishStates[currentPosition] = true
		return finishStates
	}
	for _, direction := range directions {
		nextPos := utils.Point{
			X: currentPosition.X + direction.X,
			Y: currentPosition.Y + direction.Y,
		}

		if nextPos.X >= len(gridMap[0]) || nextPos.Y >= len(gridMap) ||
			nextPos.X < 0 || nextPos.Y < 0 {
			continue
		}

		nextHeight := gridMap[nextPos.Y][nextPos.X]
		if nextHeight == currentHeight+1 {
			newFinishStates := getTrailHeadFinishStates(nextPos, gridMap)

			for key, value := range newFinishStates {
				finishStates[key] = value
			}
		}
	}

	return finishStates
}

func findAllTrailHeads(gridMap [][]int) (trailHeads []utils.Point) {
	trailHeads = []utils.Point{}
	for y, row := range gridMap {
		for x := range row {
			if gridMap[y][x] == 0 {
				trailHeads = append(trailHeads, utils.Point{X: x, Y: y})
			}
		}
	}
	return trailHeads
}

func solvePart2(input string) string {
	gridMap := utils.ParseListOfListOfInts(input, "")
	fmt.Print(gridMap)

	trailHeads := findAllTrailHeads(gridMap)

	sumScore := 0
	for _, trailHead := range trailHeads {
		sumScore += getTrailHeadSuccessfulTrails(trailHead, gridMap)
	}

	return strconv.Itoa(sumScore)
}

func getTrailHeadSuccessfulTrails(currentPosition utils.Point, gridMap [][]int) int {
	sum := 0

	currentHeight := gridMap[currentPosition.Y][currentPosition.X]
	if currentHeight == 9 {
		return 1
	}
	for _, direction := range directions {
		nextPos := utils.Point{
			X: currentPosition.X + direction.X,
			Y: currentPosition.Y + direction.Y,
		}

		if nextPos.X >= len(gridMap[0]) || nextPos.Y >= len(gridMap) ||
			nextPos.X < 0 || nextPos.Y < 0 {
			continue
		}

		nextHeight := gridMap[nextPos.Y][nextPos.X]
		if nextHeight == currentHeight+1 {
			sum += getTrailHeadSuccessfulTrails(nextPos, gridMap)
		}
	}

	return sum
}
