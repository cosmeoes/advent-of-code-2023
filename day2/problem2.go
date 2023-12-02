package day2

import (
	"log"

	"cosme.dev/aoc2023/helpers"
)

func Problem2() {
    lines, err := helpers.ReadFileLines("day2/problem1.txt")
    if err != nil {
        log.Fatalf("Error reading input file %v", err)
    }

    powerSum := 0
    for _, gameText := range lines {
        game := parseGame(gameText)
        powerSum += game.redCount * game.greenCount * game.blueCount
    }

    log.Printf("Possible games sum is %v", powerSum)
}
