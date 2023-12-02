package day1

import (
	"log"
	"strconv"
	"unicode"

	"cosme.dev/aoc2023/helpers"
)

func Problem1() {
    lines, err := helpers.ReadFileLines("day1/problem1.txt")
    if err != nil {
        log.Fatalf("Error reading input file %v", err)
    }

    numbers := make([]int, len(lines))

    for i, line := range lines {
        l := 0
        r := len(line) - 1
        for c := 0; c < len(line); c++ {
            if !unicode.IsDigit(rune(line[l])) {
                l++
            }

            if !unicode.IsDigit(rune(line[r])) {
                r--
            }
        }

        stringNumber := string(line[l]) + string(line[r])
        numbers[i], err = strconv.Atoi(stringNumber)
        if err != nil {
            log.Fatalf("Can not convert %v to integer: %v", stringNumber, err)
        }
    }

    total := 0
    for _, number := range numbers {
        total += number
    }

    log.Print(total)
}
