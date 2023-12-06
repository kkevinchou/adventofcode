package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	total := 0
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine

		lineSplit := strings.Split(line, ":")
		numbers := strings.Split(lineSplit[1], "|")

		winningNumbers := utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(numbers[0]), " "))
		myNumbers := utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(numbers[1]), " "))

		set := utils.MapFromSlice[int](winningNumbers)

		points := 0
		for _, num := range myNumbers {
			if _, ok := set[num]; !ok {
				continue
			}

			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}

		if points > 0 {
			total += points
		}
	}

	fmt.Println(total)
}
