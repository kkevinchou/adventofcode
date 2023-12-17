package main

import (
	"fmt"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

// Options
// 1 = recursively call
// 2 - something with a for loop

func main() {
	startTime := time.Now()
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var grid [][]string
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine

		row := utils.StringToStringSlice(line)
		grid = append(grid, row)
	}

	maxRow, maxCol := len(grid), len(grid[0])

	starts := [][2]int{}

	for i := 0; i < maxCol; i++ {
		starts = append(starts, [2]int{0, i})
		starts = append(starts, [2]int{maxRow - 1, i})
	}

	for i := 0; i < maxRow; i++ {
		starts = append(starts, [2]int{i, 0})
		starts = append(starts, [2]int{i, maxCol - 1})
	}

	maxTiles := 0
	for _, start := range starts {
		var dirs [][2]int

		if start[0] == 0 {
			dirs = append(dirs, DOWN)
		} else if start[0] == maxRow-1 {
			dirs = append(dirs, UP)
		}

		if start[1] == 0 {
			dirs = append(dirs, RIGHT)
		} else if start[1] == maxCol-1 {
			dirs = append(dirs, LEFT)
		}

		for _, dir := range dirs {
			memo := map[string]bool{}
			visited := map[string]bool{}
			solve(grid, maxRow, maxCol, start, dir, memo, visited)
			if len(visited) > maxTiles {
				maxTiles = len(visited)
			}
		}
	}

	fmt.Println(maxTiles)
	fmt.Println(time.Since(startTime))
}

func solve(grid [][]string, maxRow, maxCol int, position [2]int, dir [2]int, memo map[string]bool, visited map[string]bool) {
	key := genKey(position, dir)
	if _, ok := memo[key]; ok {
		return
	}
	memo[key] = true

	if grid[position[0]][position[1]] == "." {
		grid[position[0]][position[1]] = "#"
	}

	visitedKey := genVisitedKey(position)
	visited[visitedKey] = true

	cell := grid[position[0]][position[1]]
	nextDirections := travelCell(dir, cell)

	for _, nextDir := range nextDirections {
		nextPosition := [2]int{position[0] + nextDir[0], position[1] + nextDir[1]}

		if nextPosition[0] < 0 || nextPosition[0] >= maxRow {
			continue
		}

		if nextPosition[1] < 0 || nextPosition[1] >= maxCol {
			continue
		}

		solve(grid, maxRow, maxCol, nextPosition, nextDir, memo, visited)
	}
}

var dirs [][2]int = [][2]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

func genVisitedKey(position [2]int) string {
	return fmt.Sprintf("%d_%d", position[0], position[1])
}

func genKey(position [2]int, dir [2]int) string {
	return fmt.Sprintf("%d_%d_%d_%d", position[0], position[1], dir[0], dir[1])
}

func travelCell(dir [2]int, cell string) [][2]int {
	dirs := [][2]int{}

	if cell == "/" {
		if dir == RIGHT {
			dirs = append(dirs, UP)
		} else if dir == LEFT {
			dirs = append(dirs, DOWN)
		} else if dir == UP {
			dirs = append(dirs, RIGHT)
		} else if dir == DOWN {
			dirs = append(dirs, LEFT)
		} else {
			panic("WAT")
		}
	} else if cell == "\\" {
		if dir == RIGHT {
			dirs = append(dirs, DOWN)
		} else if dir == LEFT {
			dirs = append(dirs, UP)
		} else if dir == UP {
			dirs = append(dirs, LEFT)
		} else if dir == DOWN {
			dirs = append(dirs, RIGHT)
		} else {
			panic("WAT")
		}
	} else if cell == "-" {
		if dir == RIGHT {
			dirs = append(dirs, RIGHT)
		} else if dir == LEFT {
			dirs = append(dirs, LEFT)
		} else if dir == UP || dir == DOWN {
			dirs = append(dirs, LEFT)
			dirs = append(dirs, RIGHT)
		}
	} else if cell == "|" {
		if dir == RIGHT || dir == LEFT {
			dirs = append(dirs, UP)
			dirs = append(dirs, DOWN)
		} else if dir == DOWN {
			dirs = append(dirs, DOWN)
		} else if dir == UP {
			dirs = append(dirs, UP)
		}
	} else if cell == "." || cell == "#" {
		dirs = append(dirs, dir)
	} else {
		panic("WAT")
	}

	return dirs
}

var LEFT [2]int = [2]int{0, -1}
var RIGHT [2]int = [2]int{0, 1}
var UP [2]int = [2]int{-1, 0}
var DOWN [2]int = [2]int{1, 0}
