package day5

import (
	"log"
	"math"
	"strconv"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

type Range struct {
    destStart int
    sourceStart int
    rang int
}

func Problem1() {
    lines, err := helpers.ReadFileLines("day5/input.txt")
    if err != nil {
	log.Fatalf("Error reading input file %v", err)
    }


    seeds := strings.Fields(lines[0])[1:]

    var maps [7][]Range

    var currentMapIndex int
    for i := 1; i < len(lines); i++ {
	if lines[i] == "" {
	    continue
	}

	switch lines[i] {
	case "seed-to-soil map:":
	    currentMapIndex = 0
	case "soil-to-fertilizer map:":
	    currentMapIndex = 1
	case "fertilizer-to-water map:":
	    currentMapIndex = 2
	case "water-to-light map:":
	    currentMapIndex = 3
	case "light-to-temperature map:":
	    currentMapIndex = 4
	case "temperature-to-humidity map:":
	    currentMapIndex = 5
	case "humidity-to-location map:":
	    currentMapIndex = 6
	default:
	    data := strings.Fields(lines[i])
	    destStart, _ := strconv.Atoi(data[0])
	    sourceStart, _ := strconv.Atoi(data[1])
	    rang, _ := strconv.Atoi(data[2])
	    maps[currentMapIndex] = append(maps[currentMapIndex], Range{
		destStart: destStart,
		sourceStart: sourceStart,
		rang: rang,
	    })
	}
    }

    minLoc := math.MaxInt
    for _, s := range seeds {
	seed, _ := strconv.Atoi(s)
	soil := getValue(seed, maps[0])
	fert := getValue(soil, maps[1]);
	water := getValue(fert, maps[2])
	light := getValue(water, maps[3])
	temp := getValue(light, maps[4])
	hum := getValue(temp, maps[5])
	loc := getValue(hum, maps[6])
	if loc < minLoc {
	    minLoc = loc
	}

	log.Print([]int{seed, soil, fert, water, light, temp, hum, loc})
    }

    log.Printf("minLoc: %v", minLoc)
}

func getValue(source int, typeMap []Range) int {
    for _, r := range typeMap {
	if source >= r.sourceStart && source <= r.sourceStart + r.rang {
	    return r.destStart + (r.sourceStart - source)
	}
    }

    return source
}
