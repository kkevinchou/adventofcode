package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var grid [][]string
	var start [2]int
	var maxRow, maxCol int

	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine

		row := []string{}
		for col, character := range line {
			piece := string(character)
			row = append(row, piece)

			if piece == "S" {
				start = [2]int{record.ID, col}
			}
		}
		grid = append(grid, row)
	}

	maxRow = len(grid)
	maxCol = len(grid[0])

	inAndOutGrid := [][]int{}
	for i := 0; i < maxRow; i++ {
		inAndOutGrid = append(inAndOutGrid, make([]int, maxCol))
	}

	current := [2]int{start[0], start[1]}
	var neighborDir [2]int = [2]int{-100, -100}

	var path [][2]int

	count := 1
	for {
		inAndOutGrid[current[0]][current[1]] = 5
		currentCharacter := grid[current[0]][current[1]]
		if currentCharacter == "S" && count != 1 {
			break
		}

		neighborDir = getNeighborDir(grid, current[0], current[1], maxRow, maxCol, neighborDir)
		path = append(path, neighborDir)
		current = [2]int{current[0] + neighborDir[0], current[1] + neighborDir[1]}
		count += 1
	}

	current = start
	for _, dir := range path {
		current[0] = current[0] + dir[0]
		current[1] = current[1] + dir[1]

		leftDir := dirToLeft[dir]
		fill(inAndOutGrid, current[0]+leftDir[0], current[1]+leftDir[1], maxRow, maxCol, 1)
		fill(inAndOutGrid, current[0]-dir[0]+leftDir[0], current[1]-dir[1]+leftDir[1], maxRow, maxCol, 1)
		rightDir := dirToRight[dir]
		fill(inAndOutGrid, current[0]+rightDir[0], current[1]+rightDir[1], maxRow, maxCol, 2)
		fill(inAndOutGrid, current[0]-dir[0]+rightDir[0], current[1]-dir[1]+rightDir[1], maxRow, maxCol, 2)
	}

	leftCount := 0
	rightCount := 0
	zeroCount := 0
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if inAndOutGrid[r][c] == 1 {
				leftCount += 1
			}
			if inAndOutGrid[r][c] == 2 {
				rightCount += 1
			}

			if inAndOutGrid[r][c] == 0 {
				zeroCount += 1
			}
		}
	}

	total := len(inAndOutGrid) * len(inAndOutGrid[0])
	fmt.Println("TOTAL", total)
	fmt.Println("PATH SIZE", len(path))
	fmt.Println("LEFT SIZE", leftCount)
	fmt.Println("RIGHT SIZE", rightCount)
	fmt.Println("ZERO COUNT", zeroCount)
}

func fill(inAndOutGrid [][]int, r, c int, maxRow, maxCol int, fillNum int) {
	queue := [][2]int{{r, c}}

	if r < 0 || r >= maxRow {
		return
	}

	if c < 0 || c >= maxCol {
		return
	}

	for len(queue) > 0 {
		r, c := queue[0][0], queue[0][1]
		queue = queue[1:]

		if inAndOutGrid[r][c] != 5 && inAndOutGrid[r][c] != 0 && inAndOutGrid[r][c] != fillNum {
			panic("WAT")
		}

		if inAndOutGrid[r][c] == 0 {
			inAndOutGrid[r][c] = fillNum

			neighbors := getNeighbors(inAndOutGrid, r, c, maxRow, maxCol)
			for _, neighbor := range neighbors {
				nr, nc := neighbor[0], neighbor[1]
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
}

func getNeighbors(inAndOutGrid [][]int, r, c, maxRow, maxCol int) [][2]int {
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

var dirToRight map[[2]int][2]int = map[[2]int][2]int{
	[2]int{1, 0}:  [2]int{0, -1},
	[2]int{0, -1}: [2]int{-1, 0},
	[2]int{-1, 0}: [2]int{0, 1},
	[2]int{0, 1}:  [2]int{1, 0},
}
