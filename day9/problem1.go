package day9

import (
	"log"
	"strconv"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

func Problem1() {
    lines, err := helpers.ReadFileLines("day9/input.txt")
    if err != nil {
	log.Fatalf("Error reading input: %v", err)
    }

    inputs := make([][]int, len(lines))
    for i, line := range lines {
	inputs[i] = parseLine(line)
    }

    result := 0
    for _, input := range inputs {
	result += nextValue(input)
    }

    log.Printf("Result: %v", result)
}

func parseLine(line string) []int {
    stringValues := strings.Fields(line)
    values := make([]int, len(stringValues))

    for i, v := range stringValues {
	values[i], _ = strconv.Atoi(v)
    }

    return values
}

func nextValue(input []int) int {
    out := make([]int, 0, len(input) - 1)

    allZeros := true
    for i := 0; i < len(input) - 1; i += 1 {
	val := input[i + 1] - input[i]
	if val != 0 {
	    allZeros = false
	}

	out = append(out, val)
    }

    res := input[len(input) - 1]
    if allZeros {
	return res
    }

    return res + nextValue(out)
}
