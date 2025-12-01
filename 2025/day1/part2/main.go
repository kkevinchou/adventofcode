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

		result += delta / 100
		delta = delta % 100

		if dir == "R" {
			newValue := (value + delta) % 100
			if value != 0 && value > newValue {
				result++
			}
			value = newValue
		} else if dir == "L" {
			newValue := (value - delta + 100) % 100
			if value != 0 && (value < newValue || newValue == 0) {
				result++
			}
			value = newValue
		}
	}

	fmt.Println(result)
}
