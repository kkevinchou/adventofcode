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
	input := strings.Split(strData, "\n")

	current := map[string]string{}

	count := 0
	for i, line := range input {
		if line != "" {
			parseLine(current, line)
		}

		if line == "" || i == len(input)-1 {
			if eval(current) {
				count++
			}
			current = map[string]string{}
		}
	}

	fmt.Println(count)
}

func eval(data map[string]string) bool {
	// byr (Birth Year)
	// iyr (Issue Year)
	// eyr (Expiration Year)
	// hgt (Height)
	// hcl (Hair Color)
	// ecl (Eye Color)
	// pid (Passport ID)
	// cid (Country ID)

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range fields {
		if _, ok := data[field]; !ok {
			return false
		}
	}

	return true
}
