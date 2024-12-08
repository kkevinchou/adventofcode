package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	var left []int
	var right []int
	for record := range utils.Records("sample") {
		result := strings.Split(record.Line, "   ")
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
