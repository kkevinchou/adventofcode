package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"
var regionIDGen int

func main() {
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

	for r := range rCount {
		for c := range cCount {
			if _, ok := seen[[2]int{r, c}]; ok {
				continue
			}

			area := floodFillArea(grid, r, c, rCount, cCount, grid[r][c], seen, regionIDs, regionIDGen)
			areas[regionIDGen] = area
			regionIDGen++
		}
	}

	top := sweepTopDown(regionIDs, rCount, cCount, [2]int{-1, 0})
	bottom := sweepTopDown(regionIDs, rCount, cCount, [2]int{1, 0})
	left := sweepLeftRight(regionIDs, rCount, cCount, [2]int{0, -1})
	right := sweepLeftRight(regionIDs, rCount, cCount, [2]int{0, 1})

	sides := map[int]int{}
	for regionID, sideCount := range top {
		sides[regionID] += sideCount
	}
	for regionID, sideCount := range bottom {
		sides[regionID] += sideCount
	}
	for regionID, sideCount := range left {
		sides[regionID] += sideCount
	}
	for regionID, sideCount := range right {
		sides[regionID] += sideCount
	}

	var result int
	for regionID, sideCount := range sides {
		result += areas[regionID] * sideCount
	}

	fmt.Println(result)
}

func sweepTopDown(regionIDs [][]int, rCount, cCount int, dir [2]int) map[int]int {
	// top down
	counts := map[int]int{}
	for r := range rCount {
		lastRegionID := -1
		for c := range cCount {
			regionID := regionIDs[r][c]
			if edgeCheck(regionIDs, dir, r, c, rCount, cCount) {
				if regionID != lastRegionID {
					counts[regionID]++
					lastRegionID = regionID
				}
			} else {
				lastRegionID = -1
			}
		}
	}
	return counts
}

func sweepLeftRight(regionIDs [][]int, rCount, cCount int, dir [2]int) map[int]int {
	// top down
	counts := map[int]int{}
	for c := range cCount {
		lastRegionID := -1
		for r := range rCount {
			regionID := regionIDs[r][c]
			if edgeCheck(regionIDs, dir, r, c, rCount, cCount) {
				if regionID != lastRegionID {
					counts[regionID]++
					lastRegionID = regionID
				}
			} else {
				lastRegionID = -1
			}
		}
	}
	return counts
}

func edgeCheck(regionIDs [][]int, dir [2]int, r, c, rCount, cCount int) bool {
	nr, nc := r+dir[0], c+dir[1]
	if nr < 0 || nr >= rCount || nc < 0 || nc >= cCount {
		return true
	}

	return regionIDs[r][c] != regionIDs[nr][nc]
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

var dirs [][2]int = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}
