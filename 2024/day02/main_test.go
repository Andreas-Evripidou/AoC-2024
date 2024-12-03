package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputFile = "testInput.txt"
var testInput string

var solutionPart1 = "2"
var solutionPart2 = "4"

func init() {
	readBytes, err := os.ReadFile(testInputFile)
	if err != nil {
		fmt.Printf("Failed to read %s with error %s", testInputFile, err)
	}

	if string(readBytes) == "" {
		fmt.Printf("Please update the %s file", testInputFile)
		os.Exit(1)
	}

	testInput = string(readBytes)
}

func TestSolvePart1(t *testing.T) {
	assert.NotEmpty(t, solutionPart1)
	assert.Equal(t, solutionPart1, solvePart1(testInput), "Please provide the test solution for part 1")
}

func TestSolvePart2(t *testing.T) {
	assert.NotEmpty(t, solutionPart2, "Please provide the test solution for part 2")
	assert.Equal(t, solutionPart2, solvePart2(testInput))
}

func Test_calculateSumOfValidReports(t *testing.T) {
	tests := []struct {
		description          string
		inputReportsTable    [][]int
		inputToleration      int
		expectedValidReports int
	}{
		{
			description: "3 valid reports with zero toleration",
			inputReportsTable: [][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			inputToleration:      0,
			expectedValidReports: 3,
		},
		{
			description: "3 valid reports with 1 mistake toleration",
			inputReportsTable: [][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			inputToleration:      1,
			expectedValidReports: 3,
		},
		{
			description:          "1 report with one mistake and 1 mistake toleration",
			inputReportsTable:    [][]int{{1, 2, 7}},
			inputToleration:      1,
			expectedValidReports: 1,
		},
		{
			description:          "1 report with one mistake and 2 mistake toleration",
			inputReportsTable:    [][]int{{1, 7, 17}},
			inputToleration:      1,
			expectedValidReports: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			validReports := calculateSumOfValidReports(test.inputReportsTable, test.inputToleration)
			assert.Equal(t, validReports, test.expectedValidReports)
		})
	}
}

func TestIsSortedWithTolerance(t *testing.T) {
	tests := []struct {
		description    string
		inputSlice     []int
		toleration     int
		expectedOutput bool
	}{
		{
			description:    "empty slice",
			inputSlice:     []int{},
			toleration:     0,
			expectedOutput: true,
		},
		{
			description:    "single element slice",
			inputSlice:     []int{1},
			toleration:     0,
			expectedOutput: true,
		},
		{
			description:    "ascending slice",
			inputSlice:     []int{1, 2, 3, 4, 5},
			toleration:     0,
			expectedOutput: true,
		},
		{
			description:    "descending slice",
			inputSlice:     []int{5, 4, 3, 2, 1},
			toleration:     0,
			expectedOutput: true,
		},
		{
			description:    "ascending slice with violation and one mistake tolerance",
			inputSlice:     []int{1, 2, 3, 4, 6},
			toleration:     1,
			expectedOutput: true,
		},
		{
			description:    "descending slice with violation and one mistake tolerance",
			inputSlice:     []int{6, 5, 4, 3, 2, 1},
			toleration:     1,
			expectedOutput: true,
		},
		{
			description:    "ascending slice with violation and zero tolerance",
			inputSlice:     []int{1, 2, 3, 4, 5, 4, 6},
			toleration:     0,
			expectedOutput: false,
		},
		{
			description:    "descending slice with violation and zero tolerance",
			inputSlice:     []int{7, 6, 5, 4, 3, 2, 3},
			toleration:     0,
			expectedOutput: false,
		},
	}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			actualOutput := isSortedWithTolerance(test.inputSlice, test.toleration)
			assert.Equal(t, actualOutput, test.expectedOutput)
		})
	}
}
