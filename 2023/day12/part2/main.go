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
		springs := utils.StringToStringSlice(lineSplit[0] + "?" + lineSplit[0] + "?" + lineSplit[0] + "?" + lineSplit[0] + "?" + lineSplit[0])

		oneConfiguration := utils.StringSliceToIntSlice(strings.Split(lineSplit[1], ","))
		var configurations []int
		for i := 0; i < 5; i++ {
			configurations = append(configurations, oneConfiguration...)
		}

		memo := map[string]int{}
		total += solve(springs, configurations, 0, 0, memo)
	}

	fmt.Println(total)
}

func genKey(springIndex, configIndex int) string {
	return fmt.Sprintf("%d_%d", springIndex, configIndex)
}

func solve(springs []string, configurations []int, springIndex int, configIndex int, memo map[string]int) int {
	key := genKey(springIndex, configIndex)
	if _, ok := memo[key]; ok {
		return memo[key]
	}

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
		takeItCount = solve(springs, configurations, springIndex+springCount+1, configIndex+1, memo)
	}

	if springs[springIndex] == "?" || springs[springIndex] == "." {
		leaveItCount = solve(springs, configurations, springIndex+1, configIndex, memo)
	}

	// fmt.Println(springIndex, "take it count", takeItCount)
	// fmt.Println(springIndex, "leave it count", leaveItCount)

	result := takeItCount + leaveItCount

	memo[key] = result

	return result

}
