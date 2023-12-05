package main

import "fmt"

type Preamble struct {
	numList []int
}

func main() {
	preamble := Preamble{}
	recordGenerator := constructRecordGenerator("input", "\n")
	for {
		record, done := recordGenerator()
		if done {
			break
		}

		num := mustParseNum(record.Lines[0])

		if record.ID < 25 {
			preamble.numList = append(preamble.numList, num)
			continue
		}

		if !preamble.InPreamble(num) {
			fmt.Println(num)
			break
		}

		preamble.numList = preamble.numList[1:]
		preamble.numList = append(preamble.numList, num)
	}
}

func (p *Preamble) InPreamble(num int) bool {
	for _, num1 := range p.numList {
		for _, num2 := range p.numList {
			if num1 != num2 {
				if num1+num2 == num {
					return true
				}
			}
		}
	}

	return false
}
