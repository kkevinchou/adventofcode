package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	generator := utils.RecordGenerator("input", "\n")

	sum := 0
	for {
		record, done := generator()
		if done {
			break
		}

		calibration := ""
		line := record.Line

		tokens := tokenize(line)

		calibration += fmt.Sprintf("%d", tokens[0])
		calibration += fmt.Sprintf("%d", tokens[len(tokens)-1])

		sum += utils.MustParseNum(calibration)
	}
	fmt.Println(sum)
}

func tokenize(line string) []int {
	stringSoFar := ""

	var tokens []int

	for i := 0; i < len(line); i++ {
		if utils.IsNum(string(line[i])) {
			tokens = append(tokens, utils.MustParseNum(string(line[i])))
			continue
		}

		stringSoFar += string(line[i])
		if i == len(line)-1 {
			nums := extractNum(stringSoFar)
			tokens = append(tokens, nums...)
			stringSoFar = ""
		} else if utils.IsNum(string(line[i+1])) {
			nums := extractNum(stringSoFar)
			tokens = append(tokens, nums...)
			stringSoFar = ""
		}
	}

	return tokens
}

// takes a string of characters and extracts a digit from its string representation
// aaaaaaaaaaaonebbbbbbbbbbbbb => 1
// can generate multiple numbers
func extractNum(blob string) []int {
	var nums []int

	for i := 0; i < len(blob); i++ {
		digit, ok := matchNum(blob, i)
		if ok {
			nums = append(nums, digit)
		}
	}

	return nums
}

// returns the digit, next index, and whether it was successful
func matchNum(blob string, index int) (int, bool) {
	var numbers = []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for i, num := range numbers {
		if strings.HasPrefix(blob[index:], num) {
			return i + 1, true
		}
	}

	return -1, false
}
