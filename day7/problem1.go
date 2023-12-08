package day7

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
	Two Card = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	T
	J
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

func Problem1() {
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
	for i, char := range handText {
		cards[i] = cardValues[char]
		cardCounts[cards[i]]++
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
