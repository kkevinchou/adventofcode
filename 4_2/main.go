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

	result := true

	result = result && len(data["byr"]) == 4 && mustParseNum(data["byr"]) >= 1920 && mustParseNum(data["byr"]) <= 2002
	result = result && len(data["iyr"]) == 4 && mustParseNum(data["iyr"]) >= 2010 && mustParseNum(data["iyr"]) <= 2020
	result = result && len(data["eyr"]) == 4 && mustParseNum(data["eyr"]) >= 2020 && mustParseNum(data["eyr"]) <= 2030

	result = result && parseHeight(data["hgt"])
	result = result && parseHair(data["hcl"])
	result = result && (data["ecl"] == "amb" || data["ecl"] == "blu" || data["ecl"] == "brn" || data["ecl"] == "gry" || data["ecl"] == "grn" || data["ecl"] == "hzl" || data["ecl"] == "oth")
	result = result && (len(data["pid"]) == 9 && isNum(data["pid"]))

	return result
}

func parseLine(passport map[string]string, line string) {
	fields := strings.Split(line, " ")
	for _, field := range fields {
		key := strings.Split(field, ":")[0]
		val := strings.Split(field, ":")[1]
		passport[key] = val
	}
}

func mustParseNum(input string) int64 {
	out, err := parseNum(input)
	if err != nil {
		panic(err)
	}
	return out
}

func parseNum(input string) (int64, bool) {
	out, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return -1, false
	}
	return out, true
}

func isNum(input string) bool {
	_, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return false
	}
	return true
}

func parseHeight(input string) bool {
	index := strings.Index(input, "cm")
	if index != -1 {
		if num, ok := parseNum(input[:index]); ok {
			return num >= 150 && num <= 193
		}
		return false
	}

	index = strings.Index(input, "in")
	if index != -1 {
		if num, ok := parseNum(input[:index]); ok {
			return num >= 59 && num <= 76
		}
		return false
	}

	return false
}

func parseHair(input string) bool {
	if len(input) != 7 {
		return false
	}

	if input[0] != '#' {
		return false
	}

	for i := 1; i < len(input); i++ {
		c := input[i]
		if !strings.Contains("1234567890abcdef", string(c)) {
			fmt.Println(string(c))
			return false
		}
	}
	return true
}
