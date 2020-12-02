package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	input := string(dat)

	splitStr := strings.Split(input, "\n")
	numList := []int64{}
	numbersMap := map[int64]map[int64]struct{}{}

	for _, line := range splitStr {
		num, err := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		if err != nil {
			panic(err)
		}

		numList = append(numList, num)
		numbersMap[num] = map[int64]struct{}{}
	}

	for _, num1 := range numList {
		for _, num2 := range numList {
			if _, ok := numbersMap[2020-(num1+num2)]; ok {
				fmt.Println("FOUND: ", 2020-(num1+num2), num1, num2)
				fmt.Println((2020 - (num1 + num2)) * num1 * num2)
			}
		}
	}
}
