// Solution to Advent of Code 2024 Day 04 - Ceres Search

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := readFileToGrid("input.txt")
	part1Result := part1(grid)
	fmt.Println("Part1:", part1Result)
	part2Result := part2(grid)
	fmt.Println("Part2:", part2Result)
}

func readFileToGrid(filepath string) [][]rune {
	file, _ := os.Open(filepath)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}
	return grid
}

func part1(grid [][]rune) int {
	keyword := "XMAS"

	var directions = [][]int{
		{0, 1},   // right
		{0, -1},  // left
		{1, 0},   // up
		{-1, 0},  // down
		{1, 1},   // down right
		{-1, 1},  // up right
		{1, -1},  // down left
		{-1, -1}, // down right
	}

	return searchGrid(grid, keyword, directions)
}

func part2(grid [][]rune) int {
	rows := len(grid)       //x
	columns := len(grid[0]) //y
	var result int

	for x := 1; x < rows-1; x++ {
		for y := 1; y < columns-1; y++ {
			if grid[x][y] == rune("A"[0]) {
				if (grid[x-1][y-1] == rune("M"[0]) && grid[x+1][y+1] == rune("S"[0]) ||
					grid[x-1][y-1] == rune("S"[0]) && grid[x+1][y+1] == rune("M"[0])) &&
					(grid[x-1][y+1] == rune("M"[0]) && grid[x+1][y-1] == rune("S"[0]) ||
						grid[x-1][y+1] == rune("S"[0]) && grid[x+1][y-1] == rune("M"[0])) {
					result++
				}
			}
		}
	}
	return result
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
}

func searchGrid(grid [][]rune, keyword string, directions [][]int) int {
	rows := len(grid)       //x
	columns := len(grid[0]) //y
	var result int

	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			if grid[x][y] == rune(keyword[0]) {
				for _, direction := range directions {
					dx, dy := direction[0], direction[1]
					if searchInDirection(grid, keyword, x, y, dx, dy) {
						result++
					}
				}

			}
		}
	}
	return result
}

func searchInDirection(grid [][]rune, keyword string, x, y, dx, dy int) bool {
	for i := 0; i < len(keyword); i++ {
		currentX := i*dx + x
		currentY := i*dy + y

		if !isValid(currentX, currentY, len(grid), len(grid[0])) {
			return false
		}

		if grid[currentX][currentY] != rune(keyword[i]) {
			return false
		}
	}
	return true
}
