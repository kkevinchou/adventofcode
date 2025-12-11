package main

import (
	"fmt"
	"os"
	"strings"
)

var file string = "input.txt"
var results [][]string

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

	fmt.Println(solve("you", neighbors))
}

func solve(current string, neighbors map[string][]string) int {
	if current == "out" {
		return 1
	}

	var result int
	for _, neighbor := range neighbors[current] {
		result += solve(neighbor, neighbors)
	}

	return result
}
