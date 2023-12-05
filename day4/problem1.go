package day4

import (
	"log"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

func Problem1() {
    lines, err := helpers.ReadFileLines("day4/input.txt")
    if err != nil {
	log.Fatalf("Error reading input file %v", err)
    }

    var score int
    for _, line := range lines {
	numbers := strings.Split(line, ": ")[1]

	winningNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, " | ")[0]), " ")
	cardNumbers := strings.Split(strings.TrimSpace(strings.Split(numbers, " | ")[1]), " ")

	var cardScore int
	for _, cardNumber := range cardNumbers {
	    if !contains(winningNumbers, cardNumber) {
		continue
	    }

	    if cardScore == 0 {
		cardScore = 1
	    } else {
		cardScore *= 2
	    }

	}

	score += cardScore
    }

    log.Printf("Total Score %v", score)
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}
