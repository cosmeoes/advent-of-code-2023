package day1

import (
	"log"
	"strconv"
	"unicode"

	"cosme.dev/aoc2023/helpers"
)

func Problem2() {
    lines, err := helpers.ReadFileLines("problem2.txt")
    if err != nil {
        log.Fatalf("Error reading input file %v", err)
    }

    // for i, line := range lines {
    //     newLine := strings.Builder{}
    //     for j := 0; j < len(line); j++ {
    //         var nextChar byte
    //         nextChar, j = toNumber(line, j)
    //         newLine.WriteByte(nextChar)
    //     }

    //     lines[i] = newLine.String()
    // }

    numbers := make([]int, len(lines))

    for i, line := range lines {
        l := 0
        r := len(line) - 1
        leftValue := ""
        rightValue := ""
        for c := 0; c < len(line); c++ {
            if leftValue == "" {
                var found bool
                leftValue, found = toNumber(line, l)
                if !found {
                    l++
                }
            }

            if rightValue == "" {
                var found bool
                rightValue, found = toNumberReversed(line, r)
                if !found {
                    r--
                }
            }
        }

        stringNumber := leftValue + rightValue
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

func toNumber(line string, index int) (string, bool) {
    if unicode.IsDigit(rune(line[index])) {
        return string(line[index]), true
    }

    writtenNumbers := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
    for word, value := range writtenNumbers {
        size := len(word)

        if len(line) >= size+index && line[index:size+index] == word {
            return value, true
        }
    }

    return "", false
}

func toNumberReversed(line string, index int) (string, bool) {
    if unicode.IsDigit(rune(line[index])) {
        return string(line[index]), true
    }

    writtenNumbers := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

    for word, value := range writtenNumbers {
        size := len(word)

        if index-size+1 >= 0 && line[index-size+1:index+1] == word {
            return value, true
        }
    }

    return "", false
}
