package main

import "fmt"

func main() {
	grid := [][]string{}
	recordGenerator := constructRecordGenerator("input", "\n")
	for {
		record, done := recordGenerator()
		if done {
			break
		}
		line := record.Lines[0]
		row := []string{}
		for _, cell := range line {
			row = append(row, string(cell))
		}
		grid = append(grid, row)
	}

	// fmt.Println(grid)
	for {
		nextGrid := Simulate(grid)
		if equal, occupied := isEqual(grid, nextGrid); equal {
			fmt.Println(occupied)
			break
		}
		grid = nextGrid
	}
}

func isEqual(grid1 [][]string, grid2 [][]string) (bool, int) {
	occupied := 0
	for x := range grid1 {
		for y := range grid1[x] {
			if grid1[x][y] != grid2[x][y] {
				return false, -1
			}

			if grid1[x][y] == "#" {
				occupied++
			}
		}
	}
	return true, occupied
}

func Simulate(grid [][]string) [][]string {
	nextGrid := [][]string{}

	// initialize next grid
	for _, v := range grid {
		nextGrid = append(nextGrid, make([]string, len(v)))
	}

	for x := range grid {
		for y := range grid[x] {
			neighbors := getNeighbors(grid, x, y)
			occupiedCount := getOccupiedCount(grid, neighbors)
			if grid[x][y] == "L" && occupiedCount == 0 {
				nextGrid[x][y] = "#"
			} else if grid[x][y] == "#" && occupiedCount >= 5 {
				nextGrid[x][y] = "L"
			} else {
				nextGrid[x][y] = grid[x][y]
			}
		}
	}

	return nextGrid
}

func getNeighbors(grid [][]string, x, y int) [][]int {
	deltas := []int{-1, 0, 1}

	result := [][]int{}
	for _, delta1 := range deltas {
		for _, delta2 := range deltas {
			if delta1 == 0 && delta2 == 0 {
				continue
			}
			neighborX := x + delta1
			neighborY := y + delta2

			for {
				// fmt.Println("getNeighbors", neighborX, neighborY)

				if neighborX >= 0 && neighborX < len(grid) && neighborY >= 0 && neighborY < len(grid[0]) {
					if grid[neighborX][neighborY] == "." {
						neighborX += delta1
						neighborY += delta2
						continue
					} else {
						result = append(result, []int{neighborX, neighborY})
						break
					}
				} else {
					break
				}
			}
		}
	}

	return result
}

func getOccupiedCount(grid [][]string, cells [][]int) int {
	total := 0
	for _, pair := range cells {
		if grid[pair[0]][pair[1]] == "#" {
			total++
		}
	}
	return total
}
