package main

import (
	"fmt"
	"kkevinchou/adventofcode/utils"
	"strings"
)

func main() {
	file := "sample"
	generator := utils.RecordGenerator(file, "\r\n\r\n")
	var seeds []int
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
			for _, num := range nums {
				seeds = append(seeds, utils.MustParseNum(num))
			}
		} else {
			mapping := Mapping{}
			for i, line := range lines {
				if i == 0 {
					continue
				}

				mappingSplit := strings.Split(line, " ")
				mapping.Ranges = append(mapping.Ranges, Range{
					DestinationStart: utils.MustParseNum(mappingSplit[0]),
					SourceStart:      utils.MustParseNum(mappingSplit[1]),
					Count:            utils.MustParseNum(mappingSplit[2]),
				})
			}

			mappings = append(mappings, mapping)
		}
	}

	var location int
	for i, seed := range seeds {
		id := seed
		for j := 0; j < len(mappings); j++ {
			id = lookup(mappings[j], id)
		}
		if id < location || i == 0 {
			location = id
		}
	}
	fmt.Println(location)
}

type Range struct {
	SourceStart      int
	DestinationStart int
	Count            int
}

type Mapping struct {
	Ranges []Range
}

func lookup(mapping Mapping, id int) int {
	result := id
	for _, r := range mapping.Ranges {
		if id >= r.SourceStart && id <= r.SourceStart+r.Count-1 {
			result = r.DestinationStart + (id - r.SourceStart)
			break
		}
	}
	return result
}
