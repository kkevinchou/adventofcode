package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

type Coordinate [2]int

var dirs [][2]int = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func main() {
	grid, rCount, cCount := utils.ParseGrid(file)

	var result int
	for r := range rCount {
		for c := range cCount {
			if grid[r][c] == "0" {
				endPositions := map[Coordinate]int{}
				solve(grid, r, c, rCount, cCount, 0, endPositions)
				for _, numTrails := range endPositions {
					result += numTrails
				}
			}
		}
	}
	fmt.Println(result)
}

func solve(grid [][]string, r, c, rCount, cCount, num int, endPositions map[Coordinate]int) {
	if grid[r][c] != fmt.Sprintf("%d", num) {
		return
	}
	if num == 9 {
		endPositions[Coordinate{r, c}]++
		return
	}

	for dir := range 4 {
		newR, newC := r+dirs[dir][0], c+dirs[dir][1]
		if newR < 0 || newR >= rCount || newC < 0 || newC >= cCount {
			continue
		}

		solve(grid, newR, newC, rCount, cCount, num+1, endPositions)
	}
}
