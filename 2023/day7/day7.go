package day7

import (
	"aoc/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards    string
	bet      int
	strength int
	rank     int
}

/*
Five of a kind: [5 0 0 0 0] => 7
Four of a kind: [4 1 0 0 0] => 6
Full house: [3 2 0 0 0] => 5
Three of a kind: [3 1 1 0 0] => 4
Two Pair: [2 2 1 0 0] => 3
One Pair: [2 1 1 1 0] => 2
High card: [1 1 1 1 1] => 1
*/
var cardFreqToHandStrength = map[[5]int]int{
	{5, 0, 0, 0, 0}: 7,
	{4, 1, 0, 0, 0}: 6,
	{3, 2, 0, 0, 0}: 5,
	{3, 1, 1, 0, 0}: 4,
	{2, 2, 1, 0, 0}: 3,
	{2, 1, 1, 1, 0}: 2,
	{1, 1, 1, 1, 1}: 1,
}

// 65 75 81 74 84 57 56 55 54 53 52 51 50
// AKQJT98765432
var cardToValue = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var part = 1

var wildCard byte = 'J'

func getHandFromLine(line string) hand {
	leftAndRight := strings.Split(line, " ")

	cards := leftAndRight[0]
	bet, err := strconv.Atoi(leftAndRight[1])
	util.Check(err)

	h := hand{cards: cards, bet: bet}
	h.strength = getHandStrength(h)
	return h
}

func getHandStrength(h hand) int {
	cardFreq := make(map[byte]int)

	for i := 0; i < len(h.cards); i++ {
		_, ok := cardFreq[h.cards[i]]

		if !ok {
			cardFreq[h.cards[i]] = 1
		} else {
			cardFreq[h.cards[i]]++
		}
	}

	var frequencies [5]int

	i := 0
	for _, v := range cardFreq {
		frequencies[i] = v
		i++
	}

	sort.Slice(frequencies[:], func(i, j int) bool {
		return frequencies[i] > frequencies[j]
	})

	if part == 2 {
		if cardFreq[wildCard] > 0 {
			isWildCardHighestFreq := cardFreq[wildCard] == frequencies[0]
			targetToRemove := cardFreq[wildCard]

			if isWildCardHighestFreq && frequencies[0] == frequencies[1] {
				frequencies[0] += cardFreq[wildCard]
			} else if isWildCardHighestFreq {
				secondHighestFreq := frequencies[1]
				frequencies[0] += secondHighestFreq
				targetToRemove = secondHighestFreq
			} else {
				frequencies[0] += cardFreq[wildCard]
			}

			// remove wild card freq as it has been used
			for i := 0; i < len(frequencies); i++ {
				if frequencies[i] == targetToRemove {
					frequencies[i] = 0
					break
				}
			}
		}
		sort.Slice(frequencies[:], func(i, j int) bool {
			return frequencies[i] > frequencies[j]
		})
	}

	strength := cardFreqToHandStrength[frequencies]
	if strength == 0 {
		fmt.Println("Impossible card frequencies", frequencies, h.cards)
	}

	return strength
}

// if h1 < h2
func compareHands(h1 hand, h2 hand) bool {
	if h1.strength > h2.strength {
		return false
	} else if h1.strength < h2.strength {
		return true
	}

	if len(h1.cards) != 5 {
		panic("Assumption about cards length is wrong")
	}

	for i := 0; i < 5; i++ {
		h1CardVal := cardToValue[h1.cards[i]]
		h2CardVal := cardToValue[h2.cards[i]]

		if h1CardVal < h2CardVal {
			return true
		} else if h1CardVal > h2CardVal {
			return false
		}
	}

	return false
}

func day7(inputPath string) {
	if part == 2 {
		cardToValue[wildCard] = 0
	}

	scanner := util.GetScanner(inputPath)

	h := []hand{}
	for scanner.Scan() {
		h = append(h, getHandFromLine(scanner.Text()))
	}

	sort.Slice(h, func(i, j int) bool {
		return compareHands(h[i], h[j])
	})

	totalWinnings := 0
	for rank, hand := range h {
		totalWinnings += hand.bet * (rank + 1)
	}

	fmt.Printf("Part %v: %v\n", part, totalWinnings)
}

func Day7(inputPath string) {
	part = 1
	day7(inputPath)

	part = 2
	day7(inputPath)
}
