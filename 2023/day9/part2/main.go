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

		line := record.Line

		sequence := utils.StringSliceToIntSlice(strings.Split(line, " "))
		var lastValSum int

		fmt.Println("---------")
		subSeq := sequence
		fmt.Println(sequence)
		for !zeroes(subSeq) {
			subSeq = derive(subSeq)
			lastValSum += subSeq[len(subSeq)-1]
			fmt.Println(subSeq)
		}
		total += sequence[len(sequence)-1] + lastValSum
		fmt.Println("---------")

	}
	// _ = total
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
