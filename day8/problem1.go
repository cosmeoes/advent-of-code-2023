package day8

import (
	"log"
	"strings"

	"cosme.dev/aoc2023/helpers"
)

type Node struct {
    name string
    left *Node
    right *Node
}

func Problem1() {
    lines, err := helpers.ReadFileLines("day8/input.txt")
    if err != nil {
	log.Fatalf("Could not read file lines: %v", err)
    }

    instructions := []byte(lines[0])
    var start *Node
    nodes := make(map[string]*Node);
    for i := 2; i < len(lines); i++ {
	nodeString := strings.Split(lines[i], " = ")
	name := nodeString[0]
	node, ok := nodes[name]
	if !ok {
	    node = &Node{
		name: name,
	    }
	    nodes[name] = node
	}

	instructionsString := strings.Trim(nodeString[1], "() ")
	branches := strings.Split(instructionsString, ", ")
	left, ok := nodes[branches[0]]
	if !ok {
	    left = &Node{
		name: branches[0],
	    }
	    nodes[branches[0]] = left
	}

	right, ok := nodes[branches[1]]
	if !ok {
	    right = &Node{
		name: branches[1],
	    }
	    nodes[branches[1]] = right
	}

	node.left = left
	node.right = right

	if node.name == "AAA" {
	    start = node
	}

    }

    current := start
    steps := 0
    for current.name != "ZZZ" {
	instruction := instructions[steps % len(instructions)]
	steps++
	switch instruction {
	case 'R':
	    current = current.right
	case 'L':
	    current = current.left
	}
    }

    log.Println(steps)
}
