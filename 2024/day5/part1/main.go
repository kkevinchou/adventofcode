package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	dependencies := map[int][]int{}
	parseDependencies := true

	// parse rules
	for record := range utils.Records("input", "\n") {
		if record.Line == "" {
			break
		}

		if parseDependencies {
			split := strings.Split(record.Line, "|")

			dep, num := utils.MustParseNum(split[0]), utils.MustParseNum(split[1])
			if _, ok := dependencies[num]; !ok {
				dependencies[num] = []int{}
			}
			dependencies[num] = append(dependencies[num], dep)
		}
	}

	var total int

	// evaluation
	for record := range utils.Records("input", "\n") {
		if strings.Contains(record.Line, "|") || record.Line == "" {
			continue
		}

		split := strings.Split(record.Line, ",")
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
			total += nums[len(nums)/2]
		}
	}

	fmt.Println(total)
}
