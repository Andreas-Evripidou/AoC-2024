package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Andreas-Evripidou/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Day 06 input:", input)

	if input == "" {
		fmt.Print("Please update the input.txt\n")
		os.Exit(1)
	}

	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))
}

func solvePart1(input string) string {
	mapData := utils.ParseLines(input)

	visitedPlaces, _ := simulateGuard(mapData)

	return strconv.Itoa(visitedPlaces)
}

func solvePart2(input string) string {
	mapData := utils.ParseLines(input)

	_, newObstacles := simulateGuard(mapData)

	return strconv.Itoa(newObstacles)
}

var directions = []utils.Point{
	{X: 0, Y: -1}, // Up (^)
	{X: 1, Y: 0},  // Right (>)
	{X: 0, Y: 1},  // Down (v)
	{X: -1, Y: 0}, // Left (<)
}

func simulateGuard(mapData []string) (int, int) {
	rows := len(mapData)
	cols := len(mapData[0])
	guardPosition, guardDirection := findGuardsPositionAndDirection(mapData)

	visited := make(map[utils.Point]bool)
	visited[guardPosition] = true
	placed := make(map[utils.Point]bool)

	sumNewObstacles := 0
	for {
		nextPos := utils.Point{
			X: guardPosition.X + directions[guardDirection].X,
			Y: guardPosition.Y + directions[guardDirection].Y,
		}

		if nextPos.X < 0 || nextPos.Y < 0 || nextPos.X >= cols || nextPos.Y >= rows {
			break
		}

		if mapData[nextPos.Y][nextPos.X] == '#' {
			guardDirection = (guardDirection + 1) % 4
			continue
		}

		if !placed[nextPos] {
			mapWithNewObstacle := make([]string, len(mapData))
			copy(mapWithNewObstacle, mapData)
			rowToAddObstacle := []rune(mapWithNewObstacle[nextPos.Y])
			rowToAddObstacle[nextPos.X] = '#'
			mapWithNewObstacle[nextPos.Y] = string(rowToAddObstacle)
			if resultsToInfinityLoop(mapWithNewObstacle, guardPosition, guardDirection) {
				sumNewObstacles++
			}
			placed[nextPos] = true
		}

		guardPosition = nextPos
		visited[guardPosition] = true
	}

	return len(visited), sumNewObstacles
}

func resultsToInfinityLoop(mapData []string, guardPosition utils.Point, guardDirection int) bool {
	rows := len(mapData)
	cols := len(mapData[0])

	visited := make(map[utils.Point]map[int]bool)
	visited[guardPosition] = make(map[int]bool)
	visited[guardPosition][guardDirection] = true

	for {
		nextPos := utils.Point{
			X: guardPosition.X + directions[guardDirection].X,
			Y: guardPosition.Y + directions[guardDirection].Y,
		}

		if nextPos.X < 0 || nextPos.Y < 0 || nextPos.X >= cols || nextPos.Y >= rows {
			break
		}

		if mapData[nextPos.Y][nextPos.X] == '#' {
			guardDirection = (guardDirection + 1) % 4
			continue
		}

		if visited[nextPos] != nil && visited[nextPos][guardDirection] {
			return true
		}

		guardPosition = nextPos
		if visited[guardPosition] == nil {
			visited[guardPosition] = make(map[int]bool)
		}
		visited[guardPosition][guardDirection] = true
	}

	return false
}

func findGuardsPositionAndDirection(mapData []string) (utils.Point, int) {
	for y, row := range mapData {
		for x, cell := range row {
			if cell == '.' || cell == '#' {
				continue
			}

			return calculatePositionAndDirection(cell, x, y)
		}
	}

	return utils.Point{}, -1
}

func calculatePositionAndDirection(cell rune, x int, y int) (utils.Point, int) {
	if cell == '^' {
		return utils.Point{X: x, Y: y}, 0
	} else if cell == '>' {
		return utils.Point{X: x, Y: y}, 1
	} else if cell == 'v' {
		return utils.Point{X: x, Y: y}, 2
	}

	return utils.Point{X: x, Y: y}, 3
}
