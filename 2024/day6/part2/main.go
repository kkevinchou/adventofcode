package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

var dirs [][2]int = [][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func main() {
	start := time.Now()
	blockers := firstPass()

	var startR, startC int
	var grid [][]string

	for record := range utils.Records(file, "\n") {
		grid = append(grid, make([]string, len(record.SingleLine)))
		for c, char := range record.SingleLine {
			grid[record.ID][c] = string(char)

			if string(char) == "^" {
				startR = record.ID
				startC = c
			}
		}
	}
	rCount := len(grid)
	cCount := len(grid[0])

	var result int
	var lock sync.Mutex

	var wg sync.WaitGroup

	wg.Add(len(blockers))

	for i := range blockers {
		go func() {
			if try(grid, startR, startC, rCount, cCount, blockers[i]) {
				lock.Lock()
				result++
				lock.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(result)
	fmt.Println(time.Since(start))
}

func try(grid [][]string, startR, startC, rCount, cCount int, block [2]int) bool {
	r, c := startR, startC
	dir := 0
	lookup := map[string]bool{}

	// place starting position in lookup
	key := genKey(r, c, dir)
	if lookup[key] {
		return true
	}

	for {
		nextR, nextC := r+dirs[dir][0], c+dirs[dir][1]
		if nextR < 0 || nextR >= rCount || nextC < 0 || nextC >= cCount {
			break
		}

		if grid[nextR][nextC] == "#" || (nextR == block[0] && nextC == block[1]) {
			key := genKey(r, c, dir)
			if lookup[key] {
				return true
			}
			lookup[key] = true
			dir = (dir + 1) % len(dirs)
			continue
		}

		r, c = nextR, nextC
	}

	return false
}

func firstPass() [][2]int {
	var grid [][]string

	var startR, startC int
	var dir = 0

	for record := range utils.Records(file, "\n") {
		grid = append(grid, make([]string, len(record.SingleLine)))
		for c, char := range record.SingleLine {
			grid[record.ID][c] = string(char)

			if string(char) == "^" {
				startR = record.ID
				startC = c
			}
		}
	}
	rCount := len(grid)
	cCount := len(grid[0])

	grid[startR][startC] = "X"

	r, c := startR, startC

	var blockers [][2]int
	for {
		nextR, nextC := r+dirs[dir][0], c+dirs[dir][1]
		if nextR < 0 || nextR >= rCount || nextC < 0 || nextC >= cCount {
			break
		}

		if grid[nextR][nextC] == "#" {
			dir = (dir + 1) % len(dirs)
			continue
		}

		r, c = nextR, nextC

		if grid[r][c] != "X" {
			grid[r][c] = "X"
			blockers = append(blockers, [2]int{r, c})
		}
	}

	return blockers
}

func copyGrid(grid [][]string) [][]string {
	copy := make([][]string, len(grid))
	for i := range len(grid) {
		copy[i] = make([]string, len(grid[0]))
		for j := range len(grid[0]) {
			copy[i][j] = grid[i][j]
		}
	}
	return copy
}

func printGrid(grid [][]string) {
	for r := range len(grid) {
		for c := range len(grid[0]) {
			fmt.Printf(grid[r][c] + " ")
		}
		fmt.Printf("\n")
	}
}

func genKey(r, c, dir int) string {
	return fmt.Sprintf("%d_%d_%d", r, c, dir)
}
