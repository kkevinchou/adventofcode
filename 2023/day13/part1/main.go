package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n\r\n")

	var total int
	for {
		record, done := generator()
		if done {
			break
		}

		var grid [][]string

		lines := record.Lines
		for _, line := range lines {
			row := utils.StringToStringSlice(line)
			grid = append(grid, row)
		}

		maxRow := len(grid)
		maxCol := len(grid[0])
		result := solve(grid, maxRow, maxCol)

		total += result
	}

	fmt.Println(total)
}

func solve(grid [][]string, maxRow, maxCol int) int {
	// vertical case

	mirrorFound := false
	mirrorIndex := -1
	for i := 0; i < maxCol-1; i++ {
		for j := 0; j < maxCol; j++ {
			leftIndex := i - j
			rightIndex := i + j + 1

			if leftIndex < 0 {
				mirrorFound = true
				break
			}

			if rightIndex >= maxCol {
				mirrorFound = true
				break
			}

			leftColumn := column(grid, leftIndex, maxRow, maxCol)
			rightColumn := column(grid, rightIndex, maxRow, maxCol)

			if !equals(leftColumn, rightColumn) {
				break
			}
		}

		if mirrorFound {
			mirrorIndex = i
			break
		}
	}

	if mirrorFound {
		return mirrorIndex + 1
	}

	// horizontal case

	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxRow; j++ {
			upIndex := i - j
			downIndex := i + j + 1

			if upIndex < 0 {
				mirrorFound = true
				break
			}

			if downIndex >= maxRow {
				mirrorFound = true
				break
			}

			upRow := row(grid, upIndex)
			downRow := row(grid, downIndex)

			if !equals(upRow, downRow) {
				break
			}
		}

		if mirrorFound {
			mirrorIndex = i
			break
		}
	}

	if mirrorFound {
		return (mirrorIndex + 1) * 100
	}

	panic("WAT DIDNT FIND A MIRROR")
}

func equals(l1 []string, l2 []string) bool {
	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

func column(grid [][]string, index int, maxRow, maxCol int) []string {
	var result []string

	for r := 0; r < maxRow; r++ {
		result = append(result, grid[r][index])
	}

	return result
}

func row(grid [][]string, index int) []string {
	return grid[index]
}
