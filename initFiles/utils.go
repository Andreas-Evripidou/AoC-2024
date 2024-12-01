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

func ParsTwoCountMaps(input string, stringBetweenLines string) (map[string]int, map[string]int) {
	lines := parseLines(input)

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
	lines := parseLines(input)
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

func parseLines(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}

func GetIntAbsoluteValue(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
