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
	recordGenerator := utils.RecordGenerator(file, "\r\n")

	var grid [][]string
	for {
		record, done := recordGenerator()
		if done {
			break
		}

		line := record.SingleLine
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

	generatorGenerators := [4]GeneratorGenerator{north, west, south, east}

	cycleCount := 1000000000

	var cacheKeys []string

	cycleStartKey := ""
	cycleStartIndex := 0
	cycleOffset := 0

	cache := map[string][][]string{}
	for i := 0; i < cycleCount; i++ {
		cacheHitCount := 0

		cacheKeys = nil

		for generatorIndex, generatorGenerator := range generatorGenerators {
			generator := generatorGenerator(maxRow, maxCol)
			key := genKey(grid, maxRow, maxCol)
			key = fmt.Sprintf("%d|%s", generatorIndex, key)
			if cachedGrid, ok := cache[key]; ok {
				copy(cachedGrid, grid, maxRow, maxCol)
				cacheHitCount += 1
				cacheKeys = append(cacheKeys, key)
			} else {
				tilt(generator, grid, maxRow, maxCol)

				newGrid := make([][]string, maxRow)
				for r := 0; r < maxRow; r++ {
					newGrid[r] = make([]string, maxCol)
				}

				copy(grid, newGrid, maxRow, maxCol)
				cache[key] = newGrid
			}
		}

		if cacheHitCount == 4 {
			if cycleOffset == 0 {
				if cycleStartKey == "" {
					cycleStartKey = cacheKeys[3]
					cycleStartIndex = i
				} else if cycleStartKey == cacheKeys[3] {
					cycleOffset = i - cycleStartIndex
				}
			} else {
				jump := (cycleCount - i) / cycleOffset
				i += jump * cycleOffset
			}
		}
	}

	load := 0
	for r, row := range grid {
		for _, cell := range row {
			if cell == "O" {
				load += maxRow - r
			}
		}
	}
	fmt.Println(load)
}

type CacheData struct {
	Grid    [][]string
	GridKey string
}

func copy(source [][]string, target [][]string, maxRow, maxCol int) {
	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			target[r][c] = source[r][c]
		}
	}
}

func genKey(grid [][]string, maxRow, maxCol int) string {
	key := ""

	for r := 0; r < maxRow; r++ {
		for c := 0; c < maxCol; c++ {
			if grid[r][c] == "O" {
				key += fmt.Sprintf("%d_%d/", r, c)
			}
		}
	}
	return key
}

func tilt(generator Generator, grid [][]string, maxRow, maxCol int) {
	var start [2]int = [2]int{-1, -1}
	var stoneCount int

	for {
		index, dir, laneDone, gridDone := generator()
		r, c := index[0], index[1]
		cell := grid[r][c]

		if start[0] == -1 && start[1] == -1 && (cell == "O" || cell == ".") {
			start = index
		}

		if cell == "O" {
			stoneCount += 1
			grid[r][c] = "."
		}

		if cell == "#" || laneDone {
			for j := 0; j < stoneCount; j++ {
				grid[start[0]+j*dir[0]][start[1]+j*dir[1]] = "O"
			}
			start = [2]int{-1, -1}
			stoneCount = 0
		}

		if gridDone {
			break
		}
	}

}

type GeneratorGenerator func(maxRow, maxCol int) func() ([2]int, [2]int, bool, bool)
type Generator func() ([2]int, [2]int, bool, bool)

func north(maxRow, maxCol int) func() ([2]int, [2]int, bool, bool) {
	r := 0
	c := 0

	return func() ([2]int, [2]int, bool, bool) {
		index := [2]int{r, c}
		laneDone := r == maxRow-1
		gridDone := laneDone && c == maxCol-1

		r += 1
		if laneDone {
			r = 0
			c += 1
		}

		return index, [2]int{1, 0}, laneDone, gridDone
	}
}

func east(maxRow, maxCol int) func() ([2]int, [2]int, bool, bool) {
	r := 0
	c := maxCol - 1

	return func() ([2]int, [2]int, bool, bool) {
		index := [2]int{r, c}
		laneDone := c == 0
		gridDone := laneDone && r == maxRow-1

		c -= 1
		if laneDone {
			c = maxCol - 1
			r += 1
		}

		return index, [2]int{0, -1}, laneDone, gridDone
	}
}

func west(maxRow, maxCol int) func() ([2]int, [2]int, bool, bool) {
	r := 0
	c := 0

	return func() ([2]int, [2]int, bool, bool) {
		index := [2]int{r, c}
		laneDone := c == maxCol-1
		gridDone := laneDone && r == maxRow-1

		c += 1
		if laneDone {
			c = 0
			r += 1
		}

		return index, [2]int{0, 1}, laneDone, gridDone
	}
}

func south(maxRow, maxCol int) func() ([2]int, [2]int, bool, bool) {
	r := maxRow - 1
	c := 0

	return func() ([2]int, [2]int, bool, bool) {
		index := [2]int{r, c}
		laneDone := r == 0
		gridDone := laneDone && c == maxCol-1

		r -= 1
		if laneDone {
			r = maxRow - 1
			c += 1
		}

		return index, [2]int{-1, 0}, laneDone, gridDone
	}
}
