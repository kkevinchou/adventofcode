package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var sequences []string
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine

		sequences = strings.Split(line, ",")
	}

	total := 0
	for _, seq := range sequences {
		total += hash(seq)
	}

	fmt.Println(total)
}

func hash(s string) int {
	total := 0
	for _, r := range s {

		total += int(r)
		total *= 17
		total %= 256
	}

	return total
}
