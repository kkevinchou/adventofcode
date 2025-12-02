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
	var total int
	rangesSplits := strings.Split(string(data), ",")
	for _, rangeSplit := range rangesSplits {
		split := strings.Split(rangeSplit, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])

		for i := start; i <= end; i++ {
			if isInvalid(i) {
				total += i
				fmt.Println(i)
			}
		}
	}
	fmt.Println(total)
}

func isInvalid(num int) bool {
	numAsStr := fmt.Sprintf("%d", num)
	if len(numAsStr)%2 != 0 {
		return false
	}

	firstHalf := numAsStr[:len(numAsStr)/2]
	secondHalf := numAsStr[len(numAsStr)/2:]

	return firstHalf == secondHalf
}
