package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

var file string = "input.txt"

type Point struct {
	id     int
	r, c   int
	cr, cc int
}

func main() {
	start := time.Now()
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")
	var points []Point
	for i, line := range lines {
		var r, c int
		fmt.Sscanf(line, "%d,%d", &c, &r)
		points = append(points, Point{r: r, c: c, id: i})
	}

	// compress r
	sort.Slice(points, func(i, j int) bool {
		return points[i].r < points[j].r
	})
	rMap := map[int]int{}
	prev := -1
	for i := range len(points) {
		var compressed int
		if points[i].r != prev {
			compressed = len(rMap)
			rMap[points[i].r] = len(rMap)
		} else {
			compressed = rMap[points[i].r]
		}
		points[i].cr = compressed
		prev = points[i].r
	}

	// compress c
	sort.Slice(points, func(i, j int) bool {
		return points[i].c < points[j].c
	})
	cMap := map[int]int{}
	prev = -1
	for i := range len(points) {
		var compressed int
		if points[i].c != prev {
			compressed = len(cMap)
			cMap[points[i].c] = len(cMap)
		} else {
			compressed = cMap[points[i].c]
		}
		points[i].cc = compressed
		prev = points[i].c
	}

	maxr, maxc := 0, 0
	for _, p := range points {
		maxr = max(maxr, p.cr)
		maxc = max(maxc, p.cc)
	}

	grid := [][]int{}
	for range maxr + 1 {
		grid = append(grid, make([]int, maxc+1))
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i].id < points[j].id
	})

	// fill in the shape perimeter
	for i, a := range points {
		b := points[(i+1)%len(points)]

		if b.cr != a.cr {
			startr := min(b.cr, a.cr)
			maxr := max(b.cr, a.cr)

			for r := startr; r <= maxr; r++ {
				grid[r][a.cc] = 1
			}
		}

		if b.cc != a.cc {
			startc := min(b.cc, a.cc)
			maxc := max(b.cc, a.cc)

			for c := startc; c <= maxc; c++ {
				grid[a.cr][c] = 1
			}
		}
	}

	// flood fill the empty space outside of the perimeter
	seen := map[string]bool{}
	rows := []int{0, maxr}
	cols := []int{0, maxc}
	for _, r := range rows {
		for c := 0; c <= maxc; c++ {
			flood(r, c, grid, seen)
		}
	}
	for _, c := range cols {
		for r := 0; r <= maxr; r++ {
			flood(r, c, grid, seen)
		}
	}

	// find max area
	area := math.MinInt
	for i := 0; i < len(points); i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]

			minr, maxr := min(a.cr, b.cr), max(a.cr, b.cr)
			minc, maxc := min(a.cc, b.cc), max(a.cc, b.cc)

			valid := true
			for r := minr; r <= maxr; r++ {
				for c := minc; c <= maxc; c++ {
					if grid[r][c] == -1 {
						valid = false
						break
					}
				}
				if !valid {
					break
				}
			}

			if valid {
				currArea := (abs(a.r-b.r) + 1) * (abs(a.c-b.c) + 1)
				area = max(area, currArea)
			}
		}
	}

	fmt.Println(area)
	fmt.Println(time.Since(start))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func flood(r, c int, grid [][]int, seen map[string]bool) {
	key := fmt.Sprintf("%d_%d", r, c)
	if seen[key] {
		return
	}
	seen[key] = true

	if grid[r][c] == 1 || grid[r][c] == -1 {
		return
	}
	grid[r][c] = -1

	for _, rdir := range dirs {
		for _, cdir := range dirs {
			if rdir == cdir {
				continue
			}

			nr := r + rdir
			nc := c + cdir

			if nr < 0 || nr >= len(grid) {
				continue
			}

			if nc < 0 || nc >= len(grid[0]) {
				continue
			}

			flood(nr, nc, grid, seen)
		}
	}
}

var dirs []int = []int{-1, 0, 1}
