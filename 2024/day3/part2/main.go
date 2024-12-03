package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

func readNum(i int, s string) (int, int, bool) {
	var strNum string

	for index := i; index < len(s) && len(strNum) < 3; index++ {
		char := string(s[index])
		if _, err := utils.ParseNum(char); err == nil {
			strNum += char
		} else {
			break
		}
	}
	if strNum == "" {
		return -1, -1, false
	}

	num := utils.MustParseNum(strNum)
	return num, len(strNum), true
}

func readDo(i int, s string) bool {
	return strings.HasPrefix(s[i:], "do()")
}

func readDont(i int, s string) bool {
	return strings.HasPrefix(s[i:], "don't()")
}

func readLeftParen(i int, s string) bool {
	if i == len(s) {
		return false
	}
	return string(s[i]) == "("
}

func readRightParen(i int, s string) bool {
	if i == len(s) {
		return false
	}
	return string(s[i]) == ")"
}

func readComma(i int, s string) bool {
	if i == len(s) {
		return false
	}
	return string(s[i]) == ","
}

func readMul(i int, s string) bool {
	if i >= len(s)-2 {
		return false
	}

	return s[i:i+3] == "mul"
}

func main() {
	start := time.Now()
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	strData := string(data)

	var num1, num2 int
	var success bool
	var result int
	var skip int

	do := true

	for i := 0; i < len(strData); i++ {
		// process don't
		if readDont(i, strData) {
			do = false
			i += len("don't()")
		}

		// process do
		if readDo(i, strData) {
			do = true
			i += len("do()")
		}

		// processing a mul
		if success = readMul(i, strData); !success {
			continue
		}
		i += 3
		if success = readLeftParen(i, strData); !success {
			if string(strData[i]) == "d" {
				i--
			}
			continue
		}
		i += 1
		if num1, skip, success = readNum(i, strData); !success {
			if string(strData[i]) == "d" {
				i--
			}
			continue
		}
		i += skip
		if success = readComma(i, strData); !success {
			if string(strData[i]) == "d" {
				i--
			}
			continue
		}
		i += 1
		if num2, skip, success = readNum(i, strData); !success {
			if string(strData[i]) == "d" {
				i--
			}
			continue
		}
		i += skip
		if success = readRightParen(i, strData); !success {
			if string(strData[i]) == "d" {
				i--
			}
			continue
		}

		if do {
			result += num1 * num2
		}
	}

	fmt.Println(result)
	fmt.Println(time.Since(start))
}
