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

	for record := range utils.Records(file, "\n") {
		strSplit := strings.Split(record.Line, ":")
		target := utils.MustParseNum(strSplit[0])
		nums := utils.StringSliceToIntSlice(strings.Split(strings.TrimSpace(strSplit[1]), " "))

		if solve(target, len(nums)-1, nums) {
			result += target
		}
	}

	fmt.Println(result)
	fmt.Println(time.Since(start))
}

// solves in reverse order of the numbers to enable quick short circuiting
func solve(target, index int, nums []int) bool {
	if index == 0 {
		return nums[index] == target
	}

	var a, b, c bool

	// mul
	if target%nums[index] == 0 {
		a = solve(target/nums[index], index-1, nums)
	}

	// add
	if target-nums[index] > 0 {
		b = solve(target-nums[index], index-1, nums)
	}

	// concat
	strTarget := fmt.Sprintf("%d", target)
	strNum := fmt.Sprintf("%d", nums[index])
	if strings.HasSuffix(strTarget, strNum) {
		strNewTarget := strTarget[:len(strTarget)-len(strNum)]

		if strNewTarget != "" {
			c = solve(utils.MustParseNum(strNewTarget), index-1, nums)
		}
	}

	return a || b || c
}
