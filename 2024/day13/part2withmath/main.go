package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"
var part2Offset int = 10000000000000

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
			x := utils.MustParseNum(strings.Trim(xSplit[1], ",")) + part2Offset
			y := utils.MustParseNum(ySplit[1]) + part2Offset

			determinant := (a[0]*b[1] - a[1]*b[0])
			if determinant != 0 {
				// solve system of linear equations
				A := (x*b[1] - y*b[0]) / (a[0]*b[1] - a[1]*b[0])
				B := (a[0]*y - a[1]*x) / (a[0]*b[1] - a[1]*b[0])

				checkValue := [2]int{A*a[0] + B*b[0], A*a[1] + B*b[1]}
				if checkValue[0] == x && checkValue[1] == y {
					cost := 3*A + B
					result += cost
				}
			}
		}
		step = (step + 1) % 3
	}
	fmt.Println(result)
}
