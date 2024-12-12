package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "custom3"

var regionIDGen int

type Data struct {
	RegionID int
}

func key(r, c, potentialRegionID int) string {
	return fmt.Sprintf("%d_%d_%d", r, c, potentialRegionID)
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
	seenPerimeter := map[string]bool{}

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

			sides, _, _ := traversePerimeter(grid, r, c, rCount, cCount, [2]int{r, c}, isPlotCell, regionIDs, -999)
			fmt.Println("current region", regionIDGen, "area", areas[regionIDGen], "sides", sides)

			leftR := r - 1
			if leftR >= 0 {
				potentialContainerRegionID := regionIDs[leftR][c]
				if _, ok := seenPerimeter[key(r, c, potentialContainerRegionID)]; !ok {
					internalSides, neighborRegionIDs, clockwise := traversePerimeter(grid, r, c, rCount, cCount, [2]int{r, c}, isNotRegion, regionIDs, potentialContainerRegionID)
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
						fmt.Println("parent region", potentialContainerRegionID, "area", areas[potentialContainerRegionID], "sides", internalSides)
						result += areas[potentialContainerRegionID] * internalSides
					}
					floodFillSeenForPerimeter(grid, r, c, rCount, cCount, potentialContainerRegionID, regionIDs, seenPerimeter)
				}
			}

			result += area * sides
			regionIDGen++
		}
	}
	fmt.Println(result)
}

type PlotFn func(grid [][]string, dataGrid [][]int, r, c, rCount, cCount int, plot string, notRegionID int) bool

func traversePerimeter(grid [][]string, r, c, rCount, cCount int, start [2]int, plotFn PlotFn, regionIDs [][]int, notRegionID int) (int, map[int]bool, bool) {
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

		if plotFn(grid, regionIDs, lr, lc, rCount, cCount, plot, notRegionID) {
			dir = left
			r = lr
			c = lc
			sides++
			leftCount++
		} else {
			if lr >= 0 && lr < rCount && lc >= 0 && lc < cCount {
				hitRegionIDs[regionIDs[lr][lc]] = true
			} else {
				hitRegionIDs[-1] = true
			}
			fr := r + dirs[dir][0]
			fc := c + dirs[dir][1]
			// perimeter++

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

func isPlotCell(grid [][]string, dataGrid [][]int, r, c, rCount, cCount int, plot string, regionID int) bool {
	if r < 0 || r >= rCount || c < 0 || c >= cCount {
		return false
	}
	return grid[r][c] == plot
}

func isNotRegion(grid [][]string, dataGrid [][]int, r, c, rCount, cCount int, plot string, regionID int) bool {
	if r < 0 || r >= rCount || c < 0 || c >= cCount {
		return false
	}
	return dataGrid[r][c] != regionID
}

func floodFillSeenForPerimeter(grid [][]string, r, c, rCount, cCount int, potentialRegionID int, regionIDs [][]int, seenPerimeter map[string]bool) {
	if _, ok := seenPerimeter[key(r, c, potentialRegionID)]; ok {
		return
	}

	if regionIDs[r][c] == potentialRegionID {
		return
	}

	seenPerimeter[key(r, c, potentialRegionID)] = true

	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]
		if nr < 0 || nr >= rCount || nc < 0 || nc >= cCount {
			continue
		}

		floodFillSeenForPerimeter(grid, nr, nc, rCount, cCount, potentialRegionID, regionIDs, seenPerimeter)
	}
}

func floodFillArea(grid [][]string, r, c, rCount, cCount int, plot string, seen map[[2]int]bool, dataGrid [][]int, regionID int) int {
	if _, ok := seen[[2]int{r, c}]; ok {
		return 0
	}

	if grid[r][c] != plot {
		return 0
	}

	seen[[2]int{r, c}] = true
	dataGrid[r][c] = regionID

	result := 1

	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]
		if nr < 0 || nr >= rCount || nc < 0 || nc >= cCount {
			continue
		}

		result += floodFillArea(grid, nr, nc, rCount, cCount, plot, seen, dataGrid, regionID)
	}
	return result
}

var dirs [][2]int = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// 826904  - too low
// 860708  - wrong
// 2709122 - too high
// 3272136 - too high
// 1421940 - wrong
