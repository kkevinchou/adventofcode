package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

type Position [2]int

func main() {
	grid, instructions, rCount, cCount := ParseGrid(file)

	var position Position
	var done bool
	for r := range rCount {
		for c := range cCount {
			if grid[r][c] == "@" {
				position[0] = r
				position[1] = c
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
			for r := position[0] - 1; r >= 0; r-- {
				if grid[r][position[1]] == "." {
					grid[r][position[1]] = "O"

					grid[position[0]][position[1]] = "."
					position[0] = position[0] - 1
					grid[position[0]][position[1]] = "@"
					break
				} else if grid[r][position[1]] == "#" {
					break
				}
			}
		} else if instruction == "v" {
			for r := position[0] + 1; r < rCount; r++ {
				if grid[r][position[1]] == "." {
					grid[r][position[1]] = "O"
					grid[position[0]][position[1]] = "."
					position[0] = position[0] + 1
					grid[position[0]][position[1]] = "@"
					break
				} else if grid[r][position[1]] == "#" {
					break
				}
			}
		} else if instruction == "<" {
			for c := position[1] - 1; c >= 0; c-- {
				if grid[position[0]][c] == "." {
					grid[position[0]][c] = "O"
					grid[position[0]][position[1]] = "."
					position[1] = position[1] - 1
					grid[position[0]][position[1]] = "@"
					break
				} else if grid[position[0]][c] == "#" {
					break
				}
			}
		} else if instruction == ">" {
			for c := position[1] + 1; c < cCount; c++ {
				if grid[position[0]][c] == "." {
					grid[position[0]][c] = "O"
					grid[position[0]][position[1]] = "."
					position[1] = position[1] + 1
					grid[position[0]][position[1]] = "@"
					break
				} else if grid[position[0]][c] == "#" {
					break
				}
			}
		}
	}

	var result int
	for r := range rCount {
		for c := range cCount {
			if grid[r][c] == "O" {
				result += r*100 + c
			}
		}
	}

	fmt.Println(result)
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
