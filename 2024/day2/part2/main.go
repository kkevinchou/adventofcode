package main

import (
	"fmt"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

var infile string = "sample2"

func main() {
	naiveButWorks()
}

func cleverButBusted() {
	var result int

	for record := range utils.Records(infile, "\n") {
		stringSplit := strings.Split(record.Line, " ")
		nums := utils.StringSliceToIntSlice(stringSplit)

		// A B C D
		var safe bool
		for i := range nums {
			var newSlice []int
			for j := range nums {
				if i == j {
					continue
				}
				newSlice = append(newSlice, nums[j])
			}
			if solve(newSlice) {
				safe = true
				result++
				break
			}
		}

		if safe {
			fmt.Println("SAFE", nums)
		} else {
			fmt.Println("NOT SAFE", nums)
		}
	}

	fmt.Println(result)
}

func solve(nums []int) bool {
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
	return safe
}

func naiveButWorks() {
	var result int

	for record := range utils.Records(infile, "\n") {
		stringSplit := strings.Split(record.Line, " ")

		nums := utils.StringSliceToIntSlice(stringSplit)

		var asc int
		var desc int
		var eq int
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] < nums[i+1] {
				asc++
			} else if nums[i] > nums[i+1] {
				desc++
			} else {
				eq++
			}
		}

		var safeF safeFn
		if eq > 2 {
			// fmt.Println(nums)
			continue
		} else if asc > desc {
			safeF = incSafe
		} else {
			safeF = decSafe
		}

		safe := true
		var mistake bool
		left, right := 0, 1
		skipIndex := -1
		for {
			if right >= len(nums) {
				break
			}
			if left == skipIndex {
				left++
			}
			if !safeF(nums[left], nums[right]) {
				if mistake {
					safe = false
					break
				}
				mistake = true

				// drop the last element
				if right == len(nums)-1 {
					break
				}

				skipIndex = left
				left--
				if left == -1 {
					left = 1
					right = 2
				}
				continue
			}

			delta := utils.IntAbs(nums[left] - nums[right])
			if delta < 1 || delta > 3 {
				if mistake {
					safe = false
					break
				}
				mistake = true

				// drop the last element
				if right == len(nums)-1 {
					break
				}
				skipIndex = left
				left--
				if left == -1 {
					left = 1
					right = 2
				}
				continue
			}
			left++
			right++
		}

		if safe {
			fmt.Println("SAFE", nums)
			result++
		} else {
			fmt.Println("NOT SAFE", nums)
		}
	}
	fmt.Println(result)
}

type safeFn func(i, j int) bool

func incSafe(i, j int) bool {
	return i < j
}

func decSafe(i, j int) bool {
	return i > j
}
