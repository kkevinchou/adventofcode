package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	generator := utils.RecordGenerator("input", "\n")

	var left []int
	var right []int
	for {
		record, done := generator()
		if done {
			break
		}
		result := strings.Split(record.SingleLine, "   ")

		left = append(left, utils.MustParseNum(result[0]))
		right = append(right, utils.MustParseNum(result[1]))
	}

	sort.Ints(left)
	sort.Ints(right)

	var sum int
	for i := range left {

		min := min(left[i], right[i])
		max := max(left[i], right[i])

		sum += max - min
	}

	fmt.Println(sum)
}