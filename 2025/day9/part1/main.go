package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

var file string = "input.txt"

type Point struct {
	r, c int
}

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")
	var points []Point
	for _, line := range lines {
		var r, c int
		fmt.Sscanf(line, "%d,%d", &c, &r)
		points = append(points, Point{r: r, c: c})
	}

	area := math.MinInt
	for i := 0; i < len(points); i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			area = max(area, (abs(a.r-b.r)+1)*(abs(a.c-b.c)+1))
		}
	}
	fmt.Println(area)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
