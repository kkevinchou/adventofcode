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

	xPosition := []int{0, 0, 0, 0, 0}
	xSlope := []int{1, 1, 3, 5, 7}
	ySlope := []int{1, 2, 1, 1, 1}
	treeCounts := []int{0, 0, 0, 0, 0}

	numSlopes := len(xSlope)
	width := strings.Index(strData, "\n")

	// start from the second line so y = 1
	y := 1
	splitString := strings.Split(strData, "\n")
	for _, line := range splitString[1:] {
		for i := 0; i < numSlopes; i++ {
			if y%ySlope[i] != 0 {
				continue
			}

			xPosition[i] = (xPosition[i] + xSlope[i]) % width
			if line[xPosition[i]] == '#' {
				treeCounts[i] = treeCounts[i] + 1
			}
		}
		y++
	}

	result := 1
	for _, treeCount := range treeCounts {
		result *= treeCount
	}
	fmt.Println(result)
}
