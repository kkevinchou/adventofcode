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
			}
		}
	}
	fmt.Println(total)
}

func isInvalid(num int) bool {
	numAsStr := fmt.Sprintf("%d", num)

	for i := 0; i < len(numAsStr)/2; i++ {
		repeats := len(numAsStr) / (i + 1)
		guess := ""
		for range repeats {
			guess += numAsStr[:i+1]
		}
		if guess == numAsStr {
			return true
		}
	}

	return false
}

// // saved for later
// // returns the next invalid id, if the input id is already invalid, it returns itself
// func nextInvalid(input string) (string, int) {
// 	if len(input)%2 == 0 {
// 		half := input[:len(input)/2]
// 		double := half + half
// 		doubleAsInt, _ := strconv.Atoi(double)
// 		return double, doubleAsInt
// 	}

// 	length := len(input)/2 + 1
// 	half := "1"
// 	for range length - 1 {
// 		half += "0"
// 	}
// 	double := half + half
// 	doubleAsInt, _ := strconv.Atoi(double)
// 	return double, doubleAsInt
// }
