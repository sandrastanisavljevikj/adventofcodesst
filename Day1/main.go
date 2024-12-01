// Solution to Advent of Code 2024 Day 01 - Historian Hysteria

package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var array1 []float64
	var array2 []float64

	file, err := os.Open("input.txt")
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

		numbersInLine := strings.Fields(line)
		num1, err := strconv.Atoi(numbersInLine[0])
		if err != nil {
			fmt.Println("Error converting first num to int ", num1)
			return
		}
		array1 = append(array1, float64(num1))

		num2, err := strconv.Atoi(numbersInLine[1])
		if err != nil {
			fmt.Println("Error converting second num to int ", num2)
			return
		}
		array2 = append(array2, float64(num2))
	}

	part1(array1, array2)
	part2(array1, array2)
}

func part1(array1 []float64, array2 []float64) {
	sort.Float64s(array1)
	sort.Float64s(array2)

	distance := 0

	for i, element := range array1 {
		substract := element - array2[i]
		substractAbs := math.Abs(substract)
		distance += int(substractAbs)
	}
	fmt.Println("Distance", distance)
}

func part2(array1 []float64, array2 []float64) {
	mapOfArray2 := make(map[float64]int)
	for _, element := range array2 {
		mapOfArray2[element] += 1
		fmt.Println("mapOfArray2[element]", mapOfArray2[element])
		fmt.Println("element", element)
	}
	var similarity float64 = 0

	for _, currentNumber := range array1 {
		value, exists := mapOfArray2[currentNumber]
		if exists {
			similarity += currentNumber * float64(value)
		}
	}

	fmt.Println("Similarity", similarity)

}
