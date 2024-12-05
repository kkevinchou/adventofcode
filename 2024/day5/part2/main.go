package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

func main() {
	dependencies := map[int][]int{}
	parseDependencies := true

	// parse rules
	for record := range utils.Records(file, "\n") {
		if record.SingleLine == "" {
			break
		}

		if parseDependencies {
			split := strings.Split(record.SingleLine, "|")

			dep, num := utils.MustParseNum(split[0]), utils.MustParseNum(split[1])
			if _, ok := dependencies[num]; !ok {
				dependencies[num] = []int{}
			}
			dependencies[num] = append(dependencies[num], dep)
		}
	}

	var total int

	// evaluation
	for record := range utils.Records(file, "\n") {
		if strings.Contains(record.SingleLine, "|") || record.SingleLine == "" {
			continue
		}

		split := strings.Split(record.SingleLine, ",")
		nums := utils.StringSliceToIntSlice(split)

		seen := map[int]bool{}
		exists := map[int]bool{}

		for _, num := range nums {
			exists[num] = true
		}

		ordered := true
		for _, num := range nums {
			for _, dep := range dependencies[num] {
				if _, ok := seen[dep]; !ok && exists[dep] {
					ordered = false
					break
				}
			}

			if !ordered {
				break
			}

			seen[num] = true
		}

		if ordered {
			continue
		}

		var ordering []int
		seen = map[int]bool{}
		processed := map[int]bool{}

		for len(ordering) < len(nums) {
			for i, num := range nums {
				if processed[i] {
					continue
				}
				okay := true
				for _, dep := range dependencies[num] {
					if _, ok := seen[dep]; !ok && exists[dep] {
						okay = false
						break
					}
				}
				if okay {
					seen[num] = true
					ordering = append(ordering, num)
					processed[i] = true
				}
			}
		}

		total += ordering[len(ordering)/2]
	}

	fmt.Println(total)
}
