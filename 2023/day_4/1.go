package main

import (
	"fmt"
	"os"
	"strings"
)

var NEWLINE int = 10

func processCardsPart1(card string) int {
	winningNumbersMap := make(map[string]bool)

	leftAndRight := strings.Split(card, " | ")
	cardNumbers := strings.Fields(leftAndRight[1])
	winningNumbers := strings.Fields(leftAndRight[0])[2:]

	for i := 0; i < len(winningNumbers); i++ {
		winningNumbersMap[winningNumbers[i]] = true
	}

	points := 0
	for i := 0; i < len(cardNumbers); i++ {
		if winningNumbersMap[cardNumbers[i]] {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return points
}

func getCardMatches(card string) int {
	winningNumbersMap := make(map[string]bool)

	leftAndRight := strings.Split(card, " | ")
	cardNumbers := strings.Fields(leftAndRight[1])
	// cardId := strings.Fields(leftAndRight[0])[1]
	// cardId = cardId[:len(cardId)-1]
	winningNumbers := strings.Fields(leftAndRight[0])[2:]

	for i := 0; i < len(winningNumbers); i++ {
		winningNumbersMap[winningNumbers[i]] = true
	}

	matches := 0
	for i := 0; i < len(cardNumbers); i++ {
		if winningNumbersMap[cardNumbers[i]] {
			matches += 1
		}
	}

	return matches
}

func part1(cards []string) {
	sum := 0
	for i := 0; i < len(cards); i++ {
		sum += processCardsPart1(cards[i])
	}

	fmt.Printf("part 1: %v\n", sum)
}

func part2(cards []string) {
	cardToNumScrachers := make(map[int]int)

	numScratchers := len(cards)
	for i := len(cards) - 2; i >= 0; i-- {
		matches := getCardMatches(cards[i])
		cardID := i + 1

		scratchersGained := matches
		for i := cardID + 1; i <= cardID+matches; i++ {
			scratchersGained += cardToNumScrachers[i]
		}
		cardToNumScrachers[cardID] = scratchersGained
		numScratchers += scratchersGained
	}

	fmt.Println(numScratchers)
}

func main() {
	args := os.Args[1:] // don't include path

	if len(args) != 1 {
		panic("Incorrect number of arguments. Only 1 argument (input path) is allowed")
	}

	path := args[0]
	data, err := os.ReadFile(path)
	if err != nil {
		panic("unable to open input file.")
	}

	cards := []string{}

	left := 0
	for right := 0; right < len(data); right++ {
		if data[right] == byte(NEWLINE) || right+1 == len(data) {
			cards = append(cards, string(data[left:right]))
			left = right + 1
		}
	}

	// part1(cards)
	part2(cards)
}
