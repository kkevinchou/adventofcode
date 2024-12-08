package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	generator := utils.RecordGenerator("input", "\n")

	total := 0
	for {
		record, done := generator()
		if done {
			break
		}
		splits := strings.Split(record.Line, ":")

		maxSoFar := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		gameSplits := strings.Split(splits[1], ";")
		for _, gameSplit := range gameSplits {
			colorSplits := strings.Split(gameSplit, ",")

			for _, colorSplit := range colorSplits {
				colorSplit = strings.TrimSpace(colorSplit)
				amountSplit := strings.Split(colorSplit, " ")
				amount := utils.MustParseNum(amountSplit[0])
				color := amountSplit[1]

				if amount > maxSoFar[color] {
					maxSoFar[color] = amount
				}
			}
		}

		power := maxSoFar["red"] * maxSoFar["green"] * maxSoFar["blue"]
		total += power

	}
	fmt.Println(total)
}
