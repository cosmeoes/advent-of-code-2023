package day5

import (
	"log"
	"math"
	"strconv"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

type SeedRange struct {
    start int
    rang int
}

func Problem2() {
    lines, err := helpers.ReadFileLines("day5/input.txt")
    if err != nil {
	log.Fatalf("Error reading input file %v", err)
    }


    seedRanges := strings.Fields(lines[0])[1:]
    seeds := make([]SeedRange, 0, len(seedRanges) / 2)
    for i := 0; i < len(seedRanges); i += 2 {
	seedStart, err := strconv.Atoi(seedRanges[i])
	if err != nil  {
	    log.Fatalf("Error parsing number %v: %v", seedRanges[i], err)
	}
	seedRange, err := strconv.Atoi(seedRanges[i + 1])
	if err != nil  {
	    log.Fatalf("Error parsing number %v: %v", seedRanges[i + 1], err)
	}
	seeds = append(seeds, SeedRange{
	    start: seedStart,
	    rang: seedRange,
	})

    }

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
    for _, seedRange := range seeds {
	for seed := seedRange.start; seed < seedRange.start + seedRange.rang; seed++ {
	    soil := getValue2(seed, 0, maps)
	    fert := getValue2(soil, 1, maps);
	    water := getValue2(fert, 2, maps)
	    light := getValue2(water, 3, maps)
	    temp := getValue2(light, 4, maps)
	    hum := getValue2(temp, 5, maps)
	    loc := getValue2(hum, 6, maps)
	    if loc < minLoc {
		minLoc = soil
	    }
	}
    }

    log.Printf("minLoc: %v", minLoc)
}

func getValue2(source, mapPos int, maps [7][]Range) int {
    typeMap := maps[mapPos]
    res := source
    for _, r := range typeMap {
	if source >= r.sourceStart && source < r.sourceStart + r.rang {
	    res = r.destStart + (source - r.sourceStart)
	    break
	}
    }

    return res
}
