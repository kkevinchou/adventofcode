package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	generator := utils.RecordGenerator("input", "\n")

	sum := 0
	for {
		record, done := generator()
		if done {
			break
		}

		calibration := ""
		line := record.Line
		for i := 0; i < len(line); i++ {
			if utils.IsNum(string(line[i])) {
				calibration += string(line[i])
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if utils.IsNum(string(line[i])) {
				calibration += string(line[i])
				break
			}
		}
		sum += utils.MustParseNum(calibration)
	}
	fmt.Println(sum)
}
