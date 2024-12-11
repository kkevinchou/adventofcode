package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

func main() {
	var lineStr []string
	for record := range utils.Records(file) {
		lineStr = strings.Split(record.Line, " ")
	}

	memo := map[string]int{}
	nums := utils.StringSliceToIntSlice(lineStr)

	var result int
	for _, num := range nums {
		r := solve(num, 75, memo)
		result += r
	}
	fmt.Println(result)
}

func solve(value int, blinks int, memo map[string]int) int {
	if blinks == 0 {
		return 1
	}

	if count, ok := memo[key(value, blinks)]; ok {
		return count
	}

	var result int
	if value == 0 {
		result = solve(1, blinks-1, memo)
	} else if l, r, ok := eventDigit(value); ok {
		result = solve(l, blinks-1, memo) + solve(r, blinks-1, memo)
	} else {
		result = solve(value*2024, blinks-1, memo)
	}

	memo[key(value, blinks)] = result
	return result
}

func key(value, blinkCount int) string {
	return fmt.Sprintf("%d_%d", value, blinkCount)
}

func eventDigit(num int) (int, int, bool) {
	s := fmt.Sprintf("%d", num)
	if len(s)%2 == 0 {
		left := s[:len(s)/2]
		right := s[len(s)/2:]
		return utils.MustParseNum(left), utils.MustParseNum(right), true
	}
	return -1, -1, false
}
