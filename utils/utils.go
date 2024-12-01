package utils

import (
	"os"
	"strconv"
	"strings"
)

type Record struct {
	ID         int
	Lines      []string
	SingleLine string
}

func Records(inputFile, separator string) func(func(Record) bool) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	strData := string(data)
	normalizedStrData := strings.ReplaceAll(strData, "\r\n", "\n")
	splitInput := strings.Split(normalizedStrData, separator)

	return func(yield func(Record) bool) {
		for i, line := range splitInput {
			var recordLines []string
			if separator == "\n\n" {
				for _, line := range strings.Split(line, "\n") {
					recordLines = append(recordLines, strings.TrimSpace(line))
				}
			} else if separator == "\n" {
				recordLines = []string{strings.TrimSpace(line)}
			}

			record := Record{
				ID:         i,
				Lines:      recordLines,
				SingleLine: strings.TrimSpace(line),
			}

			if !yield(record) {
				return
			}
		}
	}
}

func RecordGenerator(inputFile, separator string) func() (Record, bool) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	strData := string(data)
	splitInput := strings.Split(strData, separator)

	recordNumber := 0
	return func() (Record, bool) {
		if recordNumber >= len(splitInput) {
			return Record{}, true
		}
		var recordLines []string

		if separator == "\r\n\r\n" {
			recordLines = strings.Split(splitInput[recordNumber], "\r\n")
		} else if separator == "\n" {
			recordLines = []string{strings.TrimSpace(splitInput[recordNumber])}
		}

		record := Record{
			ID:         recordNumber,
			Lines:      recordLines,
			SingleLine: strings.TrimSpace(splitInput[recordNumber]),
		}

		recordNumber++
		return record, false
	}
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
