package main

import (
	"fmt"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

var dirs [][2]int = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func main() {
	start := time.Now()
	var grid [][]string

	var startR, startC int
	var dirIndex = 0

	for record := range utils.Records("input", "\n") {
		grid = append(grid, make([]string, len(record.Line)))
		for c, char := range record.Line {
			grid[record.LineNumber][c] = string(char)

			if string(char) == "^" {
				startR = record.LineNumber
				startC = c
			}
		}
	}
	rCount := len(grid)
	cCount := len(grid[0])

	result := 1
	grid[startR][startC] = "X"

	r, c := startR, startC

	for {
		nextR, nextC := r+dirs[dirIndex][0], c+dirs[dirIndex][1]
		if nextR < 0 || nextR >= rCount || nextC < 0 || nextC >= cCount {
			break
		}

		if grid[nextR][nextC] == "#" {
			dirIndex = (dirIndex + 1) % len(dirs)
			continue
		}

		r, c = nextR, nextC

		if grid[r][c] != "X" {
			grid[r][c] = "X"
			result++
		}
	}

	fmt.Println(result)
	fmt.Println(time.Since(start))
}
