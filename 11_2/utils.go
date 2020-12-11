package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Record struct {
	ID    int
	Lines []string
}

func constructRecordGenerator(inputFile, separator string) func() (Record, bool) {
	data, err := ioutil.ReadFile("input")
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

		if separator == "\n\n" {
			recordLines = strings.Split(splitInput[recordNumber], "\n")
		} else if separator == "\n" {
			recordLines = []string{splitInput[recordNumber]}
		}

		record := Record{
			ID:    recordNumber,
			Lines: recordLines,
		}

		recordNumber++
		return record, false
	}
}

func mustParseNum(input string) int {
	out, err := parseNum(input)
	if err != nil {
		panic(err)
	}
	return out
}

func parseNum(input string) (int, error) {
	num, err := strconv.ParseInt(input, 10, 64)
	return int(num), err
}

func isNum(input string) bool {
	_, err := parseNum(input)
	return err == nil
}
