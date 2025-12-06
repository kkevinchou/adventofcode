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

	nums := [][]int{}
	ops := []string{}

	lines := strings.Split(string(data), "\r\n")
	for i, line := range lines {
		nums = append(nums, []int{})
		split := strings.Split(line, " ")
		for _, s := range split {
			if s == "" {
				continue
			}

			if s == "*" || s == "+" {
				ops = append(ops, s)
			} else {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}

				nums[i] = append(nums[i], n)
			}
		}
	}

	// var result []int
	// result := []int{}
	result := make([]int, len(ops))
	for i, op := range ops {
		if op == "*" {
			result[i] = 1
		} else {
			result[i] = 0
		}
	}

	for i := range len(nums) {
		for j := range len(nums[i]) {
			if ops[j] == "*" {
				result[j] *= nums[i][j]
			} else {
				result[j] += nums[i][j]
			}
		}
	}

	var total int
	for _, n := range result {
		total += n
	}

	// fmt.Println(nums)
	// fmt.Println(ops)
	// fmt.Println(result)
	fmt.Println(total)
}
