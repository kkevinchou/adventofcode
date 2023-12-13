package main

import (
	"fmt"
	"strings"
)

func main() {
	recordGenerator := constructRecordGenerator("input", "\n")
	p := program{
		seen:         map[int]bool{},
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
	fmt.Println(p.Run(0))
}

type program struct {
	instructionPointer int
	accumulator        int
	seen               map[int]bool
	instructions       map[int]string
}

func (p *program) Run(instructionPointer int) int {
	if instructionPointer >= len(p.instructions) {
		return p.accumulator
	}

	if p.seen[instructionPointer] {
		return p.accumulator
	}
	p.seen[instructionPointer] = true

	instruction := p.instructions[instructionPointer]
	split := strings.Split(instruction, " ")

	command := split[0]
	operation := string(split[1][0])
	num := mustParseNum(split[1][1:])

	if command == "nop" {
		return p.Run(instructionPointer + 1)
	}

	if command == "acc" {
		if operation == "+" {
			p.accumulator += num
		} else if operation == "-" {
			p.accumulator -= num
		} else {
			panic("weird op " + operation)
		}
		return p.Run(instructionPointer + 1)
	}

	if command == "jmp" {
		if operation == "+" {
			return p.Run(instructionPointer + num)
		} else if operation == "-" {
			return p.Run(instructionPointer - num)
		} else {
			panic("weird op " + operation)
		}
	}

	panic("wtf")
}

// keep track of lines that have been run (record number)
