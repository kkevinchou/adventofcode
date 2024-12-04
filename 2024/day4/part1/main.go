package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	var grid [][]rune

	for record := range utils.Records("input", "\n") {
		grid = append(grid, make([]rune, len(record.SingleLine)))
		for c, char := range record.SingleLine {
			grid[record.ID][c] = char
		}
	}

	rCount := len(grid)
	cCount := len(grid[0])

	var total int
	for r := range rCount {
		for c := range cCount {
			total += search(r, c, grid, 'X', -1, -1)
		}
	}

	// total += search(0, 4, grid, 'X', -1, -1)

	fmt.Println(total)
}

var dirs []int = []int{-1, 0, 1}

func search(r, c int, grid [][]rune, match rune, rDir, cDir int) int {
	// bounds checking
	if r < 0 || r >= len(grid) {
		return 0
	}

	if c < 0 || c >= len(grid[0]) {
		return 0
	}

	if grid[r][c] != match {
		return 0
	}

	if match == 'S' {
		return 1
	}

	// we got a match

	var nextCharacter rune

	if match == 'X' {
		nextCharacter = 'M'
	} else if match == 'M' {
		nextCharacter = 'A'
	} else if match == 'A' {
		nextCharacter = 'S'
	} else {
		panic("WAT " + string(match))
	}

	if match == 'X' {
		var total int
		for _, rDir := range dirs {
			for _, cDir := range dirs {
				total += search(r+rDir, c+cDir, grid, nextCharacter, rDir, cDir)
			}
		}
		return total
	}

	return search(r+rDir, c+cDir, grid, nextCharacter, rDir, cDir)
}
