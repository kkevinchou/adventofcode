package main

import (
	"fmt"
	"math"
)

var directionIndex int
var directions = []string{"E", "S", "W", "N"}

func main() {
	var x, y int
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
			x += num
		} else if operation == "S" {
			y -= num
		} else if operation == "W" {
			x -= num
		} else if operation == "N" {
			y += num
		} else if operation == "L" {
			directionIndex -= (num / 90)
			if directionIndex < 0 {
				fmt.Println("HI1", directionIndex, operation, num)
				directionIndex = directionIndex % len(directions)
				directionIndex = directionIndex + len(directions)
				fmt.Println("HI2", directionIndex, operation, num)
			}
		} else if operation == "R" {
			directionIndex += (num / 90)
			directionIndex = directionIndex % len(directions)
		} else if operation == "F" {
			if directions[directionIndex] == "E" {
				x += num
			} else if directions[directionIndex] == "S" {
				y -= num
			} else if directions[directionIndex] == "W" {
				x -= num
			} else if directions[directionIndex] == "N" {
				y += num
			}
		}
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}
