
package day6

import (
	"log"
	"strconv"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

func Problem2() {
    lines, err := helpers.ReadFileLines("day6/input.txt")
    if err != nil {
	log.Fatalf("Error reading input file %v", err)
    }

    duration := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
    distance := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")
    parsedDuartion, _ := strconv.Atoi(duration)
    parsedDistance, _ := strconv.Atoi(distance)

    race := Race{
	duration: parsedDuartion,
	distance: parsedDistance,
    }

    log.Println(race)
    result := 1
    startRange := -1
    endRange := 0
    for i := 0; i < race.duration; i++ {
	dist := calculateDistance(race, i)

	if dist > race.distance && startRange == -1 {
	    startRange = i
	}

	if startRange != -1 && dist <= race.distance {
	    endRange = i
	    break
	}
    }

    result *= endRange - startRange

    log.Print(result)
}
