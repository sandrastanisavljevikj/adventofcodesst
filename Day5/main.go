package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Maps in Go are reference types. They need to be
	// initialized using the make function or a map literal before they can be used.
	mapOfRules := make(map[int][]int)
	updates := [][]string{}
	file, err := os.Open("rule-input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF && len(line) == 0 {
			break
		}

		numbersInLine := strings.Split(line, "|")
		num1, err := strconv.Atoi(strings.TrimSpace(numbersInLine[0]))
		if err != nil {
			fmt.Println("Error converting first num to int ", numbersInLine[0])
			return
		}
		num2, err := strconv.Atoi(strings.TrimSpace(numbersInLine[1]))
		if err != nil {
			fmt.Println("Error converting second num to int ", numbersInLine[1])
			return
		}

		currentArrayOfRules := mapOfRules[num1]
		newArray := append(currentArrayOfRules, num2)
		mapOfRules[num1] = newArray
	}

	fileUpdates, err := os.Open("updates-input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fileUpdates.Close()
	readerForUpdates := bufio.NewReader(fileUpdates)
	for {
		line, err := readerForUpdates.ReadString('\n')
		if err == io.EOF && len(line) == 0 {
			break
		}

		updates = append(updates, strings.Split(line, ","))
	}

	result1, wrongUpdates := part1(mapOfRules, updates)
	fmt.Println("FirstResult:", result1)

	result2 := part2(mapOfRules, wrongUpdates)
	fmt.Println("SecondResult:", result2)
}

func part1(mapOfRules map[int][]int, updates [][]string) (int, [][]string) {
	result := 0
	wrongUpdates := [][]string{}

	for _, singleUpdateList := range updates {
		goodList := true
		// check all elements in one row

		for i := 0; i < len(singleUpdateList); i++ {
			element := singleUpdateList[i]
			intElement, _ := strconv.Atoi(element)
			rules := mapOfRules[intElement]
			// the current element has to be before all this elements in rules slice
			if rules == nil {
				fmt.Println("no rules for element", element)
				continue
			}
			// if you see a number appearing before the element
			// but it was in the list of rules, its wrong
			for y := 0; y < i; y++ {
				value, _ := strconv.Atoi(singleUpdateList[y])
				if elementExistsInSlice(rules, value) {
					goodList = false
					break
				}
			}
			if !goodList {
				wrongUpdates = append(wrongUpdates, singleUpdateList)
				break
			}
			continue
		}

		if goodList {
			middleElement := getMiddleElement(singleUpdateList)
			result += middleElement
		}
	}
	return result, wrongUpdates
}

func part2(mapOfRules map[int][]int, wrongUpdates [][]string) int {
	result := 0
	return result
}

func getMiddleElement(singleUpdateList []string) int {
	middleIndex := len(singleUpdateList) / 2
	middleElement, _ := strconv.Atoi(singleUpdateList[middleIndex])
	return middleElement
}

func elementExistsInSlice(slice []int, target int) bool {
	for _, value := range slice {
		if value == target {
			return true
		}
	}
	return false
}

// move repositions an item in a slice from one index to another.
func move[T any](slice []T, fromIndex, toIndex int) []T {
	// Ensure indices are within bounds
	if fromIndex < 0 || fromIndex >= len(slice) || toIndex < 0 || toIndex >= len(slice) {
		fmt.Println("Invalid indices")
		return slice
	}

	// If the item is already at the desired position, return the original slice
	if fromIndex == toIndex {
		return slice
	}

	// Extract the item to be moved
	item := slice[fromIndex]

	// Remove the item from the slice
	slice = append(slice[:fromIndex], slice[fromIndex+1:]...)

	// Insert the item at the new position
	if toIndex > fromIndex {
		// If moving forward, adjust the index due to removal
		toIndex--
	}
	slice = append(slice[:toIndex], append([]T{item}, slice[toIndex:]...)...)

	return slice
}
