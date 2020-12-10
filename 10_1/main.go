package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{}
	recordGenerator := constructRecordGenerator("input", "\n")
	for {
		record, done := recordGenerator()
		if done {
			break
		}
		nums = append(nums, mustParseNum(record.Lines[0]))
	}

	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3)

	lastJolt := 0
	oneJolt := 0
	threeJolt := 0
	for _, num := range nums {
		if num-lastJolt == 1 {
			oneJolt++
		} else if num-lastJolt == 3 {
			threeJolt++
		}
		lastJolt = num
	}

	fmt.Println(oneJolt * threeJolt)
}
