package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{}
	recordGenerator := constructRecordGenerator("input", "\n")
	max := 0
	for {
		record, done := recordGenerator()
		if done {
			break
		}
		num := mustParseNum(record.Lines[0])
		if num > max {
			max = num
		}
		nums = append(nums, num)
	}

	nums = append(nums, max+3, 0)
	sort.Ints(nums)

	fmt.Println(calc(nums))
}

func calc(nums []int) int {
	countsFromPosition := map[int]int{}
	countsFromPosition[len(nums)-1] = 1
	countsFromPosition[len(nums)-2] = 1
	countsFromPosition[len(nums)-3] = 1

	for i := len(nums) - 4; i >= 0; i-- {
		countsFromPosition[i] = subCalc(nums, countsFromPosition, i)
	}

	fmt.Println(countsFromPosition)
	return countsFromPosition[0]
}

func subCalc(nums []int, countsFromPosition map[int]int, position int) int {
	total := 0
	for i := 0; i < 3; i++ {
		if nums[position] >= (nums[position+1+i] - 3) {
			total += countsFromPosition[position+1+i]
		}
	}

	return total
}
