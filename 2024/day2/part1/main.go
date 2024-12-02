package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	var result int
	for record := range utils.Records("input", "\n") {
		stringSplit := strings.Split(record.SingleLine, " ")

		nums := utils.StringSliceToIntSlice(stringSplit)

		var decreasing bool
		if nums[0] > nums[1] {
			decreasing = true
		}

		safe := true
		for i := 0; i < len(nums)-1; i++ {
			if decreasing {
				if nums[i] <= nums[i+1] {
					safe = false
					break
				}

				delta := utils.IntAbs(nums[i] - nums[i+1])
				if delta < 1 || delta > 3 {
					safe = false
					break
				}
			} else {
				if nums[i] >= nums[i+1] {
					safe = false
					break
				}

				delta := utils.IntAbs(nums[i] - nums[i+1])
				if delta < 1 || delta > 3 {
					safe = false
					break
				}
			}
		}
		if safe {
			result++
		}
	}
	fmt.Println(result)
}
