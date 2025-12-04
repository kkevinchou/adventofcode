package main

import (
	"fmt"
	"os"
	"strings"
)

var file string = "input.txt"

var dirs []int = []int{-1, 0, 1}

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")
	grid := [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(line, ""))
	}

	rCount := len(grid)
	cCount := len(grid[0])
	var result int

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			result += fix(grid, r, c, rCount, cCount)
		}
	}

	fmt.Println(result)
}

func fix(grid [][]string, r, c, rCount, cCount int) int {
	var count int

	if grid[r][c] != "@" {
		return 0
	}

	if check(grid, r, c, rCount, cCount) < 4 {
		grid[r][c] = "."
		count++
	} else {
		return 0
	}

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

			count += fix(grid, nr, nc, rCount, cCount)
		}
	}

	return count
}

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
