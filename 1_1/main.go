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

	numbers := map[int64]struct{}{}
	for _, line := range strings.Split(input, "\n") {
		num, err := strconv.ParseInt(strings.TrimSpace(line), 10, 64)
		if err != nil {
			panic(err)
		}

		numbers[num] = struct{}{}

		if _, ok := numbers[2020-num]; ok {
			fmt.Println("FOUND: ", num)
			fmt.Println((2020 - num) * num)
		}
	}
}
