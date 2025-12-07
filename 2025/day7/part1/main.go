package main

import (
	"fmt"
	"os"
	"strings"
)

var file string = "input.txt"

var memo map[string]int

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var startR, startC int

	grid := [][]string{}
	lines := strings.Split(string(data), "\r\n")
	for r, line := range lines {
		grid = append(grid, make([]string, len(line)))
		for c, char := range line {
			grid[r][c] = string(char)
			if string(char) == "S" {
				startR = r
				startC = c
			}
		}
	}

	rCount := len(grid)
	cCount := len(grid[0])
	memo = map[string]int{}

	fmt.Println(step(startR, startC, rCount, cCount, grid))
}

func step(r, c, rCount, cCount int, grid [][]string) int {
	if _, ok := memo[fmt.Sprintf("%d_%d", r, c)]; ok {
		return 0
	}

	if r >= rCount {
		memo[fmt.Sprintf("%d_%d", r, c)] = 0
		return 0
	}

	if grid[r][c] == "S" || grid[r][c] == "." {
		result := step(r+1, c, rCount, cCount, grid)
		memo[fmt.Sprintf("%d_%d", r, c)] = result
		return result
	}

	// split
	result := 1 + step(r+1, c-1, rCount, cCount, grid) + step(r+1, c+1, rCount, cCount, grid)
	memo[fmt.Sprintf("%d_%d", r, c)] = result
	return result
}
