package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	generator := utils.RecordGenerator("input", "\n")

	counts := map[int]int{}
	for {
		record, done := generator()
		if done {
			break
		}
		result := strings.Split(record.SingleLine, "   ")

		counts[utils.MustParseNum(result[1])]++
	}

	generator = utils.RecordGenerator("input", "\n")

	var sum int
	for record, done := generator(); !done; record, done = generator() {

		if done {
			break
		}
		result := strings.Split(record.SingleLine, "   ")

		num := utils.MustParseNum(result[0])

		sum += num * counts[num]
	}

	fmt.Println(sum)
}
