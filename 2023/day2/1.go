package day2

import (
	"fmt"
	"os"
	"strconv"
)

var NEWLINE int = 10
var DELIMETER int = 59 // semi-colon
var COMMA int = 44
var SPACE int = 32
var SEMICOLON int = 59
var MAX_COLOR_MAP = make(map[string]int)
var colors = []string{"red", "green", "blue"}

type game struct {
	ID        int
	handsfuls []map[string]int
}

func Day2(inputFile string) {
	solvePartOne(inputFile)
	solvePartTwo(inputFile)
}

func solvePartOne(inputFile string) {
	MAX_COLOR_MAP["red"] = 12
	MAX_COLOR_MAP["green"] = 13
	MAX_COLOR_MAP["blue"] = 14

	games := parseGameInput(inputFile)
	possibleGames := []game{}

	for i := 0; i < len(games); i++ {
		if isGamePossible(games[i]) {
			possibleGames = append(possibleGames, games[i])
		}
	}

	sum := 0
	for i := 0; i < len(possibleGames); i++ {
		sum += possibleGames[i].ID
	}

	fmt.Printf("part 1 answer: %v\n", sum) // 2285
}

func solvePartTwo(inputFile string) {
	games := parseGameInput(inputFile)

	sum := 0
	for i := 0; i < len(games); i++ {
		sum += getPowerOfCubeSet(games[i])
	}

	fmt.Printf("part 2 answer: %v\n", sum) // 77021
}

func getPowerOfCubeSet(g game) int {
	gameMaxColorMap := make(map[string]int)

	for i := 0; i < len(colors); i++ {
		gameMaxColorMap[colors[i]] = 0
	}

	for i := 0; i < len(g.handsfuls); i++ {
		for j := 0; j < len(colors); j++ {
			handful := g.handsfuls[i]
			color := colors[j]

			// handle each color
			if handful[color] > gameMaxColorMap[color] {
				gameMaxColorMap[color] = handful[color]
			}
		}
	}

	power := 1
	for i := 0; i < len(colors); i++ {
		power *= gameMaxColorMap[colors[i]]
	}

	return power
}

// # color, # color, # color[;][<newline>]
func parseHandful(_handful []byte) map[string]int {
	colorFreqMap := make(map[string]int)

	left := 0
	right := 0
	freq := 0
	color := ""

	// find and populate map (color => freq) until we reach end
out:
	for {
		// get number
		for {
			if _handful[right] == byte(SPACE) {
				freq, _ = strconv.Atoi(string(_handful[left:right]))
				left = right + 1
				right = left
				break
			}
			right++
		}

		// get color
		for {
			if _handful[right] == byte(COMMA) || _handful[right] == byte(NEWLINE) || _handful[right] == byte(SEMICOLON) || right+1 == len(_handful) {
				color = string(_handful[left:right])
				left = right + 2

				if _handful[right] == byte(COMMA) {
					right = left
					break
				} else {
					colorFreqMap[color] = freq
					break out
				}
			}

			right++
		}

		colorFreqMap[color] = freq
	}

	return colorFreqMap
}

func parseGame(_game []byte) game {
	// handfull Form = Game <ID>: # <color>, # <color>, # <color>
	// game = handful;handful;handful<NEWLINE>
	parsedGame := game{}

	// get ID
	left := 5
	right := 5
	for {
		if _game[right] == byte(SPACE) {
			break
		}
		right++
	}

	parsedGame.ID, _ = strconv.Atoi(string(_game[left : right-1]))

	left = right + 1 // start on first number of handful
	for i := right; i < len(_game); i++ {
		if _game[i] == byte(SEMICOLON) || _game[i] == byte(NEWLINE) || i+1 == len(_game) {
			parsedGame.handsfuls = append(parsedGame.handsfuls, parseHandful(_game[left:i+1]))
			left = i + 2
		}
	}

	return parsedGame
}

func parseGameInput(inputFile string) []game {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	games := []game{}
	begin := 0
	for i := 0; i < len(data); i++ {
		if data[i] == byte(NEWLINE) || i+1 == len(data) {
			games = append(games, parseGame(data[begin:i+1]))
			begin = i + 1
		}
	}

	return games
}

func isGamePossible(g game) bool {
	gameMaxColorMap := make(map[string]int)

	for i := 0; i < len(colors); i++ {
		gameMaxColorMap[colors[i]] = 0
	}

	for i := 0; i < len(g.handsfuls); i++ {
		for j := 0; j < len(colors); j++ {
			handful := g.handsfuls[i]
			color := colors[j]

			// handle each color
			if handful[color] > gameMaxColorMap[color] {
				gameMaxColorMap[color] = handful[color]
			}
		}
	}

	for i := 0; i < len(colors); i++ {
		color := colors[i]
		gameMaxForColor := gameMaxColorMap[color]

		if gameMaxForColor > MAX_COLOR_MAP[color] {
			return false
		}
	}

	return true
}
