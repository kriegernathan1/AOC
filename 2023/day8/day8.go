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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func getNumStepsFromStartingPoint(instructions []string, nodeMap map[string]Node, startingNode Node, shouldBreak func(n Node) bool) int {
	currentNode := startingNode
	numSteps := 0
	for {
		currentInstruction := instructions[numSteps%len(instructions)]
		nextNode := nodeMap[currentNode.adjacentNodes[currentInstruction]]

		currentNode = nextNode

		numSteps++

		if shouldBreak(nextNode) {
			break
		}
	}

	return numSteps
}

func part1(instructions []string, nodeMap map[string]Node) {
	fmt.Println(getNumStepsFromStartingPoint(instructions, nodeMap, nodeMap[startingPoint], func(n Node) bool {
		return n.identifier == endingPoint
	}))
}

func part2(instructions []string, nodeMap map[string]Node) {
	const targetNodeEndingLetter = 'Z'
	const targetStartLetter = 'A'
	currentNodes := []Node{}

	for k, v := range nodeMap {
		if k[2] == targetStartLetter {
			currentNodes = append(currentNodes, v)
		}
	}

	pathLengths := []int{}
	for _, v := range currentNodes {
		length := getNumStepsFromStartingPoint(instructions, nodeMap, v, func(n Node) bool {
			return n.identifier[2] == targetNodeEndingLetter
		})

		pathLengths = append(pathLengths, length)
	}

	fmt.Println(LCM(pathLengths[0], pathLengths[1], pathLengths[2:]...))
}

const startingPoint = "AAA"
const endingPoint = "ZZZ"

func Day8(inputPath string) {

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

	part1(instructions, nodeMap)
	part2(instructions, nodeMap)

}
