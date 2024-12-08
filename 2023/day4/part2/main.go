package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	memo := map[int]int{}
	matches := map[int]int{}

	largest := 0
	total := 0
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.Line

		lineSplit := strings.Split(line, ":")
		numbers := strings.Split(lineSplit[1], "|")

		winningNumbers := utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(numbers[0]), " "))
		myNumbers := utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(numbers[1]), " "))

		set := utils.MapFromSlice[int](winningNumbers)

		for _, num := range myNumbers {
			if _, ok := set[num]; !ok {
				continue
			}
			matches[record.LineNumber] += 1
		}
		largest = record.LineNumber
	}

	for i := largest; i >= 0; i-- {
		value := 1
		for j := 0; j < matches[i]; j++ {
			value += memo[i+j+1]
		}
		memo[i] = value
	}

	for _, v := range memo {
		total += v
	}

	fmt.Println(total)
}
