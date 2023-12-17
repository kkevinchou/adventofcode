package main

import (
	"bytes"
	"container/heap"
	"encoding/binary"
	"fmt"
	"os"
	"runtime/pprof"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	profileFile, err := os.Create("profile.pprof")
	if err != nil {
		fmt.Println("Error creating profile file:", err)
		return
	}
	defer profileFile.Close()

	if err := pprof.StartCPUProfile(profileFile); err != nil {
		fmt.Println("Error starting CPU profile:", err)
		return
	}
	defer pprof.StopCPUProfile()

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

	goal := [2]int{maxRow - 1, maxCol - 1}

	result := solve(grid, goal, maxRow, maxCol)

	// for _, cell := range result.PathSoFar {
	// 	stringGrid[cell[0]][cell[1]] = "#"
	// }
	// for _, row := range stringGrid {
	// 	fmt.Println(row)
	// }

	fmt.Println(result.Score)
	fmt.Println(time.Since(startTime))
	pprof.StopCPUProfile()
}

func solve(grid [][]int, goal [2]int, maxRow, maxCol int) *Candidate {
	h := &Heap{}

	startCandidate1 := &Candidate{Position: [2]int{0, 1}, IncomingDirection: [2]int{0, 1}, IncomingDirectionCount: 1, Score: grid[0][1]}
	startCandidate2 := &Candidate{Position: [2]int{1, 0}, IncomingDirection: [2]int{1, 0}, IncomingDirectionCount: 1, Score: grid[1][0]}

	h.Push(startCandidate1)
	h.Push(startCandidate2)

	cache := map[[6]byte]int{}
	cache[genKey(startCandidate1)] = 1
	cache[genKey(startCandidate2)] = 1

	for {
		candidate := heap.Pop(h).(*Candidate)
		if candidate.Position == goal && candidate.IncomingDirectionCount >= 4 {
			return candidate
		}

		r, c := candidate.Position[0], candidate.Position[1]
		neighbors := getNeighbors(r, c, maxRow, maxCol)

		for _, neighbor := range neighbors {
			neighborDir := Direction{neighbor[0] - r, neighbor[1] - c}

			// var pathToNeighbor [][2]int
			// for _, pathCell := range candidate.PathSoFar {
			// 	pathToNeighbor = append(pathToNeighbor, pathCell)
			// }
			// pathToNeighbor = append(pathToNeighbor, neighbor)

			if candidate.IncomingDirectionCount < 4 {
				if neighborDir != candidate.IncomingDirection {
					continue
				}
			}

			if candidate.IncomingDirectionCount == 10 {
				if neighborDir == candidate.IncomingDirection {
					continue
				}
			}

			if candidate.IncomingDirection == neighborDir {
				// forward case
				neighborCandidate := &Candidate{
					Position:               neighbor,
					IncomingDirection:      neighborDir,
					IncomingDirectionCount: candidate.IncomingDirectionCount + 1,
					Score:                  candidate.Score + grid[neighbor[0]][neighbor[1]],
					// PathSoFar:              pathToNeighbor,
				}

				neighborCandidateKey := genKey(neighborCandidate)
				if _, ok := cache[neighborCandidateKey]; !ok {
					heap.Push(h, neighborCandidate)
					cache[neighborCandidateKey] = 1
				}
			} else if neighborDir == dirToBack[candidate.IncomingDirection] {
				// back case
				continue
			} else {
				// left and right case
				neighborCandidate := &Candidate{
					Position:               neighbor,
					IncomingDirection:      neighborDir,
					IncomingDirectionCount: 1,
					Score:                  candidate.Score + grid[neighbor[0]][neighbor[1]],
					// PathSoFar:              pathToNeighbor,
				}
				neighborCandidateKey := genKey(neighborCandidate)
				if _, ok := cache[neighborCandidateKey]; !ok {
					heap.Push(h, neighborCandidate)
					cache[neighborCandidateKey] = 1
				}
			}
		}
	}

}

func genKey(candidate *Candidate) [6]byte {
	position := candidate.Position
	direction := candidate.IncomingDirection
	directionCount := candidate.IncomingDirectionCount

	// Create a buffer to store the byte array
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, uint8(position[0]))
	if err != nil {
		fmt.Println("Error writing binary:", err)
	}

	err = binary.Write(buf, binary.LittleEndian, uint8(position[1]))
	if err != nil {
		fmt.Println("Error writing binary:", err)
	}

	err = binary.Write(buf, binary.LittleEndian, int8(direction[0]))
	if err != nil {
		fmt.Println("Error writing binary:", err)
	}

	err = binary.Write(buf, binary.LittleEndian, int8(direction[1]))
	if err != nil {
		fmt.Println("Error writing binary:", err)
	}

	err = binary.Write(buf, binary.LittleEndian, uint16(directionCount))
	if err != nil {
		fmt.Println("Error writing binary:", err)
	}

	var bytes [6]byte

	for i, b := range buf.Bytes() {
		bytes[i] = b
	}

	return bytes

	// return 2 * (position[0] + 2) * 3 * (position[1] + 2) * 5 * (direction[0] + 2) * 7 * (direction[1]) * 11 * (directionCount + 2)
	// return fmt.Sprintf("%d_%d_%d_%d_%d", position[0], position[1], direction[0], direction[1], directionCount)
}

// TODO - do i need to consider if it's possible to re-enter the same cell, but at a lower cost?

type Candidate struct {
	Position               [2]int
	IncomingDirection      Direction
	IncomingDirectionCount int
	Score                  int
	PathSoFar              [][2]int
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
