package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

func main() {
	var result int
	for record := range utils.Records(file, "\n") {
		strSplit := strings.Split(record.SingleLine, ":")
		target := utils.MustParseNum(strSplit[0])
		nums := utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(strSplit[1]), " "))
		if solve(nums[0], target, 1, nums) {
			result += target
		}
	}
	fmt.Println(result)
}

func solve(value, target, index int, nums []int) bool {
	if index == len(nums)-1 {
		if value*nums[index] == target {
			return true
		}

		if value+nums[index] == target {
			return true
		}

		if join(value, nums[index]) == target {
			return true
		}

		return false
	}

	return solve(value*nums[index], target, index+1, nums) || solve(value+nums[index], target, index+1, nums) || solve(join(value, nums[index]), target, index+1, nums)
}

func join(a, b int) int {
	strNum := fmt.Sprintf("%d%d", a, b)
	return utils.MustParseNum(strNum)
}
