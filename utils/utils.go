package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rec struct {
	LineNumber int
	Line       string
}

func Records(inputFile string) func(func(Rec) bool) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	strData := string(data)
	normalizedStrData := strings.ReplaceAll(strData, "\r\n", "\n")
	splitInput := strings.Split(normalizedStrData, "\n")

	return func(yield func(Rec) bool) {
		for i, line := range splitInput {
			record := Rec{
				LineNumber: i,
				Line:       strings.TrimSpace(line),
			}

			if !yield(record) {
				return
			}
		}
	}
}

func ParseGrid(inputFile string) ([][]string, int, int) {
	var grid [][]string

	for record := range Records(inputFile) {
		grid = append(grid, make([]string, len(record.Line)))
		for c, char := range record.Line {
			r := record.LineNumber
			grid[r][c] = string(char)
		}
	}
	rCount := len(grid)
	cCount := len(grid[0])

	return grid, rCount, cCount
}

func MustParseNum(input string) int {
	out, err := ParseNum(input)
	if err != nil {
		panic(err)
	}
	return out
}

func ParseNum(input string) (int, error) {
	num, err := strconv.ParseInt(input, 10, 64)
	return int(num), err
}

func IsNum(input string) bool {
	_, err := ParseNum(input)
	return err == nil
}

func MapFromSlice[T comparable](elements []T) map[T]T {
	result := map[T]T{}
	for _, element := range elements {
		result[element] = element
	}
	return result
}

func IntAbs(value int) int {
	if value < 0 {
		return -value
	}
	return value

}

func StringSliceToIntSlice(elements []string) []int {
	var result []int
	for _, element := range elements {
		if len(strings.TrimSpace(element)) == 0 {
			continue
		}

		result = append(result, MustParseNum(element))
	}
	return result
}

func StringToStringSlice(s string) []string {
	var result []string

	for _, char := range s {
		result = append(result, string(char))
	}
	return result
}

func PrintGrid(grid [][]string) {
	for r := range len(grid) {
		for c := range len(grid[0]) {
			fmt.Printf(grid[r][c] + " ")
		}
		fmt.Printf("\n")
	}
}
