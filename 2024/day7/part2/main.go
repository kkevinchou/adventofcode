package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

func main() {
	start := time.Now()
	var result int
	for record := range utils.Records(file) {
		strSplit := strings.Split(record.Line, ":")
		target := utils.MustParseNum(strSplit[0])
		nums := utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(strSplit[1]), " "))

		if solve(nums[0], target, 1, nums) {
			result += target
		}
	}
	fmt.Println(result)
	fmt.Println(time.Since(start))
}

func solve(value, target, index int, nums []int) bool {
	if index == len(nums) {
		return value == target
	} else if value > target {
		return false
	}

	return solve(value*nums[index], target, index+1, nums) ||
		solve(value+nums[index], target, index+1, nums) ||
		solve(join(value, nums[index]), target, index+1, nums)
}

func join(a, b int) int {
	strNum := fmt.Sprintf("%d%d", a, b)
	return utils.MustParseNum(strNum)
}
