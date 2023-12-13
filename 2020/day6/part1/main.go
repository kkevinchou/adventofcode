package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	strData := string(data)

	answers := map[int]map[string]bool{}

	groupNumber := 0
	for _, line := range strings.Split(strData, "\n") {
		if line == "" {
			groupNumber++
		}

		if _, ok := answers[groupNumber]; !ok {
			answers[groupNumber] = map[string]bool{}
		}

		for _, c := range line {
			answers[groupNumber][string(c)] = true
		}
	}

	total := 0
	for k := range answers {
		total += len(answers[k])
	}

	fmt.Println(total)
}
