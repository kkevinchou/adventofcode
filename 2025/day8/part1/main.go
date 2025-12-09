package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var file string = "input.txt"

type DistancePair struct {
	P0          Point
	P1          Point
	DistSquared int
}

type Point struct {
	ID      int
	X, Y, Z int
}

func key(id0, id1 int) string {
	return fmt.Sprintf("%d_%d", min(id0, id1), max(id0, id1))
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")
	var points []Point
	for _, line := range lines {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		z, _ := strconv.Atoi(split[2])

		points = append(points, Point{X: x, Y: y, Z: z, ID: len(points)})
	}

	start2 := time.Now()
	pairs := make([]DistancePair, 0, len(points)*(len(points)-1)/2)
	for i := range len(points) {
		p0 := points[i]
		for j := i + 1; j < len(points); j++ {
			p1 := points[j]
			dx := p0.X - p1.X
			dy := p0.Y - p1.Y
			dz := p0.Z - p1.Z

			distSquared := dx*dx + dy*dy + dz*dz
			pairs = append(pairs, DistancePair{P0: p0, P1: p1, DistSquared: distSquared})
		}
	}
	fmt.Println("pair creation", time.Since(start2))

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].DistSquared < pairs[j].DistSquared
	})

	idToCircuit := map[int]int{}
	circuitToIDs := [][]int{}
	for i := 0; i < 1000; i++ {
		pair := pairs[i]
		p0 := pair.P0.ID
		p1 := pair.P1.ID

		var p0Circuit = -1
		if id, ok := idToCircuit[p0]; ok {
			p0Circuit = id
		}

		var p1Circuit = -1
		if id, ok := idToCircuit[p1]; ok {
			p1Circuit = id
		}

		if p0Circuit == -1 && p1Circuit == -1 {
			circuitToIDs = append(circuitToIDs, []int{p0, p1})
			idToCircuit[p0] = len(circuitToIDs) - 1
			idToCircuit[p1] = len(circuitToIDs) - 1
		} else if p0Circuit >= 0 && p1Circuit >= 0 {
			if p0Circuit != p1Circuit {
				target := min(p0Circuit, p1Circuit)
				source := max(p0Circuit, p1Circuit)

				for _, id := range circuitToIDs[source] {
					idToCircuit[id] = target
				}

				circuitToIDs[target] = append(circuitToIDs[target], circuitToIDs[source]...)
				circuitToIDs[source] = []int{}
			}
		} else if p0Circuit == -1 {
			circuitToIDs[p1Circuit] = append(circuitToIDs[p1Circuit], p0)
			idToCircuit[p0] = p1Circuit
		} else if p1Circuit == -1 {
			circuitToIDs[p0Circuit] = append(circuitToIDs[p0Circuit], p1)
			idToCircuit[p1] = p0Circuit
		}
	}

	var sizes []int
	for _, ids := range circuitToIDs {
		sizes = append(sizes, len(ids))
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	result := 1
	for i := 0; i < 3; i++ {
		result *= sizes[i]
	}

	fmt.Println(result)
	fmt.Println(time.Since(start))
}
