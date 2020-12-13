package main

import (
	"fmt"
	"strings"
)

var directionIndex int
var directions = []string{"E", "S", "W", "N"}

func main() {
	plannedTimeStamp := 0

	recordGenerator := constructRecordGenerator("input", "\n")
	busses := map[int]bool{}
	for {
		record, done := recordGenerator()
		if done {
			break
		}
		if record.ID == 0 {
			plannedTimeStamp = mustParseNum(record.Lines[0])
		}

		if record.ID == 1 {
			split := strings.Split(record.Lines[0], ",")
			for _, id := range split {
				if id == "x" {
					continue
				}
				busses[mustParseNum(id)] = true
			}
		}
	}

	earliestBus := 0
	earliestBusScheduleDiff := 100
	for k := range busses {
		bussTimestamp := k
		for bussTimestamp < plannedTimeStamp {
			bussTimestamp += k
		}
		diff := bussTimestamp - plannedTimeStamp
		if diff < earliestBusScheduleDiff {
			earliestBus = k
			earliestBusScheduleDiff = diff
		}
	}

	fmt.Println(earliestBus * earliestBusScheduleDiff)
}
