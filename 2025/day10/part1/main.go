package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input.txt"

func main() {
	// g := "...#."
	// gLen := len(g)
	// a := 0
	// fmt.Println(bitsToPattern(a))
	// a ^= setBits([]int{0, 4}, gLen)
	// fmt.Println(bitsToPattern(a))
	// a ^= setBits([]int{0, 1, 2}, gLen)
	// fmt.Println(bitsToPattern(a))
	// a ^= setBits([]int{1, 2, 3, 4}, gLen)
	// fmt.Println(bitsToPattern(a))

	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	squareBrackets := regexp.MustCompile(`\[(.*?)\]`)
	parenthesis := regexp.MustCompile(`\((.*?)\)`)
	// braces := regexp.MustCompile(`\[(.*?)\]|\((.*?)\)|\{(.*?)\}`)
	lines := strings.Split(string(data), "\r\n")

	var result int
	for _, line := range lines {
		match := squareBrackets.FindString(line)
		match = strings.Trim(match, "[]")
		goalLength := len(match)

		goal := patternToBits(match)

		ops := []int{}
		matches := parenthesis.FindAllString(line, -1)
		for _, match := range matches {
			match := strings.Trim(match, "()")
			bitList := utils.StringSliceToIntSlice(strings.Split(match, ","))
			op := setBits(bitList, goalLength)
			ops = append(ops, op)
		}

		r := solve(0, goal, ops)
		// fmt.Println(r)
		result += r
	}
	fmt.Println(result)
}

func solve(current int, goal int, ops []int) int {
	stack := []int{current}

	if current == goal {
		return 0
	}

	var count int
	for {
		count++
		nextStack := []int{}

		for _, num := range stack {
			for _, op := range ops {
				next := num ^ op
				// strNum := bitsToPattern(num)
				// strOp := bitsToPattern(op)
				// strNext := bitsToPattern(next)
				// fmt.Println(strNum, "+", strOp, "=", strNext)

				if next == goal {
					return count
				}
				nextStack = append(nextStack, next)
			}
		}

		stack = nextStack
	}
}

func setBits(indices []int, maxLength int) int {
	n := 0
	for _, idx := range indices {
		n |= 1 << (maxLength - idx - 1)
	}
	return n
}

func patternToBits(s string) int {
	n := 0
	// for i := len(s) - 1; i >= 0; i-- {
	for _, c := range s {
		// c := s[i]
		n <<= 1 // shift left
		if c == '#' {
			n |= 1
		}
	}
	return n
}

func bitsToPattern(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
