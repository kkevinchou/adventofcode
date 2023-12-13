package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	count := 0
	for _, line := range strings.Split(string(data), "\n") {
		if evalRule(line) {
			count += 1
		}
	}

	fmt.Println(count)
}

func evalRule(line string) bool {
	split := strings.Split(line, " ")
	min := mustParse(strings.Split(split[0], "-")[0])
	max := mustParse(strings.Split(split[0], "-")[1])

	character := string(split[1][0])

	password := split[2]

	count := int64(strings.Count(password, character))

	return count >= min && count <= max
}

func mustParse(input string) int64 {
	out, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		panic(err)
	}
	return out
}
