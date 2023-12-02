package day2

import (
	"log"
	"strconv"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

func Problem1() {
    lines, err := helpers.ReadFileLines("day2/problem1.txt")
    if err != nil {
        log.Fatalf("Error reading input file %v", err)
    }

    existingCubes := map[string]int{
        "red": 12,
        "green": 13,
        "blue": 14,
    }

    idSum := 0
    for _, gameText := range lines {
        game := parseGame(gameText)

        if game.redCount <= existingCubes["red"] && 
        game.greenCount <= existingCubes["green"] &&
        game.blueCount <= existingCubes["blue"] {
            idSum += game.id
        }
    }

    log.Printf("Possible games sum is %v", idSum)
}

type Game struct {
    id int
    redCount int
    greenCount int
    blueCount int
}

func parseGame(game string) Game {
    data := strings.Split(game, ":")
    idText := strings.Replace(data[0], "Game ", "", 1)
    // Ignore error because idc
    id, _ := strconv.Atoi(idText)

    maxRed, maxGreen, maxBlue := 0, 0, 0

    subsets := strings.Split(data[1], "; ")
    for _, subset := range subsets {
        sets := strings.Split(strings.TrimSpace(subset), ", ")

        for _, set := range sets {
            cubesInfo := strings.Split(set, " ");
            count, _ := strconv.Atoi(cubesInfo[0])

            switch cubesInfo[1] {
            case "red":
                if maxRed < count {
                    maxRed = count
                }
            case "green":
                if maxGreen < count {
                    maxGreen = count
                }
            case "blue":
                if maxBlue < count {
                    maxBlue = count
                }
            }
        }
    }

    return Game{
        id: id,
        redCount: maxRed,
        greenCount: maxGreen,
        blueCount: maxBlue,
    }
}
