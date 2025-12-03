package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var file string = "input.txt"
var memo map[string]string

func main() {

	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var sum int

	lines := strings.Split(string(data), "\r\n")
	for _, line := range lines {
		memo = map[string]string{}
		val, _ := strconv.Atoi(biggest(line, 12))
		sum += val
	}

	fmt.Println(sum)
}

func biggest(nums string, numsRemaining int) string {
	if cachedValue, ok := memo[fmt.Sprintf("%s_%d", nums, numsRemaining)]; ok {
		return cachedValue
	}

	if numsRemaining == 0 {
		return ""
	}

	if len(nums) == numsRemaining {
		return nums
	}

	// take it
	left := string(nums[0]) + biggest(nums[1:], numsRemaining-1)

	// leave it
	right := biggest(nums[1:], numsRemaining)

	leftNum, _ := strconv.Atoi(left)
	rightNum, _ := strconv.Atoi(right)

	if leftNum > rightNum {
		memo[fmt.Sprintf("%s_%d", nums, numsRemaining)] = left
		return left
	} else {
		memo[fmt.Sprintf("%s_%d", nums, numsRemaining)] = right
		return right
	}
}
