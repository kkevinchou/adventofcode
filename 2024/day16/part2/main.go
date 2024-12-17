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
	splits := map[string]int{}
	result, _, nodes, success := solve(grid, startR, startC, rCount, cCount, 0, 0, seen, splits)

	fmt.Println(len(nodes))

	// for k, _ := range nodes {
	// 	grid[k[0]][k[1]] = "O"
	// }

	// for k, v := range splits {
	// 	fmt.Println(k, v)
	// 	// s := strings.Split(k, "_")
	// 	// grid[utils.MustParseNum(s[0])][utils.MustParseNum(s[1])] = "K"
	// }

	// utils.PrintGrid(grid)
	if !success {
		panic("wat")
	}
	fmt.Println(result)
}

func solve(grid [][]string, r, c, rCount, cCount int, incomingDir int, costSoFar int, seen map[string]int, splits map[string]int) (int, int, map[[2]int]bool, bool) {
	if grid[r][c] == "E" {
		return costSoFar, 0, map[[2]int]bool{{r, c}: true}, true
	}

	if previousCost, ok := seen[key(r, c, incomingDir)]; ok {
		if costSoFar > previousCost {
			return -1, -1, nil, false
		}
	}

	seen[key(r, c, incomingDir)] = costSoFar

	forward := incomingDir
	left := (incomingDir + 3) % 4
	right := (incomingDir + 1) % 4

	costs := []int{1, 1001, 1001}
	nextDir := []int{forward, left, right}

	bestCost := 9999999999
	var anySuccess bool
	var bestStepsToGoal int
	var allCosts []int
	var allNeighborNodes []map[[2]int]bool
	for i, dir := range nextDir {
		offset := dirs[dir]
		newR, newC := r+offset[0], c+offset[1]

		if newR < 0 || newR >= rCount || newC < 0 || newC >= cCount || grid[newR][newC] == "#" {
			continue
		}

		neighborCost, stepsToGoal, neighborNodes, success := solve(grid, newR, newC, rCount, cCount, dir, costSoFar+costs[i], seen, splits)
		if success {
			allCosts = append(allCosts, neighborCost)
			allNeighborNodes = append(allNeighborNodes, neighborNodes)
			anySuccess = true
			if neighborCost < bestCost {
				bestCost = neighborCost
				bestStepsToGoal = stepsToGoal
			}
		}
	}

	if anySuccess {
		if r == 7 && c == 3 && incomingDir == 3 {
			a := 1
			_ = a
		}
		nodes := map[[2]int]bool{
			{r, c}: true,
		}

		var numBestPaths int
		for i, cost := range allCosts {
			if cost == bestCost {
				numBestPaths++
				for k, v := range allNeighborNodes[i] {
					nodes[k] = v
				}
			}
		}

		if numBestPaths > 1 {
			splits[key(r, c, incomingDir)] = numBestPaths
		}
		return bestCost, bestStepsToGoal + 1, nodes, anySuccess
	}

	return -1, -1, nil, false
}

func key(r, c, dir int) string {
	// return fmt.Sprintf("%d_%d", r, c)
	return fmt.Sprintf("%d_%d_%d", r, c, dir)
}
