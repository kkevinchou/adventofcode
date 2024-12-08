package main

import (
	"fmt"
	"math"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "sample"
	generator := utils.RecordGenerator(file, "\r\n")

	var grid [][]string
	var start [2]int
	var maxRow, maxCol int

	for {
		record, done := generator()
		if done {
			break
		}

		line := record.Line

		row := []string{}
		for col, character := range line {
			piece := string(character)
			row = append(row, piece)

			if piece == "S" {
				start = [2]int{record.LineNumber, col}
			}
		}
		grid = append(grid, row)
	}

	maxRow = len(grid)
	maxCol = len(grid[0])

	current := [2]int{start[0], start[1]}
	neighborDir := getNeighborDir(grid, current[0], current[1], maxRow, maxCol, [2]int{-100, -100})

	count := 1
	for {
		currentCharacter := grid[current[0]][current[1]]
		if currentCharacter == "S" && count != 1 {
			break
		}

		neighborDir = getNeighborDir(grid, current[0], current[1], maxRow, maxCol, neighborDir)
		current = [2]int{current[0] + neighborDir[0], current[1] + neighborDir[1]}
		count += 1
	}

	fmt.Println(math.Floor(float64(count) / 2))
}

func getNeighborDir(grid [][]string, r, c, maxRow, maxCol int, lastDir [2]int) [2]int {
	var dirToStart [2]int
	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]

		if (dir[0] != 0 && (-dir[0] == lastDir[0])) || (dir[1] != 0 && (-dir[1] == lastDir[1])) {
			continue
		}

		if nr < 0 || nr >= maxRow {
			continue
		}

		if nc < 0 || nc >= maxCol {
			continue
		}

		neighborCharacter := grid[nr][nc]
		if neighborCharacter == "S" {
			dirToStart = [2]int{dir[0], dir[1]}
		}

		currentCharacter := grid[r][c]
		if dir[0] == 1 && dir[1] == 0 {
			if !(currentCharacter == "S" || currentCharacter == "|" || currentCharacter == "7" || currentCharacter == "F") || !(neighborCharacter == "|" || neighborCharacter == "L" || neighborCharacter == "J") {
				continue
			}
		} else if dir[0] == 0 && dir[1] == -1 {
			if !(currentCharacter == "S" || currentCharacter == "-" || currentCharacter == "J" || currentCharacter == "7") || !(neighborCharacter == "-" || neighborCharacter == "L" || neighborCharacter == "F") {
				continue
			}
		} else if dir[0] == -1 && dir[1] == 0 {
			if !(currentCharacter == "S" || currentCharacter == "|" || currentCharacter == "J" || currentCharacter == "L") || !(neighborCharacter == "|" || neighborCharacter == "7" || neighborCharacter == "F") {
				continue
			}
		} else if dir[0] == 0 && dir[1] == 1 {
			if !(currentCharacter == "S" || currentCharacter == "-" || currentCharacter == "F" || currentCharacter == "L") || !(neighborCharacter == "-" || neighborCharacter == "J" || neighborCharacter == "7") {
				continue
			}
		}

		return dir
	}

	return dirToStart
}

var dirs [][2]int = [][2]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

var dirToLeft map[[2]int][2]int = map[[2]int][2]int{
	[2]int{1, 0}:  [2]int{0, 1},
	[2]int{0, -1}: [2]int{1, 0},
	[2]int{-1, 0}: [2]int{0, -1},
	[2]int{0, 1}:  [2]int{-1, 0},
}
