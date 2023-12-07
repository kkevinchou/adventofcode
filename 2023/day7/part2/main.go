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

	if jCount, ok := cardCounts["J"]; ok {
		delete(cardCounts, "J")

		if len(cardCounts) == 0 {
			// 5 of a kind
			return 7
		}

		var maxCard string
		maxCount := 0
		for card, count := range cardCounts {
			if count > maxCount {
				maxCard = card
				maxCount = count
			}
		}

		cardCounts[maxCard] += jCount
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

var strengthMap map[string]int = map[string]int{
	"J": 0,
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"Q": 10,
	"K": 11,
	"A": 12,
}
