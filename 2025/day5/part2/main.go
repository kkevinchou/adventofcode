package main

import (
	"fmt"
	"os"
	"sort"
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

	for _, line := range lines {
		if line == "" {
			break
		}

		split := strings.Split(line, "-")
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])
		ranges = append(ranges, Range{Min: min, Max: max})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	windowMin := ranges[0].Min
	windowMax := ranges[0].Max

	var count int
	for _, r := range ranges {
		if r.Min > windowMax {
			count += windowMax - windowMin + 1
			windowMin = r.Min
			windowMax = r.Max
		} else if r.Min <= windowMax {
			windowMax = max(windowMax, r.Max)
		}
	}

	count += windowMax - windowMin + 1
	fmt.Println(count)
}
