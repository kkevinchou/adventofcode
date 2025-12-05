package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var file string = "input.txt"

type Range struct {
	Min int
	Max int
}

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var ranges []Range

	lines := strings.Split(string(data), "\r\n")

	mode := 0
	var count int
	for _, line := range lines {
		if line == "" {
			mode = 1
			continue
		}

		if mode == 0 {
			split := strings.Split(line, "-")
			min, _ := strconv.Atoi(split[0])
			max, _ := strconv.Atoi(split[1])
			ranges = append(ranges, Range{Min: min, Max: max})
		}

		if mode == 1 {
			num, _ := strconv.Atoi(line)
			for _, r := range ranges {
				if num >= r.Min && num <= r.Max {
					count++
					break
				}
			}
		}
	}
	fmt.Println(count)
}
