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

		line := record.SingleLine

		var row []string
		empty := true
		for _, char := range line {
			if string(char) != "." {
				empty = false
			}
			row = append(row, string(char))
		}

		if empty {
			emptyRows = append(emptyRows, record.ID)
		}
		grid = append(grid, row)
	}

	maxRow := len(grid)
	maxCol := len(grid[0])

	// fmt.Println("----------------------------")
	// for _, row := range grid {
	// 	fmt.Println(row)
	// }
	// fmt.Println("----------------------------")

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

	// fmt.Println("empty rows", emptyRows)
	// fmt.Println("empty cols", emptyCols)

	expandedGrid := expand(grid, emptyRows, emptyCols, maxRow, maxCol)
	// fmt.Println("----------------------------")
	// for _, row := range expandedGrid {
	// 	fmt.Println(row)
	// }
	// fmt.Println("----------------------------")

	maxRow, maxCol = len(expandedGrid), len(expandedGrid[0])

	solve(expandedGrid, maxRow, maxCol)
}

func solve(grid [][]string, maxRow, maxCol int) {
	var galaxies [][2]int

	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if grid[r][c] == "#" {
				galaxies = append(galaxies, [2]int{r, c})
			}
		}
	}

	// SDFMap := map[int][][]int{}
	// for i, position := range galaxies {
	// 	SDFMap[i] = computeSDF(position[0], position[1], maxRow, maxCol)
	// 	fmt.Println("----------------------------")
	// 	fmt.Println(position)
	// 	for _, row := range SDFMap[i] {
	// 		for _, cell := range row {
	// 			fmt.Printf("%02d ", cell)

	// 		}
	// 		fmt.Printf("\n")
	// 	}
	// 	fmt.Println("----------------------------")
	// }

	total := 0
	var pairs [][2]int
	for i, iPosition := range galaxies {
		for j, jPosition := range galaxies {
			if j >= i {
				continue
			}

			// total += pathfind(position[0], position[1], SDFMap[j])

			delta := int(math.Abs(float64(iPosition[0]) - float64(jPosition[0])))
			delta += int(math.Abs(float64(iPosition[1]) - float64(jPosition[1])))
			total += delta

			pairs = append(pairs, [2]int{i, j})
		}
	}

	fmt.Println(len(pairs))
	fmt.Println(total)
}

func computeSDF(r, c, maxRow, maxCol int) [][]int {
	queue := [][3]int{{r, c, 0}}

	var result [][]int
	for r := 0; r < maxRow; r++ {
		row := make([]int, maxCol)
		for i := 0; i < len(row); i++ {
			row[i] = -1
		}
		result = append(result, row)
	}

	for len(queue) > 0 {
		r, c, dist := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		result[r][c] = dist

		neighbors := getNeighbors(r, c, maxRow, maxCol)
		for _, neighbor := range neighbors {
			if result[neighbor[0]][neighbor[1]] != -1 {
				continue
			}
			queue = append(queue, [3]int{neighbor[0], neighbor[1], dist + 1})
		}
	}
	return result
}

func pathfind(r, c int, SDF [][]int) int {
	return -1
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
