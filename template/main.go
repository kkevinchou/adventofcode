package main

import (
	"fmt"
	"strings"
)

type BagRule struct {
	color string
	count int
}

func main() {
	forest := map[string][]BagRule{}
	recordGenerator := constructRecordGenerator("input", "\n")
	for {
		record, done := recordGenerator()
		if done {
			break
		}
		parseRow(forest, record.Lines[0])
	}
	fmt.Println(countColor(forest, "shiny gold") - 1)
}

func countColor(forest map[string][]BagRule, color string) int {
	sum := 1
	for _, child := range forest[color] {
		sum += child.count * countColor(forest, child.color)
	}

	return sum
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
