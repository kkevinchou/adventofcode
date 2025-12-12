package main

import (
	"fmt"
	"os"
	"strings"
)

var file string = "input.txt"

type Space struct {
	w, h, a, b, c, d, e, f int
}

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")

	var blocks []int
	var spaces []Space
	// for _, line := range lines {
	for i := range len(lines) {
		if lines[i] == "" {
			continue
		}

		if len(lines[i]) >= 2 && lines[i][1] == ':' {
			blockSize := 0
			for j := 1; j <= 3; j++ {
				line := lines[i+j]
				for k := range 3 {
					if line[k] == '#' {
						blockSize++
					}
				}
			}
			blocks = append(blocks, blockSize)
		} else if len(lines[i]) >= 3 && lines[i][2] == 'x' {
			var w int
			var h int

			var a, b, c, d, e, f int

			fmt.Sscanf(lines[i], "%dx%d: %d %d %d %d %d %d", &w, &h, &a, &b, &c, &d, &e, &f)
			s := Space{
				w: w,
				h: h,
				a: a,
				b: b,
				c: c,
				d: d,
				e: e,
				f: f,
			}

			spaces = append(spaces, s)
		}
	}

	bad := 0
	good := 0
	for _, s := range spaces {
		spaceSize := s.w * s.h
		numBlocks := s.a + s.b + s.c + s.d + s.e + s.f

		if spaceSize < s.a*blocks[0]+s.b*blocks[1]+s.c*blocks[2]+s.d*blocks[3]+s.e*blocks[4]+s.f*blocks[5] {
			bad++
		} else if spaceSize >= numBlocks*9 {
			good++
		}
	}
	maybeGood := len(spaces) - bad - good
	fmt.Println(good, bad, maybeGood)
}
