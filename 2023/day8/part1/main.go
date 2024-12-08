package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	table := map[string][2]string{}
	var instructions string

	for {
		record, done := generator()
		if done {
			break
		}

		line := record.Line

		if record.LineNumber == 0 {
			instructions = line
			continue
		} else if line == "" {
			continue
		}

		lineSplit := strings.Split(line, "=")
		left := lineSplit[1][2:5]
		right := lineSplit[1][7:10]
		table[strings.TrimSpace(lineSplit[0])] = [2]string{left, right}
	}

	current := "AAA"
	var count int
	var instructionIndex int
	for {
		strInstruction := string(instructions[instructionIndex])

		index := 0
		if strInstruction == "R" {
			index = 1
		}

		current = table[current][index]
		count += 1
		instructionIndex = (instructionIndex + 1) % len(instructions)
		if current == "ZZZ" {
			break
		}
	}
	fmt.Println(count)
}
