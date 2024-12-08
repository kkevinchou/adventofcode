package main

import (
	"fmt"
	"math"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var grid [][]string

	var emptyRows []int
	var emptyCols []int

	for {
		record, done := generator()
		if done {
			break
		}

		line := record.Line

		var row []string
		empty := true
		for _, char := range line {
			if string(char) != "." {
				empty = false
			}
			row = append(row, string(char))
		}

		if empty {
			emptyRows = append(emptyRows, record.LineNumber)
		}
		grid = append(grid, row)
	}

	maxRow := len(grid)
	maxCol := len(grid[0])

	for c := 0; c < maxCol; c++ {
		empty := true
		for r := 0; r < maxRow; r++ {
			if grid[r][c] != "." {
				empty = false
				break
			}
		}

		if empty {
			emptyCols = append(emptyCols, c)
		}
	}

	fmt.Println("empty rows", emptyRows)
	fmt.Println("empty cols", emptyCols)

	// expandedGrid := expand(grid, emptyRows, emptyCols, maxRow, maxCol)
	// fmt.Println("----------------------------")
	// for _, row := range expandedGrid {
	// 	fmt.Println(row)
	// }
	// fmt.Println("----------------------------")

	// maxRow, maxCol = len(expandedGrid), len(expandedGrid[0])

	solve(grid, maxRow, maxCol, emptyRows, emptyCols)
}

func solve(grid [][]string, maxRow, maxCol int, emptyRows, emptyCols []int) {
	var galaxies [][2]int

	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if grid[r][c] == "#" {
				galaxies = append(galaxies, [2]int{r, c})
			}
		}
	}

	fmt.Println("NUM GALAXIES", len(galaxies))

	total := 0
	pairCount := 91378
	current := 0
	for i, iPosition := range galaxies {
		for j, jPosition := range galaxies {
			if j >= i {
				continue
			}

			minRow := iPosition[0]
			maxRow := iPosition[0]

			if jPosition[0] < minRow {
				minRow = jPosition[0]
			}
			if jPosition[0] > maxRow {
				maxRow = jPosition[0]
			}

			minCol := iPosition[1]
			maxCol := iPosition[1]

			if jPosition[1] < minCol {
				minCol = jPosition[1]
			}
			if jPosition[1] > maxCol {
				maxCol = jPosition[1]
			}

			var hit int
			for _, emptyRow := range emptyRows {
				if emptyRow > minRow && emptyRow < maxRow {
					hit += 1
				}
			}
			for _, emptyCol := range emptyCols {
				if emptyCol > minCol && emptyCol < maxCol {
					hit += 1
				}
			}

			total += hit*1000000 - hit

			delta := int(math.Abs(float64(iPosition[0]) - float64(jPosition[0])))
			delta += int(math.Abs(float64(iPosition[1]) - float64(jPosition[1])))
			total += delta

			current += 1
			fmt.Printf("%d / %d\n", current, pairCount)
		}
	}

	fmt.Println(total)
}

func getNeighbors(r, c, maxRow, maxCol int) [][2]int {
	var result [][2]int
	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]
		if nr < 0 || nr >= maxRow {
			continue
		}

		if nc < 0 || nc >= maxCol {
			continue
		}

		result = append(result, [2]int{nr, nc})
	}
	return result
}

var dirs [][2]int = [][2]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

func expand(grid [][]string, emptyRows, emptyCols []int, maxRow, maxCol int) [][]string {
	var expandedGrid [][]string

	emptyRowsMap := utils.MapFromSlice[int](emptyRows)
	emptyColsMap := utils.MapFromSlice[int](emptyCols)

	for r := 0; r < maxRow; r++ {
		if _, ok := emptyRowsMap[r]; ok {
			var newRow []string
			for i := 0; i < maxCol+len(emptyCols); i++ {
				newRow = append(newRow, ".")
			}
			expandedGrid = append(expandedGrid, newRow)
		}

		var newRow []string
		for c := 0; c < maxCol; c++ {
			if _, ok := emptyColsMap[c]; ok {
				newRow = append(newRow, ".")
			}
			newRow = append(newRow, grid[r][c])
		}
		expandedGrid = append(expandedGrid, newRow)
	}
	return expandedGrid
}
