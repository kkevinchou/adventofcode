package main

import (
	"fmt"
	"strings"
)

func main() {
	recordGenerator := constructRecordGenerator("input", "\n")
	p := program{
		instructions: map[int]string{},
	}
	for {
		record, done := recordGenerator()
		if done {
			break
		}

		line := record.Lines[0]
		p.instructions[record.ID] = line
	}
	fmt.Println(p.Run(0, map[int]bool{}, 0, true))
}

type program struct {
	instructionPointer int
	instructions       map[int]string
}

func deepCopy(m map[int]bool) map[int]bool {
	result := map[int]bool{}
	for k, v := range m {
		result[k] = v
	}
	return result
}

func (p *program) Run(instructionPointer int, seen map[int]bool, accumulator int, okayToBranch bool) (int, bool) {
	if instructionPointer >= len(p.instructions) {
		return accumulator, true
	}

	if seen[instructionPointer] {
		return accumulator, false
	}

	newSeen := deepCopy(seen)
	newSeen[instructionPointer] = true

	instruction := p.instructions[instructionPointer]
	split := strings.Split(instruction, " ")

	command := split[0]
	operation := string(split[1][0])
	num := mustParseNum(split[1][1:])

	if command == "nop" {
		if operation == "+" {
		} else if operation == "-" {
			num *= -1
		} else {
			panic("weird op " + operation)
		}

		if okayToBranch {
			// normal run
			if result, terminated := p.Run(instructionPointer+1, newSeen, accumulator, true); terminated {
				return result, terminated
			}

			// branched run
			if result, terminated := p.Run(instructionPointer+num, newSeen, accumulator, false); terminated {
				return result, terminated
			}
		} else {
			// normal run
			if result, terminated := p.Run(instructionPointer+1, newSeen, accumulator, okayToBranch); terminated {
				return result, terminated
			}
		}
	}

	if command == "acc" {
		if operation == "+" {
			accumulator += num
		} else if operation == "-" {
			accumulator -= num
		} else {
			panic("weird op " + operation)
		}
		return p.Run(instructionPointer+1, newSeen, accumulator, okayToBranch)
	}

	if command == "jmp" {
		if operation == "+" {
		} else if operation == "-" {
			num *= -1
		} else {
			panic("weird op " + operation)
		}

		if okayToBranch {
			// normal run
			if result, terminated := p.Run(instructionPointer+num, newSeen, accumulator, true); terminated {
				return result, terminated
			}

			// branched run
			if result, terminated := p.Run(instructionPointer+1, newSeen, accumulator, false); terminated {
				return result, terminated
			}
		} else {
			// normal run
			if result, terminated := p.Run(instructionPointer+num, newSeen, accumulator, okayToBranch); terminated {
				return result, terminated
			}
		}
	}

	return -1, false
}

// keep track of lines that have been run (record number)
