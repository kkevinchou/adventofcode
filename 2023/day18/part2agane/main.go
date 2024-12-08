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

	// lastDir := [2]int{}
	// for _, i := range instructions {
	// 	if i.Direction == lastDir {
	// 		fmt.Println("HI")
	// 	}
	// 	lastDir = i.Direction
	// 	fmt.Println(i)
	// }

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
		fmt.Println(instruction.Direction, instruction.Distance)
	}

	// var heapSpans Heap = spans
	// h := &heapSpans
	// heap.Init(h)

	// for h.Len() > 0 {
	// 	span := heap.Pop(h).(*Span)
	// 	fmt.Println("X", span.X, "START", span.Start, "END", span.End)
	// }

	// testSimpleOverlap()

	totalArea := solve(spans)
	fmt.Println(totalArea)
}

func testNoOverlap() {
	s1 := &Span{Start: -686074, End: -500255, X: 5411}
	s2 := &Span{Start: 0, End: 56407, X: 461937}

	ra, rb, area := collide(s1, s2)
	fmt.Println(ra, rb, area)
}

func testSimpleOverlap() {
	s1 := &Span{Start: 0, End: 100, X: 0}
	s2 := &Span{Start: 50, End: 200, X: 1}

	ra, rb, area := collide(s1, s2)
	fmt.Println(ra, rb, area)
}

func test2() {
	s1 := &Span{Start: -686074, End: -500255, X: 5411}
	s2 := &Span{Start: 0, End: 56407, X: 461937}

	ra, rb, area := collide(s1, s2)
	fmt.Println(ra, rb, area)
}

func solve(spans []*Span) int {
	// initialization

	var heapSpans Heap = spans
	h := &heapSpans
	heap.Init(h)

	var aSpans []*Span
	var totalArea int = 0

	for h.Len() > 0 {
		// [ A B C ] ----- [ D E F ]

		// A iteration:
		// [ A ] ----- [ D E F ]
		// [ A ] ----- [ D ]
		// [ A0, A1 ] ----- [ E ]
		// [ A ] ----- [ F ]

		// pop off a span
		bSpans := []*Span{heap.Pop(h).(*Span)}
		var nextWorkingSet []*Span

		for _, aSpan := range aSpans {
			// THIS IS THE BEGINNING OF PROCESSING A

			aSubSpans := []*Span{aSpan}
			bSubSpans := bSpans

			// [ A1 A2 A3 ] ---- [ B1 B2 B3 ]

			var nextbSpans []*Span
			for _, bSpan := range bSubSpans {
				// THIS IS THE BEGINNING OF PROCESSING B
				var nextaSubSpans []*Span
				for _, aSubSpan := range aSubSpans {
					aRemainderFromCollision, bRemainderFromCollision, area := collide(aSubSpan, bSpan)
					nextaSubSpans = append(nextaSubSpans, aRemainderFromCollision...)
					nextbSpans = append(nextbSpans, bRemainderFromCollision...)
					if area < 0 {
						fmt.Println("HI")
					}
					totalArea += area
					fmt.Println(area)
				}
				aSubSpans = nextaSubSpans
			}

			bSpans = nextbSpans
			nextWorkingSet = append(nextWorkingSet, aSubSpans...)
		}
		aSpans = nextWorkingSet
		aSpans = append(aSpans, bSpans...)
	}

	return totalArea
}

func collideSpans(a []*Span, b *Span) []*Span {
	return nil
}

func collide(a *Span, b *Span) ([]*Span, []*Span, int) {
	var aRemainder []*Span
	var bRemainder []*Span

	// A covers top of B

	// A B
	// |
	// | |
	// | |
	//   |

	if a.Start < b.Start {
		end := Min(a.End, b.Start-1)
		aRemainder = append(aRemainder, &Span{X: b.X, Start: a.Start, End: end})
	}

	// A covers bottom of B

	// A B
	//   |
	// | |
	// | |
	// |

	if a.End > b.End {
		start := Max(a.Start, b.End+1)
		aRemainder = append(aRemainder, &Span{X: b.X, Start: start, End: a.End})
	}

	// same as above, but for b

	// B covers top of A

	// A B
	//   |
	// | |
	// | |
	// |

	if b.Start < a.Start {
		end := Min(b.End, a.Start-1)
		bRemainder = append(bRemainder, &Span{X: b.X, Start: b.Start, End: end})
	}

	// B covers bottom of A

	// A B
	// |
	// | |
	// | |
	//   |

	if b.End > a.End {
		start := Max(b.Start, a.End+1)
		bRemainder = append(bRemainder, &Span{X: b.X, Start: start, End: b.End})
	}

	// calculate vertical overlap, this is used to compute the amount of area it covers

	var area int
	if !(a.End < b.Start || b.End < a.Start) {
		overlapStart := Max(a.Start, b.Start)
		overlapEnd := Min(a.End, b.End)
		overlapLength := overlapEnd - overlapStart + 1
		area = (b.X - a.X + 1) * overlapLength
	}

	return aRemainder, bRemainder, area
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
