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

	filledSeats := map[int]bool{}

	for _, line := range strings.Split(strData, "\n") {
		seat := calcSeat(line)
		filledSeats[seat] = true
	}

	for i := 0; i <= 127*8+7; i++ {
		if _, ok := filledSeats[i]; !ok {
			if i > 0 && i < 127*8+7 {
				_, preSeatFilled := filledSeats[i-1]
				_, postSeatFilled := filledSeats[i+1]
				if preSeatFilled && postSeatFilled {
					fmt.Println(i)
				}
			}
		}
	}
}

func calcSeat(line string) int {
	row := find(0, 127, line[:7])
	col := find(0, 7, line[7:])
	return row*8 + col
}

func find(min, max int, query string) int {
	first := query[0]
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
