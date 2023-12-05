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

	xSlope := []int{1, 1, 3, 5, 7}
	ySlope := []int{1, 2, 1, 1, 1}

	treeCounts := make([]int, len(xSlope))
	xPosition := make([]int, len(xSlope))

	numSlopes := len(xSlope)
	width := strings.Index(strData, "\n")

	splitString := strings.Split(strData, "\n")
	if len(splitString) <= 1 {
		return
	}

	// start from the second line so y = 1
	for y := 1; y < len(splitString); y++ {
		line := splitString[y]
		for i := 0; i < numSlopes; i++ {
			if y%ySlope[i] != 0 {
				continue
			}

			xPosition[i] = (xPosition[i] + xSlope[i]) % width
			if line[xPosition[i]] == '#' {
				treeCounts[i]++
			}
		}
	}

	result := 1
	for _, treeCount := range treeCounts {
		result *= treeCount
	}
	fmt.Println(result)
}
