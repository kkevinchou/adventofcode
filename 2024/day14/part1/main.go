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
	quadrants := [4]int{}
	for _, robot := range robots {
		for _ = range 100 {
			robot.r += robot.vr
			robot.c += robot.vc

			robot.r = (robot.r + rCount) % rCount
			robot.c = (robot.c + cCount) % cCount
		}
		if quadrant, ok := quadrant(robot, rCount, cCount); ok {
			quadrants[quadrant]++
		}
	}

	fmt.Println(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])
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
