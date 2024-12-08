package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
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

		line := record.Line
		lineSplit := strings.Split(line, ":")
		lineType := lineSplit[0]
		if lineType == "Time" {
			durations = utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(lineSplit[1]), " "))
		} else if lineType == "Distance" {
			distances = utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(lineSplit[1]), " "))
		}
	}

	result := 1

	for i, duration := range durations {
		winCount := 0
		for j := 1; j < duration; j++ {
			distance := simulate(j, duration)
			if distance > distances[i] {
				winCount += 1
			}
		}
		result *= winCount
	}

	fmt.Println(result)
}

func simulate(pauseTime int, maxTime int) int {
	speed := pauseTime
	duration := maxTime - pauseTime

	return speed * duration
}
