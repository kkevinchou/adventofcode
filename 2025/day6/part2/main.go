package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var file string = "input.txt"

func main() {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	ops := []string{}
	widths := []int{}

	rCount := 0

	lines := strings.Split(string(data), "\r\n")
	for _, line := range lines {
		if string(line[0]) != "*" && string(line[0]) != "+" {
			rCount++
		} else {
			index := -1
			for _, s := range line {
				c := string(s)
				if c == " " {
					widths[index]++
				}

				if c == "*" || c == "+" {
					widths = append(widths, 0)
					ops = append(ops, c)
					index++
				}
			}
			widths[index]++
		}
	}

	globalStrNums := [][]string{}
	for _, line := range lines {
		if string(line[0]) == "*" || string(line[0]) == "+" {
			break
		}
		i := 0
		strNums := []string{}
		for _, width := range widths {
			strNum := ""
			for range width {
				strNum += string(line[i])
				i++
			}
			// skip the vertical column space
			i++
			strNums = append(strNums, strNum)
		}
		globalStrNums = append(globalStrNums, strNums)
	}

	cCount := len(ops)
	var result int

	for c := range cCount {
		width := widths[c]

		var nums []int
		for w := width - 1; w >= 0; w-- {
			strNum := ""
			for r := range rCount {
				strNum += string(globalStrNums[r][c][w])
			}
			num, _ := strconv.Atoi(strings.TrimSpace(strNum))
			nums = append(nums, num)
		}

		var columnResult int
		if ops[c] == "*" {
			columnResult = 1
		}

		for _, num := range nums {
			if ops[c] == "*" {
				columnResult *= num
			} else {
				columnResult += num
			}
		}

		result += columnResult
	}

	fmt.Println(result)
}
