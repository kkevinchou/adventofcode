package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var file string = "input.txt"

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	value := 50
	result := 0

	lines := strings.Split(string(data), "\r\n")
	for _, line := range lines {
		dir := string(line[0])
		delta, err := strconv.Atoi(string(line[1:]))
		if err != nil {
			panic(err)
		}

		if dir == "R" {
			value += delta
		} else if dir == "L" {
			value -= delta
		} else {
			panic("wat")
		}
		value = value % 100

		if value == 0 {
			result++
		}
	}

	fmt.Println(result)
}
