package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var total int
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine

		lineSplit := strings.Split(line, " ")
		springs := utils.StringToStringSlice(lineSplit[0])
		configurations := utils.StringSliceToIntSlice(strings.Split(lineSplit[1], ","))
		_, _ = springs, configurations

		total += solve(springs, configurations, 0, 0)
	}

	fmt.Println(total)
}

func solve(springs []string, configurations []int, springIndex int, configIndex int) int {
	if configIndex == len(configurations) {
		for i := springIndex; i < len(springs); i++ {
			if springs[i] == "#" {
				return 0
			}
		}
		return 1
	}

	springCount := configurations[configIndex]

	if springIndex == len(springs) {
		return 0
	} else if springIndex+springCount > len(springs) {
		return 0
	}

	// validate we can take it
	canTake := true
	for i := 0; i < springCount; i++ {
		if springs[springIndex+i] == "." {
			canTake = false
			break
		}
	}

	// make sure we can place . at the end, OR we're at the end of the input
	if springIndex+springCount != len(springs) && (springs[springIndex+springCount] == "#") {
		canTake = false
	}

	var takeItCount int
	var leaveItCount int

	if canTake {
		takeItCount = solve(springs, configurations, springIndex+springCount+1, configIndex+1)
	}

	if springs[springIndex] == "?" || springs[springIndex] == "." {
		leaveItCount = solve(springs, configurations, springIndex+1, configIndex)
	}

	result := takeItCount + leaveItCount
	return result
}
