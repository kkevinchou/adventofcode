package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

func main() {
	grid, rCount, cCount := utils.ParseGrid(file)

	seen1 := map[[2]int]bool{}
	seen2 := map[[2]int]bool{}

	var result int
	for r := range rCount {
		for c := range cCount {
			area := floodFillArea(grid, r, c, rCount, cCount, grid[r][c], seen1)
			perimeter := floodFillPerimeter(grid, r, c, rCount, cCount, grid[r][c], seen2)
			result += area * perimeter
			// fmt.Println(grid[r][c], "| AREA:", area, "PERIMETER:", perimeter)
		}
	}
	fmt.Println(result)
}

func floodFillArea(grid [][]string, r, c, rCount, cCount int, plot string, seen map[[2]int]bool) int {
	if _, ok := seen[[2]int{r, c}]; ok {
		return 0
	}

	if grid[r][c] != plot {
		return 0
	}

	seen[[2]int{r, c}] = true

	result := 1

	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]
		if nr < 0 || nr >= rCount || nc < 0 || nc >= cCount {
			continue
		}

		result += floodFillArea(grid, nr, nc, rCount, cCount, plot, seen)
	}
	return result
}

func floodFillPerimeter(grid [][]string, r, c, rCount, cCount int, plot string, seen map[[2]int]bool) int {
	if _, ok := seen[[2]int{r, c}]; ok {
		return 0
	}

	if grid[r][c] != plot {
		return 0
	}

	seen[[2]int{r, c}] = true

	result := 0

	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]
		if nr < 0 || nr >= rCount || nc < 0 || nc >= cCount {
			result++
			continue
		}

		if grid[nr][nc] != plot {
			result++
		}

		result += floodFillPerimeter(grid, nr, nc, rCount, cCount, plot, seen)
	}
	return result
}

var dirs [][2]int = [][2]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}
