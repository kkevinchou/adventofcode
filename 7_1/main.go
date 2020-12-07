package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type BagRule struct {
	color string
	count int
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	strData := string(data)
	splitInput := strings.Split(strData, "\n")

	forest := map[string][]BagRule{}
	for _, record := range splitInput {
		parseRow(forest, record)
	}

	count := 0
	for color := range forest {
		if findTarget(forest, color, "shiny gold") {
			count++
		}
	}

	fmt.Println(count)
}

func findTarget(forest map[string][]BagRule, color string, target string) bool {
	for _, child := range forest[color] {
		if child.color == target {
			return true
		}
		if findTarget(forest, child.color, target) {
			return true
		}
	}

	return false
}

func parseRow(forest map[string][]BagRule, row string) []BagRule {
	// todo: check if newline matters?
	color, rawRulesList := strings.Split(row, "contain")[0], strings.Split(row, "contain")[1]

	color = stripBag(color)
	rulesList := parseRulesList(rawRulesList)
	forest[color] = rulesList
	return rulesList
}

// 1 shiny coral bag, 4 dotted purple bags.
func parseRulesList(rulesList string) []BagRule {
	if rulesList == "no other bags." {
		return []BagRule{}
	}

	split := strings.Split(rulesList, ",")
	var result []BagRule
	for _, rule := range split {
		clean := stripBag(strings.TrimSpace(rule)) // 1 shiny coral
		ruleSplit := strings.SplitN(clean, " ", 2)
		if ruleSplit[0] == "no" {
			continue
		}
		result = append(result, BagRule{ruleSplit[1], mustParseNum(ruleSplit[0])})
	}
	return result
}

func stripBag(str string) string {
	return str[:strings.Index(str, " bag")]
}
