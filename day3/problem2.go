package day3

import (
	"log"

	"cosme.dev/aoc2023/helpers"
)

func Problem2()  {
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
            if line[col] == '*' {
                nearby := GetNearByNumbers(row, col, &lines)
                if len(nearby) == 2 {
                    sum += nearby[0] * nearby[1]
                }
            }
        }
    }

    log.Printf("Result is %v", sum)
}
