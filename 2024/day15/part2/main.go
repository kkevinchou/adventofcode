package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

func main() {
	smallGrid, instructions, smallRCount, smallCCount := ParseGrid(file)

	grid := make([][]string, len(smallGrid))
	for r := range smallRCount {
		grid[r] = make([]string, 2*len(smallGrid[0]))
		for c := range smallCCount {
			grid[r][c*2] = smallGrid[r][c]
			grid[r][c*2+1] = smallGrid[r][c]
			if smallGrid[r][c] == "@" {
				grid[r][c*2+1] = "."
			} else if smallGrid[r][c] == "O" {
				grid[r][c*2] = "["
				grid[r][c*2+1] = "]"
			}
		}
	}
	rCount := smallRCount
	cCount := smallCCount * 2

	var done bool
	var r int
	var c int
	for r = range rCount {
		for c = range cCount {
			if grid[r][c] == "@" {
				done = true
				break
			}
		}
		if done {
			break
		}
	}

	for _, instruction := range instructions {
		if instruction == "^" {
			r, c = MoveUp(grid, rCount, cCount, r, c)
		} else if instruction == "v" {
			r, c = MoveDown(grid, rCount, cCount, r, c)
		} else if instruction == "<" {
			r, c = MoveLeft(grid, rCount, cCount, r, c)
		} else if instruction == ">" {
			r, c = MoveRight(grid, rCount, cCount, r, c)
		}
	}

	var result int
	for r := range rCount {
		for c := range cCount {
			if grid[r][c] == "[" {
				result += r*100 + c
			}
		}
	}
	fmt.Println(result)
}

func MoveDown(grid [][]string, rCount, cCount int, r, c int) (int, int) {
	if CanMoveDown(grid, rCount, cCount, r, c, false) {
		ForceMoveDown(grid, rCount, cCount, r, c, false)
		r += 1
	}
	return r, c
}

func MoveUp(grid [][]string, rCount, cCount int, r, c int) (int, int) {
	if CanMoveUp(grid, rCount, cCount, r, c, false) {
		ForceMoveUp(grid, rCount, cCount, r, c, false)
		r -= 1
	}
	return r, c
}
func ForceMoveUp(grid [][]string, rCount, cCount int, r, c int, siblingCall bool) {
	if grid[r][c] == "." {
		return
	}
	if !siblingCall {
		if grid[r][c] == "[" {
			ForceMoveUp(grid, rCount, cCount, r, c+1, true)
		} else if grid[r][c] == "]" {
			ForceMoveUp(grid, rCount, cCount, r, c-1, true)
		}
	}

	ForceMoveUp(grid, rCount, cCount, r-1, c, false)
	grid[r-1][c] = grid[r][c]
	grid[r][c] = "."
}

func ForceMoveDown(grid [][]string, rCount, cCount int, r, c int, siblingCall bool) {
	if grid[r][c] == "." {
		return
	}

	if !siblingCall {
		if grid[r][c] == "[" {
			ForceMoveDown(grid, rCount, cCount, r, c+1, true)
		} else if grid[r][c] == "]" {
			ForceMoveDown(grid, rCount, cCount, r, c-1, true)
		}
	}

	ForceMoveDown(grid, rCount, cCount, r+1, c, false)
	grid[r+1][c] = grid[r][c]
	grid[r][c] = "."
}

func CanMoveUp(grid [][]string, rCount, cCount int, r, c int, siblingCall bool) bool {
	if grid[r][c] == "#" {
		return false
	}

	if grid[r][c] == "." {
		return true
	}

	var siblingSuccess = true
	if !siblingCall {
		if grid[r][c] == "[" {
			siblingSuccess = CanMoveUp(grid, rCount, cCount, r, c+1, true)
		} else if grid[r][c] == "]" {
			siblingSuccess = CanMoveUp(grid, rCount, cCount, r, c-1, true)
		}
	}

	return siblingSuccess && CanMoveUp(grid, rCount, cCount, r-1, c, false)
}

func CanMoveDown(grid [][]string, rCount, cCount int, r, c int, siblingCall bool) bool {
	if grid[r][c] == "#" {
		return false
	}

	if grid[r][c] == "." {
		return true
	}

	var siblingSuccess = true
	if !siblingCall {
		if grid[r][c] == "[" {
			siblingSuccess = CanMoveDown(grid, rCount, cCount, r, c+1, true)
		} else if grid[r][c] == "]" {
			siblingSuccess = CanMoveDown(grid, rCount, cCount, r, c-1, true)
		}
	}

	return siblingSuccess && CanMoveDown(grid, rCount, cCount, r+1, c, false)
}

func MoveRight(grid [][]string, rCount, cCount int, r, c int) (int, int) {
	startC := c

	for col := c + 1; col < cCount; col++ {
		if grid[r][col] == "." {
			for k := col; k > startC; k-- {
				grid[r][k] = grid[r][k-1]
			}
			grid[r][startC] = "."
			return r, c + 1
		} else if grid[r][col] == "#" {
			return r, c
		}
	}
	return r, c
}
func MoveLeft(grid [][]string, rCount, cCount int, r, c int) (int, int) {
	startC := c

	for col := c - 1; col >= 0; col-- {
		if grid[r][col] == "." {
			for k := col; k < startC; k++ {
				grid[r][k] = grid[r][k+1]
			}
			grid[r][startC] = "."
			return r, c - 1
		} else if grid[r][col] == "#" {
			return r, c
		}
	}
	return r, c
}

func ParseGrid(inputFile string) ([][]string, []string, int, int) {
	var grid [][]string
	var instructions []string

	step := 0
	for record := range utils.Records(inputFile) {
		if record.Line == "" {
			step++
		}
		if step == 0 {
			grid = append(grid, make([]string, len(record.Line)))
			for c, char := range record.Line {
				r := record.LineNumber
				grid[r][c] = string(char)
			}
		} else if step == 1 {
			for _, char := range record.Line {
				instructions = append(instructions, string(char))
			}
		}
	}
	rCount := len(grid)
	cCount := len(grid[0])

	return grid, instructions, rCount, cCount
}
