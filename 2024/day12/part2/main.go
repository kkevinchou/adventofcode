package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "custom4"

var regionIDGen int

type Data struct {
	RegionID int
}

func key(r, c int) string {
	return fmt.Sprintf("%d_%d", r, c)
}

func main() {
	solve()

	// file = "sample"
	// solve()
	// file = "sample4"
	// solve()
	// file = "sample2"
	// solve()
	// file = "sample3"
	// solve()
	// file = "sample5"
	// solve()
	// file = "custom"
	// solve()
	// file = "custom2"
	// solve()

	/*
		18 * (4 + 4)
		5 * (8)
		5 * (8)
		2 * 4
	*/

	// // 80, 436, 236, 368, 1206, 276, 232
}

func solve() {
	grid, rCount, cCount := utils.ParseGrid(file)

	seen := map[[2]int]bool{}
	var regionIDs [][]int
	areas := map[int]int{}

	for i := range rCount {
		regionIDs = append(regionIDs, make([]int, cCount))
		for j := range regionIDs[i] {
			regionIDs[i][j] = -1
		}
	}

	var result int
	for r := range rCount {
		for c := range cCount {
			if _, ok := seen[[2]int{r, c}]; ok {
				continue
			}

			area := floodFillArea(grid, r, c, rCount, cCount, grid[r][c], seen, regionIDs, regionIDGen)
			areas[regionIDGen] = area

			sides, _, _ := traversePerimeter(grid, r, c, rCount, cCount, [2]int{r, c}, isPlotCell, regionIDs, -999, nil)
			// fmt.Println("current region", regionIDGen, "area", areas[regionIDGen], "sides", sides)
			result += area * sides
			regionIDGen++
		}
	}

	seenPerimeter := map[string]bool{}
	for r := range rCount {
		for c := range cCount {
			leftR := r - 1
			if leftR < 0 {
				continue
			}

			potentialContainerRegionID := regionIDs[leftR][c]
			if potentialContainerRegionID < 0 {
				panic("WAT")
			}

			if potentialContainerRegionID == regionIDs[r][c] {
				continue
			}

			if _, ok := seenPerimeter[key(r, c)]; ok {
				continue
			}

			sides, neighborRegionIDs, clockwise := traversePerimeter(grid, r, c, rCount, cCount, [2]int{r, c}, isNotRegion, regionIDs, potentialContainerRegionID, seenPerimeter)
			if len(neighborRegionIDs) == 1 && clockwise {
				for k, _ := range neighborRegionIDs {
					if k != potentialContainerRegionID {
						panic("WAT")
					}
					break
				}
				if potentialContainerRegionID == -1 {
					panic("WAT")
				}

				// fmt.Println("parent region", potentialContainerRegionID, "area", areas[potentialContainerRegionID], "sides", sides)
				result += areas[potentialContainerRegionID] * sides
			}
		}
	}
	fmt.Println(result)
}

func traversePerimeter(grid [][]string, r, c, rCount, cCount int, start [2]int, plotFn PlotFn, regionIDs [][]int, notRegionID int, seenPerimeter map[string]bool) (int, map[int]bool, bool) {
	// var perimeter int
	var dir int

	plot := grid[r][c]
	var sides int
	hitRegionIDs := map[int]bool{}

	leftCount := 0
	rightCount := 0

	for {
		// fmt.Printf("[%d,%d] - DIR: [%d,%d]\n", r, c, dirs[dir][0], dirs[dir][1])
		left := (dir + 3) % 4
		lr, lc := r+dirs[left][0], c+dirs[left][1]

		if seenPerimeter != nil {
			seenPerimeter[key(r, c)] = true
		}

		if plotFn(grid, regionIDs, lr, lc, rCount, cCount, plot, notRegionID) {
			dir = left
			r = lr
			c = lc
			sides++
			leftCount++
		} else {
			fr := r + dirs[dir][0]
			fc := c + dirs[dir][1]

			if lr >= 0 && lr < rCount && lc >= 0 && lc < cCount {
				hitRegionIDs[regionIDs[lr][lc]] = true
			} else {
				hitRegionIDs[-1] = true
			}

			// can't move forward
			if !plotFn(grid, regionIDs, fr, fc, rCount, cCount, plot, notRegionID) {
				dir = (dir + 1) % 4
				sides++
				rightCount++
			} else {
				r = fr
				c = fc
			}
		}
		if dir == 0 && start[0] == r && start[1] == c {
			break
		}
	}

	return sides, hitRegionIDs, rightCount > leftCount
}

func floodFillSeenForPerimeter(grid [][]string, r, c, rCount, cCount int, potentialRegionID int, regionIDs [][]int, seenPerimeter map[string]bool) {
	if _, ok := seenPerimeter[key(r, c)]; ok {
		return
	}

	if regionIDs[r][c] == potentialRegionID {
		return
	}

	seenPerimeter[key(r, c)] = true

	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]
		if nr < 0 || nr >= rCount || nc < 0 || nc >= cCount {
			continue
		}

		floodFillSeenForPerimeter(grid, nr, nc, rCount, cCount, potentialRegionID, regionIDs, seenPerimeter)
	}
}

func floodFillArea(grid [][]string, r, c, rCount, cCount int, plot string, seen map[[2]int]bool, regionIDs [][]int, regionID int) int {
	if _, ok := seen[[2]int{r, c}]; ok {
		return 0
	}

	if grid[r][c] != plot {
		return 0
	}

	seen[[2]int{r, c}] = true
	regionIDs[r][c] = regionID

	result := 1

	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]
		if nr < 0 || nr >= rCount || nc < 0 || nc >= cCount {
			continue
		}

		result += floodFillArea(grid, nr, nc, rCount, cCount, plot, seen, regionIDs, regionID)
	}
	return result
}

type PlotFn func(grid [][]string, regionIDs [][]int, r, c, rCount, cCount int, plot string, notRegionID int) bool

func isPlotCell(grid [][]string, regionIDs [][]int, r, c, rCount, cCount int, plot string, regionID int) bool {
	if r < 0 || r >= rCount || c < 0 || c >= cCount {
		return false
	}
	return grid[r][c] == plot
}

func isNotRegion(grid [][]string, regionIDs [][]int, r, c, rCount, cCount int, plot string, regionID int) bool {
	if r < 0 || r >= rCount || c < 0 || c >= cCount {
		return false
	}
	return regionIDs[r][c] != regionID
}

var dirs [][2]int = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// 826904  - too low
// 849724  - wrong
// 852824  - wrong
// 860708  - wrong
// 867836  - broken samples - this double counts should it should theoretically be the max
// 2709122 - too high
// 3272136 - too high
// 1421940 - wrong
