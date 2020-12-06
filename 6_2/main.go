package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	answers := map[int]map[string]int{}
	groupSizes := map[int]int{}

	recordGenerator := constructRecordGenerator("input")
	for {
		recordNumber, recordSize, record, done := recordGenerator()
		if done {
			break
		}

		answers[recordNumber] = map[string]int{}
		for _, line := range strings.Split(record, "\n") {
			for _, c := range line {
				answers[recordNumber][string(c)]++
			}
		}

		groupSizes[recordNumber] = recordSize
	}

	total := 0
	for recordNumber, v := range answers {
		for _, answerCount := range v {
			if answerCount == groupSizes[recordNumber] {
				total++
			}
		}
	}

	fmt.Println(total)
}

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
		recordNumber++

		if currentLineNumber >= len(splitInput) {
			return -1, -1, "", true
		}

		records := []string{}
		for i := currentLineNumber; i < len(splitInput); i++ {
			currentLineNumber++
			if splitInput[i] == "" {
				return recordNumber, len(records), strings.Join(records, "\n"), false
			}
			records = append(records, splitInput[i])
		}
		return recordNumber, len(records), strings.Join(records, "\n"), false
	}
}
