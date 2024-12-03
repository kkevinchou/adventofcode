package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	start := time.Now()
	r := regexp.MustCompile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
	r2 := regexp.MustCompile(`\d{1,3}`)
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}
	strData := string(data)

	var result int
	do := true

	for {
		loc := r.FindStringIndex(strData)
		if loc == nil {
			break
		}
		matchStr := strData[loc[0]:loc[1]]
		strData = strData[loc[1]:]

		if do && strings.HasPrefix(matchStr, "mul") {
			strNums := r2.FindAllString(matchStr, 2)
			args := utils.StringSliceToIntSlice(strNums)
			result += args[0] * args[1]
		}

		if matchStr == "do()" {
			do = true
		} else if matchStr == "don't()" {
			do = false
		}
	}
	fmt.Println(result)
	fmt.Println(time.Since(start))
}
