package day4

import (
	"log"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

func Problem2() {
    lines, err := helpers.ReadFileLines("day4/input.txt")
    if err != nil {
	log.Fatalf("Error reading input file %v", err)
    }

    var dupCounts = make([]int, len(lines))
    for i := range dupCounts {
	dupCounts[i] = 1
    }

    for i, line := range lines {
	numbers := strings.Split(line, ": ")[1]

	winningNumbers := strings.Fields(strings.TrimSpace(strings.Split(numbers, " | ")[0]))
	cardNumbers := strings.Fields(strings.TrimSpace(strings.Split(numbers, " | ")[1]))

	foundCount := 0
	for _, cardNumber := range cardNumbers {
	    if !contains(winningNumbers, cardNumber) {
		continue
	    }

	    foundCount++
	}

	dupsOfCurrentCard := dupCounts[i]
	for pos := i + 1; pos <= i + foundCount; pos++ {
	    dupCounts[pos] += dupsOfCurrentCard
	}
    }

    totalCards := 0
    for _, dups := range dupCounts {
	totalCards += dups
    }

    log.Printf("Total cards %v", totalCards)
}
