package utils

import (
	"os"
	"strings"
)

type Record struct {
	LineNumber int
	Lines      []string
	Line       string
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
			LineNumber: recordNumber,
			Lines:      recordLines,
			Line:       strings.TrimSpace(splitInput[recordNumber]),
		}

		recordNumber++
		return record, false
	}
}
