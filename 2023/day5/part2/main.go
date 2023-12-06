package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n\r\n")
	var ranges [][2]int
	var mappings []Mapping

	for {
		record, done := generator()
		if done {
			break
		}

		lines := record.Lines

		typeSplit := strings.Split(lines[0], ":")
		if typeSplit[0] == "seeds" {
			nums := strings.Split(strings.TrimSpace(typeSplit[1]), " ")
			for i := 0; i < len(nums); i += 2 {
				ranges = append(ranges, [2]int{utils.MustParseNum(nums[i]), utils.MustParseNum(nums[i+1])})
			}
		} else {
			mapping := Mapping{}
			for i, line := range lines {
				if i == 0 {
					continue
				}

				mappingSplit := strings.Split(line, " ")
				currentRange := Range{
					DestinationStart: utils.MustParseNum(mappingSplit[0]),
					SourceStart:      utils.MustParseNum(mappingSplit[1]),
					Count:            utils.MustParseNum(mappingSplit[2]),
				}

				mapping.Ranges = append(mapping.Ranges, currentRange)
			}
			sort.Slice(mapping.Ranges, func(i, j int) bool {
				return mapping.Ranges[i].SourceStart < mapping.Ranges[j].SourceStart
			})

			var fillerRanges []Range
			for i := range mapping.Ranges {
				if i == len(mapping.Ranges)-1 {
					break
				}
				currentRange := mapping.Ranges[i]
				nextRange := mapping.Ranges[i+1]
				count := nextRange.SourceStart - (currentRange.SourceStart + currentRange.Count)
				if count > 0 {
					fillerRange := Range{
						SourceStart:      currentRange.SourceStart + currentRange.Count,
						DestinationStart: currentRange.SourceStart + currentRange.Count,
						Count:            count,
					}
					fillerRanges = append(fillerRanges, fillerRange)
				}
			}
			mapping.Ranges = append(mapping.Ranges, fillerRanges...)

			sort.Slice(mapping.Ranges, func(i, j int) bool {
				return mapping.Ranges[i].SourceStart < mapping.Ranges[j].SourceStart
			})

			mappings = append(mappings, mapping)
		}
	}

	// fmt.Println(ranges)
	// for _, mapping := range mappings {
	// 	fmt.Println(mapping)
	// }

	// fmt.Println("--------------")

	for i := 0; i < len(mappings); i++ {
		ranges = lookup(mappings[i], ranges)
		// fmt.Println(ranges)
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	fmt.Println(ranges[0][0])
}

type Range struct {
	DestinationStart int
	SourceStart      int
	Count            int
}

type Mapping struct {
	Ranges []Range
}

// TODO - handle case where there's a lookup miss and the ID should remain the same
func lookup(mapping Mapping, ranges [][2]int) [][2]int {
	var newRanges [][2]int
	for _, lookupRange := range ranges {
		start := lookupRange[0]
		count := lookupRange[1]

		// start of ranges
		if start < mapping.Ranges[0].SourceStart {
			fillerCount := count
			if start+count > mapping.Ranges[0].SourceStart {
				fillerCount = mapping.Ranges[0].SourceStart - start
			}
			newRanges = append(newRanges, [2]int{start, fillerCount})
		}

		// end of ranges
		lastRange := mapping.Ranges[len(mapping.Ranges)-1]
		endRangeIndex := lastRange.SourceStart + lastRange.Count - 1
		if start+count-1 > endRangeIndex {
			fillerStart := endRangeIndex + 1
			if start > fillerStart {
				fillerStart = start
			}

			fillerCount := start + count - fillerStart

			newRanges = append(newRanges, [2]int{fillerStart, fillerCount})
		}

		// middle of ranges
		for _, r := range mapping.Ranges {
			if start >= r.SourceStart+r.Count {
				continue
			} else if start+count-1 < r.SourceStart {
				continue
			}

			biggestMin := start
			if r.SourceStart > biggestMin {
				biggestMin = r.SourceStart
			}

			smallestMax := start + count
			if r.SourceStart+r.Count < smallestMax {
				smallestMax = r.SourceStart + r.Count
			}

			newStart := r.DestinationStart + (biggestMin - r.SourceStart)
			newCount := smallestMax - biggestMin

			newRanges = append(newRanges, [2]int{newStart, newCount})
		}
	}
	return newRanges
}
