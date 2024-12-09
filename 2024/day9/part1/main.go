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

	var result int
	var position int

	right := len(nums) - 1
	fileSize := nums[right]

	for left, num := range nums {
		if left == right {
			for _ = range fileSize {
				result += (position * left / 2)
				position++
			}
			break
		}

		if left%2 == 0 {
			// file block
			for _ = range num {
				result += (position * left / 2)
				position++
			}
		} else {
			// empty block
			emptySize := num
			for emptySize > 0 {
				for emptySize > 0 && fileSize > 0 {
					result += (position * right / 2)
					position++

					emptySize--
					fileSize--
				}

				if fileSize == 0 {
					right -= 2
					fileSize = nums[right]
				}
			}
		}
	}
	fmt.Println(result)
	fmt.Println(time.Since(start))
}
