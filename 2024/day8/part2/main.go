package main

import (
	"fmt"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

type Coordinate [2]int

func main() {
	start := time.Now()
	grid, rCount, cCount := utils.ParseGrid(file)

	signals := map[string][]Coordinate{}
	for r := range rCount {
		for c := range cCount {
			if grid[r][c] == "." {
				continue
			}
			signals[grid[r][c]] = append(signals[grid[r][c]], Coordinate{r, c})
		}
	}

	result := map[Coordinate]bool{}
	for _, signalList := range signals {
		coords := genCoords(signalList, rCount, cCount)
		for coord := range coords {
			result[coord] = true
		}
	}
	fmt.Println(len(result))
	fmt.Println(time.Since(start))
}

func genCoords(signals []Coordinate, rCount, cCount int) map[Coordinate]bool {
	result := map[Coordinate]bool{}

	for i := range signals {
		iSignal := signals[i]
		for j := range signals {
			jSignal := signals[j]
			if i == j {
				continue
			}

			dr := jSignal[0] - iSignal[0]
			dc := jSignal[1] - iSignal[1]

			curR := jSignal[0]
			curC := jSignal[1]

			for {
				curR += dr
				curC += dc
				antinode := Coordinate{curR, curC}
				if antinode[0] < 0 || antinode[0] >= rCount {
					break
				}
				if antinode[1] < 0 || antinode[1] >= cCount {
					break
				}
				result[antinode] = true
			}

			curR = iSignal[0]
			curC = iSignal[1]
			for {
				curR -= dr
				curC -= dc
				antinode := Coordinate{curR, curC}
				if antinode[0] < 0 || antinode[0] >= rCount {
					break
				}
				if antinode[1] < 0 || antinode[1] >= cCount {
					break
				}
				result[antinode] = true
			}

			result[iSignal] = true
			result[jSignal] = true
		}
	}
	return result
}
