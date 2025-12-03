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

	var sum int

	lines := strings.Split(string(data), "\r\n")
	for _, line := range lines {
		var right, num = -1, -1
		for i := len(line) - 2; i >= 0; i-- {
			curLeft, _ := strconv.Atoi(string(line[i]))
			curRight, _ := strconv.Atoi(string(line[i+1]))

			left := curLeft
			right = max(right, curRight)

			curNum := 10*left + right
			if curNum > num {
				num = curNum
			}
		}
		sum += num
	}

	fmt.Println(sum)
}
