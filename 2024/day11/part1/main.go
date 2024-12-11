package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

type Stone struct {
	left  *Stone
	right *Stone
	value int
	dummy bool
}

func main() {
	var lineStr []string
	for record := range utils.Records(file) {
		lineStr = strings.Split(record.Line, " ")
	}

	leftDummy := &Stone{left: nil, right: nil, value: -1, dummy: true}
	rightDummy := &Stone{left: nil, right: nil, value: -1, dummy: true}

	lastStone := leftDummy
	for _, char := range lineStr {
		stone := &Stone{left: lastStone, right: nil, value: utils.MustParseNum(char)}
		lastStone.right = stone
		lastStone = stone
	}
	lastStone.right = rightDummy
	rightDummy.left = lastStone

	for _ = range 25 {
		blink(leftDummy)
	}

	fmt.Println(countStones(leftDummy))
}

func printStone(stone *Stone) {
	for stone != nil {
		fmt.Printf("%d ", stone.value)
		stone = stone.right
	}
	fmt.Println()
}

func countStones(stone *Stone) int {
	count := 0
	for stone != nil {
		if !stone.dummy {
			count++
		}
		stone = stone.right
	}
	return count
}

func blink(stone *Stone) {
	for stone != nil {
		if stone.dummy {
			stone = stone.right
			continue
		}
		var nextStone *Stone
		if stone.value == 0 {
			nextStone = stone.right
			stone.value = 1
		} else if l, r, ok := eventDigit(stone.value); ok {
			nextStone = stone.right

			leftStone := &Stone{left: stone.left, value: l, dummy: false}
			rightStone := &Stone{right: stone.right, value: r, dummy: false}

			stone.left.right = leftStone
			stone.right.left = rightStone

			leftStone.right = rightStone
			rightStone.left = leftStone
		} else {
			nextStone = stone.right
			stone.value *= 2024
		}

		stone = nextStone
	}
}

func eventDigit(value int) (int, int, bool) {
	s := fmt.Sprintf("%d", value)
	if len(s)%2 == 0 {
		left := s[:len(s)/2]
		right := s[len(s)/2:]
		return utils.MustParseNum(left), utils.MustParseNum(right), true
	}
	return -1, -1, false
}
