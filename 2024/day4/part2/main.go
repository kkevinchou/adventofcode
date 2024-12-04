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
			if r+1 < rCount && c+1 < cCount && grid[r+1][c+1] == 'A' {
				if grid[r][c] == 'M' || grid[r][c] == 'S' {
					// top left
					s1, m1 := countMS(r, c, grid)
					// bottom right
					s2, m2 := countMS(r+2, c+2, grid)

					if s1+m1+s2+m2 != 2 {
						continue
					}

					if !((s1 + s2) == (m1 + m2)) {
						continue
					}

					// top right
					s1, m1 = countMS(r, c+2, grid)
					// bottom left
					s2, m2 = countMS(r+2, c, grid)

					if s1+m1+s2+m2 != 2 {
						continue
					}

					if !((s1 + s2) == (m1 + m2)) {
						continue
					}

					total += 1
				}
			}
		}
	}

	fmt.Println(total)
}

func countMS(r, c int, grid [][]rune) (int, int) {
	// bounds checking
	if r < 0 || r >= len(grid) {
		return 0, 0
	}

	if c < 0 || c >= len(grid[0]) {
		return 0, 0
	}

	if grid[r][c] == 'M' {
		return 0, 1
	} else if grid[r][c] == 'S' {
		return 1, 0
	}

	return 0, 0
}
