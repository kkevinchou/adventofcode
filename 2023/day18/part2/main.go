package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

type Span struct {
	Start int
	End   int

	X int
}

func main() {
	file := "sample"
	generator := utils.RecordGenerator(file, "\r\n")

	dirMap := map[string]Direction{
		"0": RIGHT,
		"1": DOWN,
		"2": LEFT,
		"3": UP,
	}

	var instructions []Instruction

	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine

		lineSplit := strings.Split(line, " ")
		hex := lineSplit[2][1 : len(lineSplit[2])-1]

		distanceHex := hex[:len(hex)-1]
		distanceHex = distanceHex[1:]
		directionHex := hex[len(hex)-1:]

		distance64, err := strconv.ParseInt(distanceHex, 16, 64)
		if err != nil {
			panic(err)
		}

		// fmt.Println(distance64, directionHex)
		instructions = append(instructions, Instruction{Direction: dirMap[directionHex], Distance: int(distance64)})
	}

	fmt.Println(instructions)
	// walk the instructions and create each vertical span
	// call solve withspans

}

func solve(spans []*Span) {
	// insert all spans into the min heap based on X coordinate
	// pop off a span
	// collide against working set of spans
	// update working set of spans based on collisions
	//		for each span (A) from the working set, collide against the popped off span (B)
	// 		if there is a remainder left to span B, add it to the working set
	// 		sum the consumed area
	// repeat
}

func collide(a *Span, b *Span) ([]*Span, []*Span, int) {
	// no overlaps
	if a.End < b.Start {
		return []*Span{a}, []*Span{b}, 0
	}

	if b.End < a.Start {
		return []*Span{a}, []*Span{b}, 0
	}

	// a fully covers b
	if a.Start < b.Start && a.End > b.End {
		return []*Span{
			{X: b.X, Start: a.Start, End: b.Start - 1},
			{X: b.X, Start: b.End + 1, End: a.End},
		}, nil, (b.X - a.X) * (b.End - b.Start + 1)
	}

	// b fully covers a

	// a covers top of b

	// a covers bottom of b

	return nil, nil, 0
}

type Instruction struct {
	Direction Direction
	Distance  int
}

type Direction [2]int

var LEFT Direction = Direction{0, -1}
var RIGHT Direction = Direction{0, 1}
var UP Direction = Direction{-1, 0}
var DOWN Direction = Direction{1, 0}
