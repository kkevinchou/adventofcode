package main

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

func main() {
	start := time.Now()
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
	fmt.Println(time.Since(start))
}

func solve(value, target, index int, nums []int) bool {
	if index == len(nums) {
		return value == target
	} else if value > target {
		return false
	}

	var a, b, c bool

	if index < 3 {
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			a = solve(value*nums[index], target, index+1, nums)
			wg.Done()
		}()

		go func() {
			b = solve(value+nums[index], target, index+1, nums)
			wg.Done()
		}()

		go func() {
			c = solve(join(value, nums[index]), target, index+1, nums)
			wg.Done()
		}()

		wg.Wait()
	} else {
		a = solve(value*nums[index], target, index+1, nums)
		b = solve(value+nums[index], target, index+1, nums)
		c = solve(join(value, nums[index]), target, index+1, nums)
	}

	return a || b || c
}

func join(a, b int) int {
	strNum := fmt.Sprintf("%d%d", a, b)
	return utils.MustParseNum(strNum)
}
