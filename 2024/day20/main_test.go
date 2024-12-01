package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputFile = "testInput.txt"
var testInput string

var solutionPart1 = ""
var solutionPart2 = ""

func init() {
	readBytes, err := os.ReadFile(testInputFile)
	if err != nil {
		fmt.Printf("Failed to read [%!s(int=20)] with error [%!s(MISSING)]", testInputFile, err)
	}

	if string(readBytes) == "" {
		fmt.Printf("Please update the [%!s(MISSING)] file", testInputFile)
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