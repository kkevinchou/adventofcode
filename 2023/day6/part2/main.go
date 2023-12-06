package main

import (
	"fmt"
	"kkevinchou/adventofcode/utils"
	"strings"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var distances []string
	var durations []string
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine
		lineSplit := strings.Split(line, ":")
		lineType := lineSplit[0]
		if lineType == "Time" {
			durations = strings.Split(strings.TrimSpace(lineSplit[1]), " ")
		} else if lineType == "Distance" {
			distances = strings.Split(strings.TrimSpace(lineSplit[1]), " ")
		}
	}

	goal := utils.MustParseNum(strings.Join(distances, ""))
	maxDuration := utils.MustParseNum(strings.Join(durations, ""))

	var winCount int
	for j := 1; j < maxDuration; j++ {
		distance := simulate(j, maxDuration)
		if distance > goal {
			winCount += 1
		}
	}
	fmt.Println(winCount)

	// start := 0
	// for j := 1; j < maxDuration; j++ {
	// 	distance := simulate(j, maxDuration)
	// 	if distance > goal {
	// 		start = j
	// 		break
	// 	}
	// }
	// end := bs(start, maxDuration, maxDuration, goal)
	// fmt.Println(end - start + 1)
}

func bs(start, end int, maxDuration int, goal int) int {
	for {
		mid := (end-start)/2 + start
		midResult := simulate(mid, maxDuration)

		if midResult > goal && simulate(mid+1, maxDuration) <= goal {
			return mid
		}

		if midResult > goal {
			start = mid
		} else {
			end = mid
		}
	}
}

func simulate(pauseTime int, maxTime int) int {
	speed := pauseTime
	duration := maxTime - pauseTime

	return speed * duration
}
