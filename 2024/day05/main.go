package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Andreas-Evripidou/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	// fmt.Println("Day 05 input:", input)

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
	rulesString, booksString := utils.ParseTwoPartInput(input, "\n\n")
	rulesList := utils.ParseListOfListOfInts(rulesString, "|")
	booksList := utils.ParseListOfListOfInts(booksString, ",")

	sum := 0
	for _, book := range booksList {
		sum += checkBook(book, rulesList)
	}
	return strconv.Itoa(sum)
}

func solvePart2(input string) string {
	rulesString, booksString := utils.ParseTwoPartInput(input, "\n\n")
	rulesList := utils.ParseListOfListOfInts(rulesString, "|")
	booksList := utils.ParseListOfListOfInts(booksString, ",")

	sum := 0
	for _, book := range booksList {

		if checkBook(book, rulesList) > 0 {
			continue
		}

		sum += correctBook(book, rulesList)
	}

	return strconv.Itoa(sum)
}

func checkBook(book []int, rules [][]int) int {
	for _, rule := range rules {
		small := findIndex(book, rule[0])
		if small == -1 {
			continue
		}

		big := findIndex(book, rule[1])
		if big == -1 {
			continue
		}

		if big < small {
			return 0
		}
	}
	return book[len(book)/2]
}

func correctBook(book []int, rules [][]int) int {
	for ruleIndex, rule := range rules {
		small := findIndex(book, rule[0])
		if small == -1 {
			continue
		}

		big := findIndex(book, rule[1])
		if big == -1 {
			continue
		}

		if big < small {
			book = moveItemToIndex(book, big, small)
			correctBook(book, rules[:ruleIndex-1])
		}
	}

	return book[len(book)/2]
}

func moveItemToIndex(slice []int, from, to int) []int {
	item := slice[from]

	slice = append(slice[:from], slice[from+1:]...)
	slice = append(slice[:to], append([]int{item}, slice[to:]...)...)

	return slice
}

func findIndex(slice []int, val int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == val {
			return i
		}
	}

	return -1
}
