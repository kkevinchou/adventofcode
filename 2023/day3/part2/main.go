package main

import (
	"fmt"

	"github.com/kkevinchou/adventofcode/utils"
)

type Data struct {
	ID  int
	Num int
}

var numIDGenerator int = 100

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\n")
	lookup := [][]Data{}

	var rCount, cCount int
	var total int

	for {
		record, done := generator()
		if done {
			break
		}

		line := record.Line
		lookup = append(lookup, make([]Data, len(line)))
		rCount = len(line)

		var i int
		for i < len(line)-1 {
			if utils.IsNum(string(line[i])) {
				num, length := parseNum(line, i)

				for j := 0; j < length; j++ {
					lookup[len(lookup)-1][i+j] = Data{
						ID:  numIDGenerator,
						Num: num,
					}
				}
				numIDGenerator += 1

				i += length
			} else {
				i += 1
			}
		}
		cCount += 1
	}

	generator = utils.RecordGenerator(file, "\n")
	var r int
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.Line
		for i := 0; i < len(line); i++ {
			char := string(line[i])
			if char != "*" {
				continue
			}
			seen := map[int]bool{}

			partCount := 0
			gearRatio := 1

			neighbors := getNeighbors(r, i, rCount, cCount)
			for _, neighbor := range neighbors {
				nr, nc := neighbor[0], neighbor[1]
				data := lookup[nr][nc]
				if data.ID != 0 {
					if _, ok := seen[data.ID]; ok {
						continue
					}
					seen[data.ID] = true
					gearRatio *= data.Num
					partCount += 1

				}
			}

			if partCount == 2 {
				total += gearRatio
			}
		}

		r += 1
	}
	fmt.Println(total)
}

func getNeighbors(r, c, rCount, cCount int) [][2]int {
	dirs := [][2]int{{1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	var neighbors [][2]int
	for _, dir := range dirs {
		nr, nc := r+dir[0], c+dir[1]

		if nr < 0 || nr >= rCount {
			continue
		}

		if nc < 0 || nc >= cCount {
			continue
		}

		neighbors = append(neighbors, [2]int{nr, nc})
	}
	return neighbors
}

// returns the number, and it's length
func parseNum(line string, index int) (int, int) {
	digitsSoFar := ""
	for i := index; i < len(line); i++ {
		if utils.IsNum(string(line[i])) {
			digitsSoFar += string(line[i])
		} else {
			break
		}
	}

	return utils.MustParseNum(digitsSoFar), len(digitsSoFar)
}
