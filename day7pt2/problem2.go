package day7pt2

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

type HandType uint

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

type Card uint

const (
	J Card = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	T
	Q
	K
	A
)

var cardValues = map[rune]Card{
	'2': Two,
	'3': Three,
	'4': Four,
	'5': Five,
	'6': Six,
	'7': Seven,
	'8': Eight,
	'9': Nine,
	'T': T,
	'J': J,
	'Q': Q,
	'K': K,
	'A': A,
}

type Hand struct {
	handType HandType
	cards    [5]Card
	bid      int
}

func Problem2() {
	lines, err := helpers.ReadFileLines("day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	hands := make([]Hand, len(lines))
	for i, line := range lines {
		hands[i] = parseHand(line)
	}

	sortHands(hands)

	res := 0
	for i, hand := range hands {
		res += hand.bid * (i + 1)
	}

	log.Print(res)
}

func parseHand(line string) Hand {
	data := strings.Fields(line)

	handText := data[0]
	bid, _ := strconv.Atoi(data[1])

	var cards [5]Card
	cardCounts := make([]Card, 13)
	jCount := 0
	for i, char := range handText {
		cards[i] = cardValues[char]
		if (cards[i] == J) {
			jCount++
		} else {
			cardCounts[cards[i]]++
		}
	}

	sort.Slice(cardCounts, func (i, j int) bool {
		return cardCounts[i] > cardCounts[j]
	})

	var handType HandType = HighCard
	for _, count := range cardCounts {
		switch count {
		case 5:
			handType = FiveOfKind
			break
		case 4:
			handType = FourOfKind
			break
		case 3:
			handType = ThreeOfKind
		case 2:
			if handType == ThreeOfKind {
				handType = FullHouse
			} else if handType == OnePair {
				handType = TwoPair
			} else {
				handType = OnePair
			}
		}
	}

	switch handType {
	case FourOfKind:
		if jCount >= 1 {
			handType = FiveOfKind
		}
	case ThreeOfKind:
		if jCount >= 2 {
			handType = FiveOfKind
		} else if  jCount == 1 {
			handType = FourOfKind
		}
	case TwoPair:
		if jCount >= 1 {
			handType = FullHouse
		}
	case OnePair:
		if jCount >= 3 {
			handType = FiveOfKind
		} else if jCount >= 2 {
			handType = FourOfKind
		} else if jCount >= 1 {
			handType = ThreeOfKind
		}
	case HighCard:
		if jCount >= 4 {
			handType = FiveOfKind
		} else if jCount >= 3 {
			handType = FourOfKind
		} else if jCount >= 2 {
			handType = ThreeOfKind
		} else if jCount >= 1 {
			handType = OnePair
		}
	}


	return Hand{
		handType: handType,
		cards: cards,
		bid:   bid,
	}
}

func sortHands(hands []Hand) {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType < hands[j].handType {
			return true
		} else if hands[i].handType > hands[j].handType {
			return false
		}

		for k := 0; k < 5; k++ {
			cardI, cardJ := hands[i].cards[k], hands[j].cards[k]

			if cardI < cardJ {
				return true
			} else if cardI > cardJ {
				return false
			}
		}

		return true
	})
}
