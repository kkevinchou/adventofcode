package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	counts := map[int]int{}
	for record := range utils.Records("input", "\n") {
		result := strings.Split(record.Line, "   ")
		counts[utils.MustParseNum(result[1])]++
	}

	var sum int
	for record := range utils.Records("input", "\n") {
		result := strings.Split(record.Line, "   ")
		num := utils.MustParseNum(result[0])
		sum += num * counts[num]
	}

	fmt.Println(sum)
}
