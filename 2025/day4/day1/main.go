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

	lines := strings.Split(string(data), "\r\n")
	// var grid [][]string
	grid := [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	rCount := len(grid)
	cCount := len(grid[0])

	var result int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] != "@" {
				continue
			}
			count := check(grid, r, c, rCount, cCount)
			if count < 4 {
				result++
			}
		}
	}

	fmt.Println(result)
}

var dirs []int = []int{-1, 0, 1}

func check(grid [][]string, r, c, rCount, cCount int) int {
	var count int
	for _, rDir := range dirs {
		for _, cDir := range dirs {
			if rDir == 0 && cDir == 0 {
				continue
			}

			nr := r + rDir
			nc := c + cDir

			if nr < 0 || nr >= rCount {
				continue
			}
			if nc < 0 || nc >= cCount {
				continue
			}

			if grid[nr][nc] == "@" {
				count++
			}
		}
	}
	return count
}
