package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

// options
// 1 - use a stack of positions - LUL
// 2 - scan all spaces between blockages and allocate rocks

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var grid [][]string
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.Line
		row := utils.StringToStringSlice(line)
		grid = append(grid, row)
	}

	maxRow := len(grid)
	maxCol := len(grid[0])

	var columns [][]string
	for c := 0; c < maxCol; c++ {
		col := make([]string, maxRow)
		columns = append(columns, col)

		for r := 0; r < maxRow; r++ {
			columns[c][r] = grid[r][c]
		}
	}

	for _, c := range columns {
		fmt.Println(c)
	}

	var score int
	for colIndex, column := range columns {
		fmt.Println(colIndex, "------------------")
		start := -1
		stoneCount := 0

		if colIndex == 1 {
			fmt.Println("HI")
		}

		for i := 0; i < maxRow; i++ {
			if start == -1 && (column[i] == "O" || column[i] == ".") {
				start = i
			}

			if column[i] == "O" {
				stoneCount += 1
			}

			if column[i] == "#" || i == maxRow-1 {
				for j := 0; j < stoneCount; j++ {
					delta := (maxRow - (start + j))
					fmt.Println("adding", delta)
					score += delta
				}
				start = -1
				stoneCount = 0
			}
		}
	}

	fmt.Println(score)
}
