package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	strData := string(data)

	maxSeat := 0

	for _, line := range strings.Split(strData, "\n") {
		seat := calcSeat(line)
		if seat > maxSeat {
			maxSeat = seat
		}
	}

	fmt.Println(maxSeat)
}

func calcSeat(line string) int {
	row := find(0, 127, line[:7])
	col := find(0, 7, line[7:])
	return row*8 + col
}

func find(min, max int, query string) int {
	first := query[0]
	fmt.Println(min, max, string(first), query)
	if max-min == 1 {
		if first == 'F' || first == 'L' {
			return min
		} else if first == 'B' || first == 'R' {
			return max
		} else {
			panic("WTF1")
		}
	}

	if first == 'F' || first == 'L' {
		return find(min, (max+min-1)/2, query[1:])
	} else if first == 'B' || first == 'R' {
		return find((max+min+1)/2, max, query[1:])
	}

	panic("WTF2")
	return -1
}
