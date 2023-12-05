package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	strData := string(data)

	treeCount := 0
	width := strings.Index(strData, "\n")

	x := 0
	// start from the second line so y = 1
	splitString := strings.Split(strData, "\n")
	for _, line := range splitString[1:] {
		x = (x + 3) % width
		if line[x] == '#' {
			treeCount++
		}
	}

	fmt.Println(treeCount)
}
