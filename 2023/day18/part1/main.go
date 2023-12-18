package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var instructions []Instruction
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine

		lineSplit := strings.Split(line, " ")

		var dir Direction
		if lineSplit[0] == "R" {
			dir = RIGHT
		} else if lineSplit[0] == "D" {
			dir = DOWN
		} else if lineSplit[0] == "L" {
			dir = LEFT
		} else if lineSplit[0] == "U" {
			dir = UP
		}

		amount := utils.MustParseNum(lineSplit[1])
		color := lineSplit[2]

		var movement [2]int
		if dir == UP {
			movement = [2]int{-amount, 0}
		} else if dir == DOWN {
			movement = [2]int{amount, 0}
		} else if dir == LEFT {
			movement = [2]int{0, -amount}
		} else if dir == RIGHT {
			movement = [2]int{0, amount}
		}

		instruction := Instruction{Direction: dir, Movement: movement, Color: color, Amount: amount}
		instructions = append(instructions, instruction)
	}

	maxRow := -10000
	maxCol := -10000

	minRow := 10000
	minCol := 10000

	currentRow := 0
	currentCol := 0
	for _, instruction := range instructions {
		if instruction.Direction == LEFT || instruction.Direction == RIGHT {
			currentCol += instruction.Movement[1]
			if currentCol > maxCol {
				maxCol = currentCol
			} else if currentCol < minCol {
				minCol = currentCol
			}
		} else if instruction.Direction == UP || instruction.Direction == DOWN {
			currentRow += instruction.Movement[0]
			if currentRow > maxRow {
				maxRow = currentRow
			} else if currentRow < minRow {
				minRow = currentRow
			}
		}
	}

	startRow := minRow
	if startRow < 0 {
		startRow = -startRow
	}
	startCol := minCol
	if startCol < 0 {
		startCol = -startCol
	}

	maxRow = maxRow - minRow + 1
	maxCol = maxCol - minCol + 1

	var grid [][]string
	for r := 0; r < maxRow; r++ {
		grid = append(grid, make([]string, maxCol))
		for c := 0; c < maxCol; c++ {
			grid[r][c] = "."
		}
	}

	fmt.Println(maxRow, maxCol)
	current := [2]int{startRow, startCol}
	fmt.Println("START", current)

	for _, instruction := range instructions {
		for i := 0; i < instruction.Amount; i++ {
			current[0] += instruction.Direction[0]
			current[1] += instruction.Direction[1]
			grid[current[0]][current[1]] = "#"
			// fmt.Println(current)
		}
	}

	start := [2]int{4, 155}
	// start := [2]int{1, 1}

	fill(grid, maxRow, maxCol, start)

	// for _, row := range grid {
	// 	// fmt.Println(len(row))
	// 	for _, s := range row {
	// 		fmt.Printf("%s", s)
	// 	}
	// 	fmt.Printf("\n")
	// }

	total := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == "#" {
				total += 1
			}
		}
	}
	fmt.Println(total)
}

func fill(grid [][]string, maxRow, maxCol int, start [2]int) {
	nodes := [][2]int{start}

	seen := map[[2]int]bool{}

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]

		if _, ok := seen[node]; ok {
			continue
		}

		r, c := node[0], node[1]
		seen[node] = true

		grid[r][c] = "#"
		neighbors := getNeighbors(r, c, maxRow, maxCol)
		for _, neighbor := range neighbors {
			if grid[neighbor[0]][neighbor[1]] == "." {
				nodes = append(nodes, neighbor)
			}
		}
	}
}

type Instruction struct {
	Direction [2]int
	Movement  [2]int
	Amount    int
	Color     string
}

func getNeighbors(r, c, maxRow, maxCol int) [][2]int {
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

var dirToBack map[[2]int][2]int = map[[2]int][2]int{
	[2]int{1, 0}:  [2]int{-1, 0},
	[2]int{0, -1}: [2]int{0, 1},
	[2]int{-1, 0}: [2]int{1, 0},
	[2]int{0, 1}:  [2]int{0, -1},
}

type Direction [2]int

var LEFT Direction = Direction{0, -1}
var RIGHT Direction = Direction{0, 1}
var UP Direction = Direction{-1, 0}
var DOWN Direction = Direction{1, 0}
