package main

import (
	"fmt"
	"math"
)

var directionIndex int
var directions = []string{"E", "S", "W", "N"}

func main() {
	var x, y int
	var wayx, wayy = 10, 1
	directionIndex = 0

	recordGenerator := constructRecordGenerator("input", "\n")
	for {
		record, done := recordGenerator()
		if done {
			break
		}
		instruction := record.Lines[0]
		operation := string(instruction[0])
		num := mustParseNum(instruction[1:])

		if operation == "E" {
			wayx += num
		} else if operation == "S" {
			wayy -= num
		} else if operation == "W" {
			wayx -= num
		} else if operation == "N" {
			wayy += num
		} else if operation == "L" {
			rotationCount := num % 360

			for i := 0; i < rotationCount/90; i++ {
				wayxOld := wayx
				wayyOld := wayy

				wayx = wayyOld * -1
				wayy = wayxOld
			}
		} else if operation == "R" {
			rotationCount := num % 360

			for i := 0; i < rotationCount/90; i++ {
				wayxOld := wayx
				wayyOld := wayy

				wayx = wayyOld
				wayy = wayxOld * -1
			}
		} else if operation == "F" {

			for i := 0; i < num; i++ {
				x += wayx
				y += wayy
			}

			// if directions[directionIndex] == "E" {
			// 	x += num
			// 	wayx += num
			// } else if directions[directionIndex] == "S" {
			// 	y -= num
			// 	wayy -= num
			// } else if directions[directionIndex] == "W" {
			// 	x -= num
			// 	wayx -= num
			// } else if directions[directionIndex] == "N" {
			// 	y += num
			// 	wayy += num
			// }
		}
		fmt.Println(x, y)
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}
