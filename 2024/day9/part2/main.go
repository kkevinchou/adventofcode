package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

var file string = "input"

type Coordinate [2]int

func main() {
	start := time.Now()
	var chars []string
	for record := range utils.Records(file) {
		chars = strings.Split(record.Line, "")
		break
	}
	nums := utils.StringSliceToIntSlice(chars)

	positions := make([]int, len(nums))
	var positionSoFar int
	for i, num := range nums {
		positions[i] = positionSoFar
		positionSoFar += num
	}

	var result int

	for right := len(nums) - 1; right > 0; right -= 2 {
		candidate := nums[right]

		for left := 1; left < len(nums); left += 2 {
			if right < left {
				break
			}
			if nums[left] >= candidate {
				nums[left] -= candidate
				nums[right] = 0

				for _ = range candidate {
					result += (positions[left] * (right / 2))
					positions[left]++
				}
				break
			}
		}
	}

	for left := 0; left < len(nums); left += 2 {
		num := nums[left]
		for _ = range num {
			result += (positions[left] * (left / 2))
			positions[left]++
		}
	}

	fmt.Println(result)
	fmt.Println(time.Since(start))
}
