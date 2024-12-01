package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := flag.String("year", "", "Root directory for the Advent of Code structure")
	flag.Parse()

	if *rootDir == "" {
		fmt.Println("Error: The -year flag is required. Please specify the year in order to create a year directory.")
		os.Exit(1)
	}

	numDays := 25

	err := createRootDirectory(*rootDir)
	if err != nil {
		fmt.Printf("Failed to create root directory: %v\n", err)
		return
	}
	fmt.Println("Created root directory:", rootDir)

	createChallengesDict(numDays, *rootDir)
}

func createChallengesDict(numDays int, rootDir string) {
	for day := 1; day <= numDays; day++ {
		dayDir := filepath.Join(rootDir, fmt.Sprintf("day%02d", day))
		err := os.Mkdir(dayDir, 0755)
		if err != nil && !os.IsExist(err) {
			fmt.Printf("Failed to create directory for day %02d: %v\n", day, err)
			continue
		}

		mainFile := filepath.Join(dayDir, "main.go")
		mainTemplate := readInput("initFiles/main.go")
		mainContent := fmt.Sprintf(mainTemplate, day)
		err = os.WriteFile(mainFile, []byte(mainContent), 0644)
		if err != nil {
			fmt.Printf("Failed to create main.go for day %02d: %v\n", day, err)
			continue
		}

		mainTestFile := filepath.Join(dayDir, "main_test.go")
		mainTestTemplate := readInput("initFiles/main_Test.go")
		mainTestContent := fmt.Sprintf(mainTestTemplate, day)
		err = os.WriteFile(mainTestFile, []byte(mainTestContent), 0644)
		if err != nil {
			fmt.Printf("Failed to create main.go for day %02d: %v\n", day, err)
			continue
		}

		inputFile := filepath.Join(dayDir, "input.txt")
		err = os.WriteFile(inputFile, []byte{}, 0644)
		if err != nil {
			fmt.Printf("Failed to create input.txt for day %02d: %v\n", day, err)
			continue
		}

		testInputFile := filepath.Join(dayDir, "testInput.txt")
		err = os.WriteFile(testInputFile, []byte{}, 0644)
		if err != nil {
			fmt.Printf("Failed to create testInput.txt for day %02d: %v\n", day, err)
			continue
		}

		fmt.Printf("Created structure for day %02d\n", day)
	}
}

func createRootDirectory(rootDir string) error {
	err := os.Mkdir(rootDir, 0755)
	if os.IsExist(err) {
		return nil
	}
	return err
}

func readInput(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
	}
	return strings.TrimSpace(string(data))
}
