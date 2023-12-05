package main

import (
	"fmt"
	"kkevinchou/adventofcode/utils"
	"strings"
)

func main() {
	generator := utils.RecordGenerator("input", "\n")

	max := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	total := 0
	for {
		record, done := generator()
		if done {
			break
		}
		splits := strings.Split(record.SingleLine, ":")
		gameNumber := utils.MustParseNum(splits[0][5:])

		gameOK := true
		gameSplits := strings.Split(splits[1], ";")
		for _, gameSplit := range gameSplits {
			colorSplits := strings.Split(gameSplit, ",")

			for _, colorSplit := range colorSplits {
				colorSplit = strings.TrimSpace(colorSplit)
				amountSplit := strings.Split(colorSplit, " ")
				amount := utils.MustParseNum(amountSplit[0])
				color := amountSplit[1]

				if max[color] < amount {
					gameOK = false
				}
			}
		}

		if gameOK {
			total += gameNumber
		}
	}
	fmt.Println(total)
}
