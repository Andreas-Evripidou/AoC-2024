package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

var year = "2024"

// ReadInput reads the content of a file and returns it as a string
func ReadInput(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return strings.TrimSpace(string(data))
}

func ParseTwoPartInput(input string, stringBetweenParts string) (part1 string, part2 string) {
	parts := strings.Split(input, stringBetweenParts)
	return parts[0], parts[1]
}

func ParseListOfListOfInts(input string, stringBetweenLines string) [][]int {
	lines := ParseLines(input)

	mainList := make([][]int, 0)
	for _, line := range lines {
		splitLine := strings.Split(line, stringBetweenLines)
		subList := make([]int, 0)
		for _, splitSubLine := range splitLine {
			if splitSubLine == "." {
				splitSubLine = "-1"
			}
			intSplitLine, _ := strconv.Atoi(splitSubLine)
			subList = append(subList, intSplitLine)
		}
		mainList = append(mainList, subList)
	}

	return mainList
}

func ParseListOfListOfString(input string) [][]string {
	lines := ParseLines(input)

	mainList := make([][]string, 0)
	for _, line := range lines {
		subList := make([]string, 0)
		for _, character := range line {
			subList = append(subList, string(character))
		}
		mainList = append(mainList, subList)
	}

	return mainList
}

func ParsTwoCountMaps(input string, stringBetweenLines string) (map[string]int, map[string]int) {
	lines := ParseLines(input)

	var map1 = make(map[string]int)
	var map2 = make(map[string]int)

	for _, line := range lines {
		splitLine := strings.Split(line, stringBetweenLines)

		map1[splitLine[0]]++
		map2[splitLine[1]]++
	}

	return map1, map2
}

func ParseTwoNumberLists(input string, stringBetweenLines string) ([]int, []int) {
	lines := ParseLines(input)
	fileLength := len(lines)

	list1 := make([]int, fileLength)
	list2 := make([]int, fileLength)
	for i, line := range lines {
		splitLine := strings.Split(line, stringBetweenLines)
		list1[i], _ = strconv.Atoi(splitLine[0])
		list2[i], _ = strconv.Atoi(splitLine[1])
	}
	return list1, list2
}

func ParseLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func GetIntAbsoluteValue(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func StringToInt(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}

//------------------ Map related -----------------------\\

type Point struct {
	X int
	Y int
}
