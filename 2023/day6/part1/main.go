package main

import (
	"fmt"
	"kkevinchou/adventofcode/utils"
	"strings"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var durations []int
	var distances []int
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine
		lineSplit := strings.Split(line, ":")
		lineType := lineSplit[0]
		if lineType == "Time" {
			durations = utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(lineSplit[1]), " "))
		} else if lineType == "Distance" {
			distances = utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(lineSplit[1]), " "))
		}
	}

	fmt.Println(durations, distances)

	result := 1

	for i, duration := range durations {
		winCount := 0
		for j := 1; j < distances[i]; j++ {
			distance := simulate(j, duration)
			if distance > distances[i] {
				winCount += 1
			}
		}
		fmt.Println(winCount)
		result *= winCount
	}

	fmt.Println(result)
}

func simulate(pauseTime int, maxTime int) int {
	speed := pauseTime
	duration := maxTime - pauseTime

	return speed * duration
}
