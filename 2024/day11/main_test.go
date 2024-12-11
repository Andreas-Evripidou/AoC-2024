package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputFile = "testInput.txt"
var testInput string

var solutionPart1 = "55312"
var solutionPart2 = ""

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

func Test_splitStone(t *testing.T) {
	tests := []struct {
		name        string
		stringStone string
		want        []int
	}{
		{
			name:        "Test 1",
			stringStone: "12",
			want:        []int{1, 2},
		},
		{
			name:        "Test 2",
			stringStone: "1234",
			want:        []int{12, 34},
		},
		{
			name:        "Test 5",
			stringStone: "123456",
			want:        []int{123, 456},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, splitStone(tt.stringStone))
		})
	}
}
