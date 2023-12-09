package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var total int
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine

		sequence := utils.StringSliceToIntSlice(strings.Split(line, " "))
		var firstVals []int

		firstVals = append(firstVals, sequence[0])

		subSeq := sequence
		for !zeroes(subSeq) {
			subSeq = derive(subSeq)
			firstVals = append(firstVals, subSeq[0])
		}

		var subTotal int
		for i := len(firstVals) - 2; i >= 0; i-- {
			subTotal = firstVals[i] - subTotal
		}
		total += subTotal
	}
	fmt.Println(total)
}

func derive(sequence []int) []int {
	var result []int
	for i := range sequence {
		if i == len(sequence)-1 {
			break
		}

		result = append(result, sequence[i+1]-sequence[i])
	}
	return result
}

func zeroes(sequence []int) bool {
	for _, num := range sequence {
		if num != 0 {
			return false
		}
	}
	return true
}
