package main

import (
	"fmt"
	"strings"
)

func main() {
	answers := map[int]map[string]int{}
	groupSizes := map[int]int{}

	recordGenerator := constructRecordGenerator("input")
	for {
		recordNumber, recordSize, record, done := recordGenerator()
		if done {
			break
		}

		answers[recordNumber] = map[string]int{}
		for _, line := range strings.Split(record, "\n") {
			for _, c := range line {
				answers[recordNumber][string(c)]++
			}
		}

		groupSizes[recordNumber] = recordSize
	}

	total := 0
	for recordNumber, v := range answers {
		for _, answerCount := range v {
			if answerCount == groupSizes[recordNumber] {
				total++
			}
		}
	}

	fmt.Println(total)
}
