package day3

import (
	"log"
	"os"
	"strconv"
	"unicode"

	"cosme.dev/aoc2023/helpers"
)

func Problem1() {
    inputLines, err := helpers.ReadFileLines("day3/input.txt")
    if err != nil {
        log.Fatalf("Error reading input file %v", err)
    }

    lines := make([][]byte, len(inputLines))

    for i, line := range inputLines {
        lines[i] = []byte(line)
    }

    var sum int
    for row, line := range lines {
        for col := 0; col < len(line); col++ {
            if line[col] != '.' && !unicode.IsDigit(rune(line[col])) {
                nearby := GetNearByNumbers(row, col, &lines)
                for _, number := range nearby {
                    sum += number
                }
            }
        }
    }

    writeDebugOutputFile(lines)
    log.Printf("Result is %v", sum)
}

func GetNearByNumbers(row, col int, lines *[][]byte) []int {
    l := *lines

    var res []int
    if row > 0 {
        // top left
        if col > 0 && unicode.IsDigit(rune(l[row - 1][col - 1])) {
            res = append(res, GetFullNumber(row - 1, col - 1, lines))
        }

        // top center
        if unicode.IsDigit(rune(l[row - 1][col])) {
            res = append(res, GetFullNumber(row - 1, col, lines))
        }

        // top right
        if col < len(l[row]) - 1 && unicode.IsDigit(rune(l[row - 1][col + 1])) {
            res = append(res, GetFullNumber(row - 1, col + 1, lines))
        }
    }

    if row < len(l) - 1 {
        // bottom left
        if col > 0 && unicode.IsDigit(rune(l[row + 1][col - 1])) {
            res = append(res, GetFullNumber(row + 1, col - 1, lines))
        }

        // bottom center
        if unicode.IsDigit(rune(l[row + 1][col])) {
            res = append(res, GetFullNumber(row + 1, col, lines))
        }

        // bottom right
        if col < len(l[row]) - 1 && unicode.IsDigit(rune(l[row + 1][col + 1])) {
            res = append(res, GetFullNumber(row + 1, col + 1, lines))
        }
    }

    //  left
    if col > 0 && unicode.IsDigit(rune(l[row][col - 1])) {
        res = append(res, GetFullNumber(row, col - 1, lines))
    }

    // right
    if col < len(l[row]) - 1 && unicode.IsDigit(rune(l[row][col + 1])) {
        res = append(res, GetFullNumber(row, col + 1, lines))
    }

    return res
}

func GetFullNumber(row, col int, lines *[][]byte) int {
    l := *lines
    start := col

    for start >= 0 && unicode.IsDigit(rune(l[row][start])) {
        start--
    }

    start++

    end := col

    for end < len(l[row]) && unicode.IsDigit(rune(l[row][end])) {
        end++
    }

    numberText := string(l[row][start:end])
    value, _ := strconv.Atoi(numberText)

    for i := start; i < end; i++ {
        (*lines)[row][i] = '.'
    }


    return value
}

func writeDebugOutputFile(lines [][]byte) {
    // flat := make([]byte, len(lines[0]) * len(lines))
    fo, _ := os.Create("debug_output.txt")

    for _, chunk := range lines {
        fo.WriteString(string(chunk) + "\n")
    }

    fo.Close()
}
