package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Andreas-Evripidou/advent-of-code/utils"
)

func main() {
	input := utils.ReadInput("input.txt")
	fmt.Println("Day 09 input:", input)

	if input == "" {
		fmt.Print("Please update the input.txt\n")
		os.Exit(1)
	}

	// Part 1
	fmt.Println("Part 1:", solvePart1(input))

	// Part 2
	// fmt.Println("Part 2:", solvePart2(input))
}

func solvePart1(input string) string {
	disk := input
	if len(input)%2 == 0 {
		disk = input[:len(input)-1]
	}

	temp := make([]int, 0)
	tempWithoutSpaces := make([]int, 0)
	id := 0
	for index, digit := range disk {
		itemToAppend := -1
		if index%2 == 0 {
			itemToAppend = id
			id++
			for i := 0; i < int(digit-'0'); i++ {
				tempWithoutSpaces = append(tempWithoutSpaces, itemToAppend)
			}
		}

		for i := 0; i < int(digit-'0'); i++ {
			temp = append(temp, itemToAppend)
		}
	}

	newTemp := make([]int, 0)
	for index := range len(tempWithoutSpaces) {
		digit := temp[index]

		if digit != -1 {
			newTemp = append(newTemp, digit)
			continue
		}

		for reverseIndex := len(temp) - 1; reverseIndex >= 0; reverseIndex-- {
			if temp[reverseIndex] == -1 {
				continue
			}
			newTemp = append(newTemp, temp[reverseIndex])
			temp = append(temp[:reverseIndex], temp[reverseIndex+1:]...)
			break
		}
	}

	checksum := 0
	for index, digit := range newTemp {
		checksum += index * digit
	}

	return strconv.Itoa(checksum)
}

type File struct {
	Id   int
	Size int
}
type Space struct {
	Size int
}

func solvePart2(input string) string {
	disk := input
	if len(input)%2 == 0 {
		disk = input[:len(input)-1]
	}

	diskMap := []interface{}{}
	id := 0
	for index, digit := range disk {
		if index%2 == 0 {
			diskMap = append(diskMap, File{Id: id, Size: int(digit - '0')})
			id++
		} else {
			diskMap = append(diskMap, Space{Size: int(digit - '0')})
		}
	}

	fmt.Print(diskMap...)

	for reverseIndex := len(diskMap) - 1; reverseIndex > 0; reverseIndex-- {
		if _, ok := diskMap[reverseIndex].(Space); ok {
			continue
		}

		file, ok := diskMap[reverseIndex].(File)
		if !ok {
			fmt.Print("You have done a terrible mistake")
			os.Exit(1)
		}

		fmt.Print("\n " + strconv.Itoa(len(diskMap)) + "\n")
		fmt.Print(diskMap...)

		for index, potentialSpace := range diskMap {
			if index > reverseIndex {
				break
			}
			if space, ok := potentialSpace.(Space); ok {
				if space.Size < file.Size {
					continue
				}

				if space.Size == file.Size {
					diskMap[index] = file
					diskMap[reverseIndex] = space
					break
				}

				space.Size = space.Size - file.Size
				diskMap[reverseIndex] = Space{Size: file.Size}
				if reverseIndex > 0 {
					if prevSpace, ok := diskMap[reverseIndex-1].(Space); ok {
						space.Size = prevSpace.Size + space.Size
						diskMap = append(diskMap[:reverseIndex-1], diskMap[reverseIndex:]...)
					}
				}
				if reverseIndex < len(diskMap)-1 {
					if prevSpace, ok := diskMap[reverseIndex+1].(Space); ok {
						space.Size = prevSpace.Size + space.Size
						diskMap = append(diskMap[:reverseIndex+1], diskMap[reverseIndex+2:]...)
					}
				}

				temp := append([]interface{}{file, space}, diskMap[index+1:]...)
				diskMap = append(diskMap[:index], temp...)
				break
			}
		}
	}
	fmt.Print("\n " + strconv.Itoa(len(diskMap)) + "\n")
	fmt.Print(diskMap...)

	checksum := 0
	for _, potentialFile := range diskMap {
		if file, ok := potentialFile.(File); ok {
			checksum += file.Id * file.Size
		}
	}

	return strconv.Itoa(checksum)
}
