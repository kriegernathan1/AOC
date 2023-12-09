package day8

import (
	"aoc/util"
	"fmt"
	"strings"
)

type Node struct {
	identifier    string
	adjacentNodes map[string]string // L and R are keys
}

func Day8(inputPath string) {
	const startingPoint = "AAA"
	const endingPoint = "ZZZ"

	nodeMap := make(map[string]Node)
	scanner := util.GetScanner(inputPath)

	scanner.Scan()
	instructions := strings.Split(scanner.Text(), "")
	scanner.Scan()

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()

		leftAndRight := strings.Split(line, "=")
		identifier := strings.Fields(leftAndRight[0])[0]
		adjacentNodes := strings.Fields(strings.Join(strings.Split(strings.TrimSpace(leftAndRight[1])[1:len(leftAndRight[1])-2], ","), "")) // Not proud of myself :(

		nodeMap[identifier] = Node{identifier: identifier, adjacentNodes: map[string]string{"L": adjacentNodes[0], "R": adjacentNodes[1]}}
	}

	currentNode := nodeMap[startingPoint]
	numSteps := 0
	for {
		currentInstruction := instructions[numSteps%len(instructions)]
		nextNode := nodeMap[currentNode.adjacentNodes[currentInstruction]]

		currentNode = nextNode

		numSteps++

		if nextNode.identifier == endingPoint {
			break
		}
	}

	fmt.Println(numSteps)
}
