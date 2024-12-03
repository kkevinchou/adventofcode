package main

import (
	"fmt"
	"os"

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

// length = 4
// 0 1 2 3
// a m u l
func readMul(i int, s string) bool {
	if i >= len(s)-2 {
		return false
	}

	return s[i:i+3] == "mul"
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	strData := string(data)

	var num1, num2 int
	var success bool
	var result int
	var skip int

	for i := 0; i < len(strData); i++ {
		if success = readMul(i, strData); !success {
			continue
		}
		i += 3
		if success = readLeftParen(i, strData); !success {
			continue
		}
		i += 1
		if num1, skip, success = readNum(i, strData); !success {
			continue
		}
		i += skip
		if success = readComma(i, strData); !success {
			continue
		}
		i += 1
		if num2, skip, success = readNum(i, strData); !success {
			continue
		}
		i += skip
		if success = readRightParen(i, strData); !success {
			continue
		}
		result += num1 * num2
	}

	fmt.Println(result)
}
