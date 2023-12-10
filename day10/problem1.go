package day10

import (
	"log"

	"cosme.dev/aoc2023/helpers"
	"cosme.dev/aoc2023/stack"
)

type Position struct {
    x, y, distanceFromS int
    origin *Position
}

func Problem1() {
    lines, err := helpers.ReadFileLines("day10/input.txt")
    if err != nil {
	log.Fatal(err)
    }

    var sPos Position
    for y := 0; y < len(lines); y++ {
	for x := 0; x < len(lines[y]); x++ {
	    if lines[y][x] == 'S' {
		sPos = Position{
		    x: x,
		    y: y,
		}
		break
	    }
	}
    }

    toVisit := stack.New[Position]()
    if sPos.y > 0 &&  isOneOf(lines[sPos.y - 1][sPos.x], "|F7") {
	toVisit.Push(Position{
	    x: sPos.x,
	    y: sPos.y - 1,
	    distanceFromS: 1,
	    origin: &sPos,
	})
    }

    if sPos.y < len(lines) - 1 &&  isOneOf(lines[sPos.y + 1][sPos.x], "|LJ") {
	toVisit.Push(Position{
	    x: sPos.x,
	    y: sPos.y + 1,
	    distanceFromS: 1,
	    origin: &sPos,
	})
    }

    if sPos.x > 0 && isOneOf(lines[sPos.y][sPos.x - 1], "-LF") {
	toVisit.Push(Position{
	    x: sPos.x - 1,
	    y: sPos.y,
	    distanceFromS: 1,
	    origin: &sPos,
	})
    }

    if sPos.x < len(lines[sPos.y]) - 1 && isOneOf(lines[sPos.y][sPos.x + 1], "-J7") {
	toVisit.Push(Position{
	    x: sPos.x + 1,
	    y: sPos.y,
	    distanceFromS: 1,
	    origin: &sPos,
	})
    }
    for toVisit.Len() > 0 {
	p, _ := toVisit.Pop()
	currentValue := lines[p.y][p.x]
	switch currentValue {
	case 'S':
	    log.Printf("Done: %v", p.distanceFromS / 2)
	    return
	case '|':
	    if p.y < p.origin.y {
		if p.y > 0 &&  isOneOf(lines[p.y - 1][p.x], "|F7") {
		    toVisit.Push(Position{
			x: p.x,
			y: p.y - 1,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    } else {
		if p.y < len(lines) - 1 &&  isOneOf(lines[p.y + 1][p.x], "|LJ") {
		    toVisit.Push(Position{
			x: p.x,
			y: p.y + 1,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    }
	case '-':
	    if p.x > p.origin.x {
		if p.x < len(lines[p.y]) - 1 &&  isOneOf(lines[p.y][p.x + 1], "-J7") {
		    toVisit.Push(Position{
			x: p.x + 1,
			y: p.y,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    } else {
		if p.x > 0 &&  isOneOf(lines[p.y][p.x - 1], "-LF") {
		    toVisit.Push(Position{
			x: p.x - 1,
			y: p.y,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    }
	case 'L':
	    if p.y > p.origin.y {
		if p.x < len(lines[p.y]) - 1 &&  isOneOf(lines[p.y][p.x + 1], "-J7") {
		    toVisit.Push(Position{
			x: p.x + 1,
			y: p.y,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    } else {
		if p.y > 0 && isOneOf(lines[p.y - 1][p.x], "|7F") {
		    toVisit.Push(Position{
			x: p.x,
			y: p.y - 1,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    }
	case 'J':
	    if p.x > p.origin.x {
		if p.y > 0 &&  isOneOf(lines[p.y - 1][p.x], "|7F") {
		    toVisit.Push(Position{
			x: p.x,
			y: p.y - 1,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    } else {
		if p.x > 0 && isOneOf(lines[p.y][p.x - 1], "-FL") {
		    toVisit.Push(Position{
			x: p.x - 1,
			y: p.y,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    }
	case '7':
	    if p.y < p.origin.y {
		if p.x > 0 && isOneOf(lines[p.y][p.x - 1], "-FL") {
		    toVisit.Push(Position{
			x: p.x - 1,
			y: p.y,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    } else {
		if p.y < len(lines) - 1 &&  isOneOf(lines[p.y + 1][p.x], "|LJ") {
		    toVisit.Push(Position{
			x: p.x,
			y: p.y + 1,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    }
	case 'F':
	    if p.y < p.origin.y {
		if p.x < len(lines[p.y]) - 1 &&  isOneOf(lines[p.y][p.x + 1], "-J7") {
		    toVisit.Push(Position{
			x: p.x + 1,
			y: p.y,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    } else {
		if p.y < len(lines) - 1 &&  isOneOf(lines[p.y + 1][p.x], "|JL") {
		    toVisit.Push(Position{
			x: p.x,
			y: p.y + 1,
			distanceFromS: p.distanceFromS + 1,
			origin: &p,
		    })
		}
	    }
	}
    }

    log.Fatalf("Error, shouldn't get here")
}

func isOneOf(c byte, target string) bool {
    for _, t := range target {
	if c == byte(t) || c == 'S' {
	    return true
	}
    }

    return false
}
