package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

type Position [2]int

var dirs [][2]int = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

var dirToString map[int]string = map[int]string{
	0: "east",
	1: "south",
	2: "west",
	3: "north",
}

func main() {
	grid, rCount, cCount := utils.ParseGrid(file)

	var startR, startC int
	for r := range rCount {
		for c := range cCount {
			if grid[r][c] == "S" {
				startR = r
				startC = c
			}
		}
	}

	seen := map[string]int{}
	result, success := solve(grid, startR, startC, rCount, cCount, 0, 0, seen)
	if !success {
		panic("wat")
	}
	fmt.Println(result)
}

func solve(grid [][]string, r, c, rCount, cCount int, incomingDir int, costSoFar int, seen map[string]int) (int, bool) {
	if grid[r][c] == "E" {
		return costSoFar, true
	}

	if previousCost, ok := seen[key(r, c)]; ok {
		if costSoFar > previousCost {
			return 0, false
		}
	}

	seen[key(r, c)] = costSoFar

	forward := incomingDir
	left := (incomingDir + 3) % 4
	right := (incomingDir + 1) % 4

	costs := []int{1, 1001, 1001}
	nextDir := []int{forward, left, right}

	bestCost := 9999999999
	var anySuccess bool
	for i, dir := range nextDir {
		offset := dirs[dir]
		newR, newC := r+offset[0], c+offset[1]

		if r == 7 && c == 3 && dir == 3 {
			a := 1
			_ = a
		}

		if newR < 0 || newR >= rCount || newC < 0 || newC >= cCount || grid[newR][newC] == "#" {
			continue
		}

		neighborCost, success := solve(grid, newR, newC, rCount, cCount, dir, costSoFar+costs[i], seen)
		if success {
			anySuccess = true
			if neighborCost < bestCost {
				bestCost = neighborCost
			}
		}
	}

	if anySuccess {
		return bestCost, anySuccess
	}

	return -1, false
}

func key(r, c int) string {
	return fmt.Sprintf("%d_%d", r, c)
	// return fmt.Sprintf("%d_%d_%d", r, c, dir)
}
