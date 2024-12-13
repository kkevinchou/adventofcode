package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

func main() {
	var a [2]int
	var b [2]int

	var result int
	step := 0
	for record := range utils.Records(file) {
		if record.Line == "" {
			continue
		}

		if step == 0 {
			lineSplit := strings.Split(record.Line, " ")
			xSplit := strings.Split(lineSplit[2], "+")
			ySplit := strings.Split(lineSplit[3], "+")
			x := utils.MustParseNum(strings.Trim(xSplit[1], ","))
			y := utils.MustParseNum(ySplit[1])
			a[0], a[1] = x, y
		} else if step == 1 {
			lineSplit := strings.Split(record.Line, " ")
			xSplit := strings.Split(lineSplit[2], "+")
			ySplit := strings.Split(lineSplit[3], "+")
			x := utils.MustParseNum(strings.Trim(xSplit[1], ","))
			y := utils.MustParseNum(ySplit[1])
			b[0], b[1] = x, y
		} else if step == 2 {
			lineSplit := strings.Split(record.Line, " ")
			xSplit := strings.Split(lineSplit[1], "=")
			ySplit := strings.Split(lineSplit[2], "=")
			x := utils.MustParseNum(strings.Trim(xSplit[1], ","))
			y := utils.MustParseNum(ySplit[1])

			cache := map[[2]int]int{}
			cost, possible := solve(a, b, x, y, cache)
			if possible {
				// fmt.Println(cost)
				result += cost
			}
		}

		step = (step + 1) % 3
	}
	fmt.Println(result)
}

func solve(a [2]int, b [2]int, goalX, goalY int, cache map[[2]int]int) (int, bool) {
	if cost, ok := cache[[2]int{goalX, goalY}]; ok {
		if cost >= 0 {
			return cost, true
		} else {
			return -1, false
		}
	}

	if goalX < 0 || goalY < 0 {
		return -1, false
	}

	if goalX == 0 && goalY == 0 {
		return 0, true
	}

	var cost int

	// pick a
	aCost, aPossible := solve(a, b, goalX-a[0], goalY-a[1], cache)

	// pick b
	bCost, bPossible := solve(a, b, goalX-b[0], goalY-b[1], cache)

	if !aPossible && !bPossible {
		cache[[2]int{goalX, goalY}] = -1
		return -1, false
	} else if aPossible && bPossible {
		cost = min((3 + aCost), (1 + bCost))
	} else if aPossible {
		cost = 3 + aCost
	} else if bPossible {
		cost = 1 + bCost
	}

	cache[[2]int{goalX, goalY}] = cost

	return cost, true
}
