package main

import (
	"fmt"
	"os"
	"strings"
)

var file string = "input.txt"
var svrmemo map[string]int

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")
	neighbors := map[string][]string{}
	for _, line := range lines {
		source := line[:3]
		targetString := line[5:]
		targets := strings.Split(targetString, " ")
		neighbors[source] = targets
	}

	memo := map[string]int{}
	a := solve("svr", "fft", neighbors, memo)
	memo = map[string]int{}
	b := solve("fft", "dac", neighbors, memo)
	memo = map[string]int{}
	c := solve("dac", "out", neighbors, memo)
	fmt.Println(a * b * c)
}

func solve(current string, target string, neighbors map[string][]string, memo map[string]int) int {
	if value, ok := memo[fmt.Sprintf("%s_%s", current, target)]; ok {
		return value
	}

	if current == target {
		memo[fmt.Sprintf("%s_%s", current, target)] = 1
		return 1
	}

	var result int
	for _, neighbor := range neighbors[current] {
		result += solve(neighbor, target, neighbors, memo)
	}

	memo[fmt.Sprintf("%s_%s", current, target)] = result
	return result
}
