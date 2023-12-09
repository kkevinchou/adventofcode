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

		line := record.SingleLine

		if record.ID == 0 {
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

	var currents []string
	for k, _ := range table {
		if string(k[2]) == "A" {
			currents = append(currents, k)
		}
	}
	fmt.Println(len(instructions), "total instructions")
	fmt.Println(currents)

	var finishes []int
	for _, current := range currents {
		var count int
		var instructionIndex int
		for {
			strInstruction := string(instructions[instructionIndex])

			index := 0
			if strInstruction == "R" {
				index = 1
			}

			done := true
			current = table[current][index]
			if string(current[2]) != "Z" {
				done = false
			}

			instructionIndex = (instructionIndex + 1) % len(instructions)

			count += 1
			if done {
				finishes = append(finishes, count)
				break
			}
		}
	}
	fmt.Println("finishes", finishes)

	var result int = 1
	for _, f := range finishes {
		result *= f
	}
	fmt.Println(result)

	// 14257 - 53 * 269
	// 21251 - 79 * 269
	// 11567 - 43 * 269
	// 19099 - 71 * 269
	// 16409 - 61 * 269
	// 12643 - 47 * 269

	fmt.Println(53 * 79 * 43 * 71 * 61 * 47 * 269)
}
