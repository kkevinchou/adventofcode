package main

import (
	"container/heap"
	"fmt"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	startTime := time.Now()
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var grid [][]int
	var stringGrid [][]string
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine
		grid = append(grid, utils.StringSliceToIntSlice(utils.StringToStringSlice(line)))
		stringGrid = append(stringGrid, utils.StringToStringSlice(line))
	}

	maxRow, maxCol := len(grid), len(grid[0])

	start := [2]int{0, 0}
	goal := [2]int{maxRow - 1, maxCol - 1}

	result := solve(grid, start, goal, maxRow, maxCol)
	fmt.Println(result.Score)
	fmt.Println(time.Since(startTime))
}

func solve(grid [][]int, start, goal [2]int, maxRow, maxCol int) *Candidate {
	h := &Heap{}

	startCandidate := &Candidate{Position: start}
	h.Push(startCandidate)

	cache := map[string]bool{}
	cache[genKey(startCandidate)] = true

	for {
		candidate := heap.Pop(h).(*Candidate)
		if candidate.Position == goal {
			return candidate
		}

		r, c := candidate.Position[0], candidate.Position[1]
		neighbors := getNeighbors(r, c, maxRow, maxCol)

		for _, neighbor := range neighbors {
			neighborDir := Direction{neighbor[0] - r, neighbor[1] - c}

			if candidate.IncomingDirection == neighborDir {
				if candidate.IncomingDirectionCount < 3 {
					// forward case
					neighborCandidate := &Candidate{
						Position:               neighbor,
						IncomingDirection:      neighborDir,
						IncomingDirectionCount: candidate.IncomingDirectionCount + 1,
						Score:                  candidate.Score + grid[neighbor[0]][neighbor[1]],
					}

					neighborCandidateKey := genKey(neighborCandidate)
					if _, ok := cache[neighborCandidateKey]; !ok {
						heap.Push(h, neighborCandidate)
						cache[neighborCandidateKey] = true
					}
				}
			} else if neighborDir == dirToBack[candidate.IncomingDirection] {
				// backwards case
				continue
			} else {
				// left and right case
				neighborCandidate := &Candidate{
					Position:               neighbor,
					IncomingDirection:      neighborDir,
					IncomingDirectionCount: 1,
					Score:                  candidate.Score + grid[neighbor[0]][neighbor[1]],
				}
				neighborCandidateKey := genKey(neighborCandidate)
				if _, ok := cache[neighborCandidateKey]; !ok {
					heap.Push(h, neighborCandidate)
					cache[neighborCandidateKey] = true
				}
			}
		}
	}
}

func genKey(candidate *Candidate) string {
	position := candidate.Position
	direction := candidate.IncomingDirection
	directionCount := candidate.IncomingDirectionCount
	return fmt.Sprintf("%d_%d_%d_%d_%d", position[0], position[1], direction[0], direction[1], directionCount)
}

// TODO - do i need to consider if it's possible to re-enter the same cell, but at a lower cost?

type Candidate struct {
	Position               [2]int
	IncomingDirection      Direction
	IncomingDirectionCount int
	Score                  int
}

type Heap []*Candidate

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Score < h[j].Score }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Candidate))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getNeighbors(r, c, maxRow, maxCol int) [][2]int {
	var result [][2]int
	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]
		if nr < 0 || nr >= maxRow {
			continue
		}

		if nc < 0 || nc >= maxCol {
			continue
		}

		result = append(result, [2]int{nr, nc})
	}
	return result
}

var dirs [][2]int = [][2]int{
	{1, 0},
	{0, -1},
	{-1, 0},
	{0, 1},
}

var dirToLeft map[[2]int][2]int = map[[2]int][2]int{
	[2]int{1, 0}:  [2]int{0, 1},
	[2]int{0, -1}: [2]int{1, 0},
	[2]int{-1, 0}: [2]int{0, -1},
	[2]int{0, 1}:  [2]int{-1, 0},
}

var dirToRight map[[2]int][2]int = map[[2]int][2]int{
	[2]int{1, 0}:  [2]int{0, -1},
	[2]int{0, -1}: [2]int{-1, 0},
	[2]int{-1, 0}: [2]int{0, 1},
	[2]int{0, 1}:  [2]int{1, 0},
}

var dirToBack map[[2]int][2]int = map[[2]int][2]int{
	[2]int{1, 0}:  [2]int{-1, 0},
	[2]int{0, -1}: [2]int{0, 1},
	[2]int{-1, 0}: [2]int{1, 0},
	[2]int{0, 1}:  [2]int{0, -1},
}

type Direction [2]int

var LEFT Direction = Direction{0, -1}
var RIGHT Direction = Direction{0, 1}
var UP Direction = Direction{-1, 0}
var DOWN Direction = Direction{1, 0}
