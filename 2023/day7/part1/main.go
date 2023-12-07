package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/kkevinchou/adventofcode/utils"
)

func main() {
	file := "input"
	generator := utils.RecordGenerator(file, "\r\n")

	var hands []Hand
	for {
		record, done := generator()
		if done {
			break
		}

		line := record.SingleLine
		lineSplit := strings.Split(line, " ")
		cards := lineSplit[0]
		bid := utils.MustParseNum(lineSplit[1])
		hands = append(hands, Hand{Cards: cards, Type: typeFromCards(cards), Bid: bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		return lessThan(hands[i], hands[j])
	})

	total := 0
	for i, hand := range hands {
		rank := i + 1
		total += rank * hand.Bid
	}
	fmt.Println(hands)
	fmt.Println(total)
}

type Hand struct {
	Cards string
	Type  int
	Bid   int
}

func typeFromCards(cards string) int {
	cardCounts := map[string]int{}
	for _, card := range cards {
		cardCounts[string(card)] += 1
	}

	if len(cardCounts) == 5 {
		// high card
		return 1
	} else if len(cardCounts) == 4 {
		// one pair
		return 2
	} else if len(cardCounts) == 3 {
		// two pair, three of a kind
		threeOfAKind := false
		for _, c := range cardCounts {
			if c == 3 {
				threeOfAKind = true
				break
			}
		}
		if threeOfAKind {
			return 4
		} else {
			return 3
		}
	} else if len(cardCounts) == 2 {
		// four of a kind, full house
		fullHouse := false
		for _, c := range cardCounts {
			if c == 3 {
				fullHouse = true
			}
		}

		if fullHouse {
			return 5
		} else {
			return 6
		}
	} else if len(cardCounts) == 1 {
		return 7
	}

	panic("wat")
}

var strengthMap map[string]int = map[string]int{
	"2": 0,
	"3": 1,
	"4": 2,
	"5": 3,
	"6": 4,
	"7": 5,
	"8": 6,
	"9": 7,
	"T": 8,
	"J": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}

func lessThan(h1, h2 Hand) bool {
	if h1.Type < h2.Type {
		return true
	} else if h1.Type > h2.Type {
		return false
	}

	for i := 0; i < len(h1.Cards); i++ {
		if strengthMap[string(h1.Cards[i])] < strengthMap[string(h2.Cards[i])] {
			return true
		} else if strengthMap[string(h1.Cards[i])] > strengthMap[string(h2.Cards[i])] {
			return false
		}
	}
	panic("wat")
}
