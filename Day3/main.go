// Solution to Advent of Code 2024 Day 03 - Mull It Over

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	part1()
	part2()

}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't open the input file")
	}

	defer file.Close()
	var result int
	scanner := bufio.NewScanner(file)
	const maxBufferSize = 1 * 1024 * 1024 // 10 MB
	buffer := make([]byte, maxBufferSize)
	scanner.Buffer(buffer, maxBufferSize)

	// Compile the regex
	pattern := `mul\(\d+,\d+\)`

	for scanner.Scan() {
		// Compile the regex
		re, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Println("Could not compile the regex")
		}
		line := scanner.Text()
		//The -1 argument specifies no limit on the number of matches.
		matches := re.FindAllString(line, -1)
		result += getNumbersAndMultiply(matches)
		// Check for scanning errors
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
		}
	}
	fmt.Println("Final result:", result)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't open the input file")
	}

	defer file.Close()
	var result int
	var fullInputString string
	pattern := `mul\(\d+,\d+\)|don't\(\)|do\(\)`
	regex, _ := regexp.Compile(pattern)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fullInputString += scanner.Text()
	}
	fmt.Println(fullInputString)
	matches := regex.FindAllString(fullInputString, -1)
	fmt.Println(matches)
	enabled := true
	for _, element := range matches {
		if element == "don't()" {
			enabled = false
			continue
		}
		if element == "do()" {
			enabled = true
			continue
		}
		if enabled {
			result += getNumbersAndMultiply([]string{element})
		}
	}

	fmt.Println("Final result2:", result)

}

func getNumbersAndMultiply(matches []string) int {
	result := 0
	for _, element := range matches {
		// Regular expression to match only the numbers in mul(x,y)
		re := regexp.MustCompile(`\d+`)

		// Find all matches
		matches := re.FindAllString(element, -1)

		// Convert matches to integers
		var numbers []int
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err == nil {
				numbers = append(numbers, num)
			}
		}
		result += numbers[0] * numbers[1]
	}
	return result
}
