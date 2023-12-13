package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

// options
// 1 - brute force every smudge
// 2 - early opt out of rows / columns that don't have the other elements of the row / column matching
// 3 - XOR rows, place mirror and check

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

		result := solve(grid, maxRow, maxCol, -9999)

		foundNewMirror := false

		for i := 0; i < maxRow; i++ {
			for j := 0; j < maxCol; j++ {
				value := grid[i][j]

				newValue := "#"
				if value == "#" {
					newValue = "."
				}

				grid[i][j] = newValue
				newResult := solve(grid, maxRow, maxCol, result)
				if newResult != -1 && newResult != result {
					result = newResult
					foundNewMirror = true
				}
				grid[i][j] = value

				if foundNewMirror {
					break
				}
			}
			if foundNewMirror {
				break
			}
		}

		if foundNewMirror == false {
			panic("WAT")
		}

		total += result
	}

	fmt.Println(total)
}

func solve(grid [][]string, maxRow, maxCol int, oldResult int) int {
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
			if i+1 != oldResult {
				mirrorIndex = i
				break
			} else {
				mirrorFound = false
			}
		}
	}

	if mirrorFound {
		return mirrorIndex + 1
	}

	// horizontal case

	for i := 0; i < maxRow-1; i++ {
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
			if (i+1)*100 != oldResult {
				mirrorIndex = i
				break
			} else {
				mirrorFound = false
			}
		}
	}

	if mirrorFound {
		return (mirrorIndex + 1) * 100
	}
	return -1
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
