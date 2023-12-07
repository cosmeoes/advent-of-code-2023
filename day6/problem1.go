package day6

import (
	"log"
	"strconv"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

type Race struct {
    duration int
    distance int
}

func Problem1() {
    lines, err := helpers.ReadFileLines("day6/input.txt")
    if err != nil {
	log.Fatalf("Error reading input file %v", err)
    }

    durations := strings.Fields(lines[0])[1:]
    distances := strings.Fields(lines[1])[1:]

    races := make([]Race, len(durations))
    for i, duration := range durations {
	parsedDuartion, _ := strconv.Atoi(duration)
	parsedDistance, _ := strconv.Atoi(distances[i])
	races[i] = Race{
	    duration: parsedDuartion,
	    distance: parsedDistance,
	}
    }

    result := 1
    for _, race := range races {
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
    }

    log.Print(result)
}

func calculateDistance(race Race, pressTime int) int {
    timeToRun := race.duration - pressTime

    return timeToRun * pressTime
}
