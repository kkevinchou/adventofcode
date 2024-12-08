package main

import (
	"cmp"
	"container/heap"
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
	file := "input"
	generator := utils.RecordGenerator(file, "\n")

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

		line := record.Line

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

	var current [2]int
	var spans []*Span

	for _, instruction := range instructions {
		if instruction.Direction == UP {
			dist := instruction.Distance
			spans = append(spans, &Span{X: current[1], Start: current[0] - dist, End: current[0]})
			current[0] -= dist
		} else if instruction.Direction == DOWN {
			dist := instruction.Distance
			spans = append(spans, &Span{X: current[1], Start: current[0], End: current[0] + dist})
			current[0] += dist
		} else if instruction.Direction == RIGHT || instruction.Direction == LEFT {
			dist := instruction.Distance
			if instruction.Direction == LEFT {
				dist *= -1
			}
			current[1] += dist
		}
	}

	// var heapSpans Heap = spans
	// h := &heapSpans
	// heap.Init(h)

	// for h.Len() > 0 {
	// 	span := heap.Pop(h).(*Span)
	// 	fmt.Println("X", span.X, "START", span.Start, "END", span.End)
	// }

	totalArea := solve(spans)
	for _, instruction := range instructions {
		totalArea += instruction.Distance
	}
	fmt.Println(totalArea)
}

func testNoOverlap() {
	s1 := &Span{Start: -686074, End: -500255, X: 5411}
	s2 := &Span{Start: 0, End: 56407, X: 461937}

	ra, area := collide(s1, s2)
	fmt.Println(ra, area)
}

func testSimpleOverlap() {
	s1 := &Span{Start: 0, End: 100, X: 0}
	s2 := &Span{Start: 50, End: 200, X: 1}

	ra, area := collide(s1, s2)
	fmt.Println(ra, area)
}

func testFullOverlap() {
	s1 := &Span{Start: 1, End: 100, X: 0}
	s2 := &Span{Start: 1, End: 100, X: 1}

	ra, area := collide(s1, s2)
	fmt.Println(ra, area)
}

func test2() {
	s1 := &Span{Start: -686074, End: -500255, X: 5411}
	s2 := &Span{Start: 0, End: 56407, X: 461937}

	ra, area := collide(s1, s2)
	fmt.Println(ra, area)
}

func testCorner() {
	s1 := &Span{Start: 1, End: 2, X: 0}
	s2 := &Span{Start: 2, End: 3, X: 1}

	ra, area := collide(s1, s2)
	fmt.Println(ra, area)
}

func solve(spans []*Span) int {
	// initialization
	var heapSpans Heap = spans
	h := &heapSpans
	heap.Init(h)

	var totalArea int = 0

	aSpans := []*Span{heap.Pop(h).(*Span)}
	for h.Len() > 0 {
		bSpan := heap.Pop(h).(*Span)

		var nextaSpans []*Span
		for i, aSpan := range aSpans {
			aRemainder, area := collide(aSpan, bSpan)
			totalArea += area

			// collision detected
			if area > 0 {
				// add the new spans to the working set
				nextaSpans = append(nextaSpans, aRemainder...)
				// add the the remaining spans
				for j := i; j < len(aSpans); j++ {
					nextaSpans = append(nextaSpans, aSpans[j])
				}
				break
			} else {
				nextaSpans = append(nextaSpans, aSpan)
			}
		}
		aSpans = nextaSpans
	}

	return totalArea
}

func collideSpans(a []*Span, b *Span) []*Span {
	return nil
}

// new working set, area produced
// outside this code, if B did not collide with anything, add it to the working set
// otherwise it must have collided with something and collide will have produced the proper working set spans
func collide(a *Span, b *Span) ([]*Span, int) {
	// no collisions
	if a.Start > b.End || b.Start > a.End {
		return []*Span{a}, 0
	}

	// A matches B exactly

	// | - - |
	// |     |
	// | - - |
	//
	if a.Start == b.Start && a.End == b.End {
		return nil, (a.End - a.Start - 1) * (b.X - a.X - 1)
	}

	// we can have an extension of the span such that the overlap is more than just from
	// sharing the horizontal tunnel

	// handle span extensions

	// A   B
	// |
	// |
	// | - |
	//     |
	//     |

	if a.End == b.Start {
		// counts the area above the b span as well
		return []*Span{{X: b.X, Start: a.Start, End: b.End}}, (a.End - a.Start - 1) * (b.X - a.X)
	}

	// A   B
	//     |
	//     |
	// | - |
	// |
	// |

	if a.End == b.Start {
		// counts the area below the b span as well
		return []*Span{{X: b.X, Start: b.Start, End: a.End}}, (a.End - a.Start - 1) * (b.X - a.X)
	}

	// B collapses bottom of A

	// A     B
	// |
	// |     |
	// |     |
	// | - - |

	if a.Start < b.Start && a.End == b.End {
		return []*Span{{X: b.X, Start: a.Start, End: b.Start}}, (a.End-a.Start-1)*(b.X-a.X-1) + (b.Start - a.Start - 1)
	}

	// B collapses top of A

	// A     B
	// | - - |
	// |     |
	// |
	// |

	if a.Start == b.Start && a.End > b.End {
		return []*Span{{X: b.X, Start: b.End, End: a.End}}, (a.End-a.Start-1)*(b.X-a.X-1) + (a.End - b.End - 1)
	}

	// A covers bottom of B (impossible?)

	// A     B
	//       |
	//       |
	// |     |
	// | - - |

	if a.Start > b.Start && a.End == b.End {
		panic("A covers bottom of B")
	}

	// A fully covers B

	// A
	//
	// |
	// |
	// |  |
	// |  |
	// |
	// |

	if a.Start < b.Start && a.End > b.End {
		return []*Span{
			{X: b.X, Start: a.Start, End: b.Start},
			{X: b.X, Start: b.End, End: a.End},
		}, (a.End-a.Start-1)*(b.X-a.X-1) + (b.Start - a.Start - 1) + (a.End - b.End - 1)
	}

	panic("WAT")
}

func Min[T cmp.Ordered](a T, b T) T {
	if a <= b {
		return a
	}
	return b
}

func Max[T cmp.Ordered](a T, b T) T {
	if a >= b {
		return a
	}
	return b
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

type Heap []*Span

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].X < h[j].X }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Span))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
