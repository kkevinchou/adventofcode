package main

import (
	"fmt"
	"strings"
)

func main() {
	recordGenerator := constructRecordGenerator("input", "\n")
	busses := map[int]int{}
	for {
		record, done := recordGenerator()
		if done {
			break
		}

		if record.ID == 1 {
			split := strings.Split(record.Lines[0], ",")
			for i, id := range split {
				if id == "x" {
					continue
				}
				busses[i] = mustParseNum(id)
			}
		}
	}

	fmt.Println(busses)

	step := 29 * 409
	var temp int
	for i := 0; i < 1000000000000000; i += step {
		temp = (i - 29)
		if i%(step*1000000000) == 0 {
			fmt.Println(i)
		}
		// if (i-29)%29 == 0 && (i-6)%37 == 0 && (i+17)%17 == 0 && (i+18)%13 == 0 && (i+19)%19 == 0 && (i+23)%23 == 0 && (i+31)%353 == 0 && (i+72)%41 == 0 {
		if (temp+23)%37 == 0 && (temp+46)%17 == 0 && (temp+47)%13 == 0 && (temp+48)%19 == 0 && (temp+52)%23 == 0 && (temp+60)%353 == 0 && (temp+101)%41 == 0 {
			fmt.Println("ANSWER:", temp)
			break
		}
	}

	// i := 438265956323733
	// if (i+23)%37 == 0 && (i+46)%17 == 0 && (i+47)%13 == 0 && (i+48)%19 == 0 && (i+52)%23 == 0 && (i+60)%353 == 0 && (i+101)%41 == 0 {
	// 	fmt.Println("YEP")
	// }

	// 	var nums []int
	// 	for i := 0; i <= 144377; i++ {
	// 		if (i-31)%409 == 0 {
	// 			nums = append(nums, i)
	// 		}
	// 	}
	// 	fmt.Println(nums)
	// 	fmt.Println(len(nums))
	// }
}
