// Solution to Advent of Code 2024 Day 02 - Red-Nosed Reports

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var safeReports int = 0
	var safeReportsWithDampener int = 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Can't open the input file")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF && len(line) == 0 {
			fmt.Println("Reached EOF")
			break
		}

		lineArray := strings.Fields(line)
		safeReports += part1(lineArray)
		safeReportsWithDampener += part2(lineArray)
	}
	fmt.Println("The number of safe reports is:", safeReports)
	fmt.Println("The number of safe reports with dampener is:", safeReportsWithDampener)

}

func part1(lineArray []string) int {
	fmt.Println("This is the array of numbers:", lineArray)
	if isSafe(lineArray) {
		fmt.Println("This whole line is safe")
		return 1
	}
	return 0
}

func part2(lineArray []string) int {
	var counter = 0
	fmt.Println("This is the array of numbers:", lineArray)

	if !isSafe(lineArray) {
		for i := 0; i < len(lineArray); i++ {

			// Create a new independent slice
			newLineArray := make([]string, 0, len(lineArray)-1)
			newLineArray = append(newLineArray, lineArray[:i]...)
			newLineArray = append(newLineArray, lineArray[i+1:]...)

			fmt.Println("This is the new array of numbers:", newLineArray)

			if isSafe(newLineArray) {
				fmt.Println("This whole line is safe with one modification")
				counter++
				break
			}
		}
	}

	if isSafe(lineArray) {
		counter++
		fmt.Println("This whole line is safe as is")
		return counter
	}

	return counter
}

func isSafe(lineArray []string) bool {
	value0, _ := strconv.Atoi(lineArray[0])
	value01, _ := strconv.Atoi(lineArray[1])
	increasing := value0 < value01
	fmt.Println("Line is increasing:", increasing)

	var safe bool = true
	for i := 1; i <= len(lineArray)-1; i++ {
		value1, _ := strconv.Atoi(lineArray[i-1])
		value2, _ := strconv.Atoi(lineArray[i])
		if value1 < value2 {
			if increasing {
				substract := math.Abs(float64(value1) - float64(value2))
				if substract > 3 {
					fmt.Println("substract > 3, increasing line not safe")
					safe = false
				}
			}

			if !increasing {
				fmt.Println("this line should be decreasing instead")
				safe = false
			}
		}
		if value1 > value2 {
			if !increasing {
				substract := math.Abs(float64(value1) - float64(value2))
				if substract > 3 {
					fmt.Println("substract > 3, decreasing line not safe")
					safe = false
				}
			}

			if increasing {
				fmt.Println("this line should be increasing instead")
				safe = false
			}
		}
		if value1 == value2 {
			fmt.Println("the values should not be same")
			safe = false
		}
	}
	return safe
}
