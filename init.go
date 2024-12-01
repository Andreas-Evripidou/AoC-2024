package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	rootDir := flag.String("year", "", "Root directory for the Advent of Code structure")
	flag.Parse()

	if *rootDir == "" {
		fmt.Println("Error: The -year flag is required. Please specify the year in order to create a year directory.")
		os.Exit(1)
	}

	utilsDir := filepath.Join(*rootDir, "utils")
	numDays := 25

	err := createRootDirectory(*rootDir)
	if err != nil {
		fmt.Printf("Failed to create root directory: %v\n", err)
		return
	}
	fmt.Println("Created root directory:", rootDir)

	err = createUtilsFolder(utilsDir)
	if err != nil {
		fmt.Printf("Failed to create utils directory: %v\n", err)
		return
	}
	fmt.Println("Created utils directory:", utilsDir)

	err = createUtilFiles(utilsDir)
	if err != nil {
		fmt.Printf("Failed to create utils.go: %v\n", err)
		return
	}
	fmt.Println("Created utils.go file in utils directory")

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
		mainContent := fmt.Sprintf(`package main

import (
	"fmt"
	"../utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Day %02d input:", input)

	// Part 1
	fmt.Println("Part 1:", solvePart1(input))

	// Part 2
	fmt.Println("Part 2:", solvePart2(input))
}

func solvePart1(input string) string {
	// TODO: Implement Part 1
	return "Not implemented"
}

func solvePart2(input string) string {
	// TODO: Implement Part 2
	return "Not implemented"
}
`, day)
		err = os.WriteFile(mainFile, []byte(mainContent), 0644)
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

		fmt.Printf("Created structure for day %02d\n", day)
	}
}

func createUtilFiles(utilsDir string) error {
	utilsFile := filepath.Join(utilsDir, "utils.go")
	err := os.WriteFile(utilsFile, []byte(`package utils

import (
	"log"
	"os"
	"strings"
)

// ReadInput reads the content of a file and returns it as a string
func ReadInput(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	return strings.TrimSpace(string(data))
}
`), 0644)
	return err
}

func createUtilsFolder(utilsDir string) error {
	err := os.Mkdir(utilsDir, 0755)
	if os.IsExist(err) {

		return nil
	}
	return err
}

func createRootDirectory(rootDir string) error {
	err := os.Mkdir(rootDir, 0755)
	if os.IsExist(err) {
		return nil
	}
	return err
}
