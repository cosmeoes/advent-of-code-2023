package day8

import (
	"log"
	"strings"

	"cosme.dev/aoc2023/helpers"
)


// Lil bro tried to do cpu cache optimization NAHHHH
var nodesAOS []*NodeAOS

type NodeAOS struct {
    name string
    left int
    right int
}

func Problem2() {
    lines, err := helpers.ReadFileLines("day8/input.txt")
    if err != nil {
	log.Fatalf("Could not read file lines: %v", err)
    }

    instructions := []byte(lines[0])

    ins := make([]string, len(instructions))
    for i, in := range instructions {
	ins[i] = string(in)
    }

    var starts []int
    nodesMap := make(map[string]int);
    nodesAOS = make([]*NodeAOS, len(lines) + 2)
    pos := 0
    for i := 2; i < len(lines); i++ {
	nodeString := strings.Split(lines[i], " = ")
	name := nodeString[0]
	index, ok := nodesMap[name]
	if !ok {
	    nodesAOS[pos] = &NodeAOS{
		name: name,
	    }
	    nodesMap[name] = pos
	    index = pos
	    pos++
	}
	node := nodesAOS[index]

	instructionsString := strings.Trim(nodeString[1], "() ")
	branches := strings.Split(instructionsString, ", ")
	leftIndex, ok := nodesMap[branches[0]]
	if !ok {
	    nodesAOS[pos] = &NodeAOS{
		name: branches[0],
	    }
	    nodesMap[branches[0]] = pos
	    leftIndex = pos
	    pos++
	}

	rightIndex, ok := nodesMap[branches[1]]
	if !ok {
	    nodesAOS[pos] = &NodeAOS{
		name: branches[1],
	    }
	    nodesMap[branches[1]] = pos
	    rightIndex = pos
	    pos++
	}

	node.left = leftIndex
	node.right = rightIndex

	if strings.HasSuffix(node.name, "A") {
	    starts = append(starts, index)
	}
    }

    zCounts := make([]int, len(starts))
    for i, start := range starts {
	steps := 0
	current := start
	for nodesAOS[current].name[2] != 'Z'  {
	    instruction := instructions[steps % len(instructions)]
	    current = move(current, instruction)
	    steps++
	}

	zCounts[i] = steps
    }

    log.Println(lcmAll(zCounts[0], zCounts[1:]...))
}

func done(currents []int) bool {
    for _, idx := range currents {
	currentNode := nodesAOS[idx]
	if currentNode.name[2] != 'Z' {
	    return false
	}
    }

    return true
}

func move(current int, instruction byte) int {
    var target int
    node := nodesAOS[current]

    switch instruction {
    case 'R':
	target = node.right
    case 'L':
	target = node.left
    default:
    }

    return target
}

// gcd,lcm,lcmAll copied from:
// https://github.com/torbensky/advent-of-code-2023/blob/main/day08/main.go#L70C1-L88C2
func gcd(a, b int) int {
    if b == 0 {
	return a
    }
    return gcd(b, a%b)
}
func lcm(a, b int) int {
    return a / gcd(a, b) * b
}

func lcmAll(a int, bs ...int) int {
    result := a
    for _, b := range bs {
	result = lcm(result, b)
    }

    return result
}
