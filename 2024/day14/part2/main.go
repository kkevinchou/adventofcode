package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

type Robot struct {
	r  int
	c  int
	vr int
	vc int
}

func main() {
	var robots []*Robot

	for record := range utils.Records(file) {
		spaceSplit := strings.Split(record.Line, " ")
		positionSlice := strings.Split(strings.Trim(spaceSplit[0], "p="), ",")
		position := utils.StringSliceToIntSlice(positionSlice)
		c := position[0]
		r := position[1]

		velocitySlice := strings.Split(strings.Trim(spaceSplit[1], "v="), ",")
		velocity := utils.StringSliceToIntSlice(velocitySlice)
		vc := velocity[0]
		vr := velocity[1]

		robots = append(robots, &Robot{r: r, c: c, vr: vr, vc: vc})
	}

	rCount, cCount := 103, 101
	for j := range 10001 {
		for _, robot := range robots {
			robot.r += robot.vr
			robot.c += robot.vc

			robot.r = (robot.r + rCount) % rCount
			robot.c = (robot.c + cCount) % cCount
		}

		grid := make([][]string, rCount)
		for i := range rCount {
			grid[i] = make([]string, cCount)
		}
		for r := range rCount {
			for c := range cCount {
				grid[r][c] = " "
			}
		}
		for _, robot := range robots {
			grid[robot.r][robot.c] = "O"
		}

		if check(grid, rCount, cCount) {
			utils.PrintGrid(grid)
			fmt.Println(j+1, "-------------------------")
		}
	}

	// fmt.Println(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])
}

func check(grid [][]string, rCount, cCount int) bool {
	lineSize := 5

	for r := range rCount {
		for c := range cCount {
			if c+lineSize >= cCount {
				continue
			}

			horizontal := true
			nr := r
			for offset := range lineSize {
				nc := c + offset + 1

				if grid[nr][nc] != "O" {
					horizontal = false
				}
			}
			if horizontal {
				return true
			}
		}
	}
	return false
}

func quadrant(robot *Robot, rCount, cCount int) (int, bool) {
	if robot.r == rCount/2 || robot.c == cCount/2 {
		return -1, false
	}

	if robot.r < rCount/2 {
		if robot.c < cCount/2 {
			return 0, true
		} else {
			return 1, true
		}
	} else {
		if robot.c < cCount/2 {
			return 2, true
		} else {
			return 3, true
		}
	}
}

// 240346656 - wrong
