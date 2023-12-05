package main

import (
	"io/ioutil"
	"strings"
)

func constructRecordGenerator(inputFile string) func() (int, int, string, bool) {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	strData := string(data)
	splitInput := strings.Split(strData, "\n")

	currentLineNumber := 0
	recordNumber := -1
	return func() (int, int, string, bool) {
		if currentLineNumber >= len(splitInput) {
			return -1, -1, "", true
		}

		recordNumber++
		lines := []string{}
		for i := currentLineNumber; i < len(splitInput); i++ {
			currentLineNumber++
			if splitInput[i] == "" {
				return recordNumber, len(lines), strings.Join(lines, "\n"), false
			}
			lines = append(lines, splitInput[i])
		}
		return recordNumber, len(lines), strings.Join(lines, "\n"), false
	}
}
