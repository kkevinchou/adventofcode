package main

import (
	"fmt"
	"os"
	"strings"
)

var file string = "input.txt"

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var startC int
	grid := [][]string{}
	lines := strings.Split(string(data), "\r\n")

	for r, line := range lines {
		grid = append(grid, make([]string, len(line)))
		for c, char := range line {
			grid[r][c] = string(char)
			if string(char) == "S" {
				startC = c
			}
		}
	}
	rCount := len(grid)
	cCount := len(grid[0])

	totals := make([]int, cCount)
	totals[startC] = 1

	for r := range rCount {
		for c := range cCount {
			if grid[r][c] == "^" {
				totals[c-1] += totals[c]
				totals[c+1] += totals[c]
				totals[c] = 0
			}
		}
	}

	var result int
	for i := 0; i < len(totals); i++ {
		result += totals[i]
	}
	fmt.Println(result)
}
