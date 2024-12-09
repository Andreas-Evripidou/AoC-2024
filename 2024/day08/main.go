package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Andreas-Evripidou/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Day 08 input:", input)

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
	antennasGrid := utils.ParseListOfListOfString(input)
	antennasLocations := findAntennaPairs(antennasGrid)

	maxRow := len(antennasGrid)
	maxCol := len(antennasGrid[0])

	uniqueAntinodes := map[utils.Point]bool{}
	for frequency := range antennasLocations {
		for pointIndex, antennaPoint := range antennasLocations[frequency] {
			for _, secondAntennaPoint := range antennasLocations[frequency][pointIndex+1:] {

				antinode1, antinode2 := calculateAntinodesInNTimesDistance(antennaPoint, secondAntennaPoint, 1)

				if isPointInBounds(antinode1, maxRow, maxCol) {
					uniqueAntinodes[antinode1] = true
				}
				if isPointInBounds(antinode2, maxRow, maxCol) {
					uniqueAntinodes[antinode2] = true
				}
			}
		}
	}

	return strconv.Itoa(len(uniqueAntinodes))
}

func solvePart2(input string) string {
	antennasGrid := utils.ParseListOfListOfString(input)
	antennasLocations := findAntennaPairs(antennasGrid)

	maxRow := len(antennasGrid)
	maxCol := len(antennasGrid[0])

	uniqueAntinodes := map[utils.Point]bool{}
	for frequency := range antennasLocations {
		for pointIndex, antennaPoint := range antennasLocations[frequency] {
			for _, secondAntennaPoint := range antennasLocations[frequency][pointIndex+1:] {
				findAllAntinodes(antennaPoint, secondAntennaPoint, maxRow, maxCol, uniqueAntinodes)

			}
		}
	}

	return strconv.Itoa(len(uniqueAntinodes))
}

func findAllAntinodes(antennaPoint utils.Point, secondAntennaPoint utils.Point, maxRow int, maxCol int, uniqueAntinodes map[utils.Point]bool) {
	n := 0

	for {
		antinode1, antinode2 := calculateAntinodesInNTimesDistance(antennaPoint, secondAntennaPoint, n)
		n++
		found := 0
		if isPointInBounds(antinode1, maxRow, maxCol) {
			uniqueAntinodes[antinode1] = true
			found++
		}
		if isPointInBounds(antinode2, maxRow, maxCol) {
			uniqueAntinodes[antinode2] = true
			found++
		}
		if found == 0 {
			break
		}
	}

}

func findAntennaPairs(grid [][]string) map[string][]utils.Point {
	antennasLocations := map[string][]utils.Point{}
	for rIndex, row := range grid {
		for cIndex, character := range row {
			if character == "." {
				continue
			}
			if antennasLocations[character] == nil {
				antennasLocations[character] = []utils.Point{}
			}

			antennasLocations[character] = append(
				antennasLocations[character],
				utils.Point{
					X: rIndex,
					Y: cIndex,
				},
			)

		}
	}
	return antennasLocations
}

func isPointInBounds(antinode1 utils.Point, maxRow int, maxCol int) bool {
	return antinode1.X >= 0 && antinode1.Y >= 0 && antinode1.X < maxRow && antinode1.Y < maxCol
}

func calculateAntinodesInNTimesDistance(a, b utils.Point, n int) (utils.Point, utils.Point) {
	dx, dy := n*(b.X-a.X), n*(b.Y-a.Y)
	antinode1 := utils.Point{X: a.X - dx, Y: a.Y - dy}
	antinode2 := utils.Point{X: b.X + dx, Y: b.Y + dy}
	return antinode1, antinode2
}
