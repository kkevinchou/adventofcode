package main

import (
	"fmt"
)

func main() {
	recordGenerator := constructRecordGenerator("input", "\n")
	var target int = 32321523 // from part 1

	var numList []int
	for {
		record, done := recordGenerator()
		if done {
			break
		}

		num := mustParseNum(record.Lines[0])
		numList = append(numList, num)
	}

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if i >= j {
				continue
			}

			if sum(numList, i, j) == target {
				smallest := 9999999999
				largest := 0
				for y := i; y < j; y++ {
					if numList[y] < smallest {
						smallest = numList[y]
					}
					if numList[y] > largest {
						largest = numList[y]
					}
				}
				fmt.Println(smallest + largest)
				return
			}
		}
	}
}

func sum(numList []int, i, j int) int {
	var sum int
	for x := i; x < j; x++ {
		sum += numList[x]
	}
	return sum
}
