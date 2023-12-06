package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getTimeOrDistance(line string, part int) []int {
	if part == 1 {
		timeOrDistanceStr := strings.Fields(strings.Split(line, ":")[1])

		timeOrDistanceInt := make([]int, len(timeOrDistanceStr))
		for i := 0; i < len(timeOrDistanceStr); i++ {
			timeOrDistanceInt[i], _ = strconv.Atoi(timeOrDistanceStr[i])
		}

		return timeOrDistanceInt
	} else {
		timeOrDistanceStr, _ := strconv.Atoi(strings.Join(strings.Fields(strings.Split(line, ":")[1]), ""))
		return []int{timeOrDistanceStr}
	}
}

func canWinWithPressOfLength(length int, totalTime int, neededDistance int) bool {
	timeRemaining := totalTime - length
	return timeRemaining*length > neededDistance
}

func day6(inputPath string, part int) {
	fd, err := os.Open(inputPath)
	check(err)

	scanner := bufio.NewScanner(fd)
	timeAndDistances := make([][]int, 2)

	for i := 0; scanner.Scan(); i++ {
		timeAndDistances[i] = getTimeOrDistance(scanner.Text(), part)
	}

	numberOfRaces := len(timeAndDistances[0])
	waysToWin := []int{}
	for i := 0; i < numberOfRaces; i++ {
		time := timeAndDistances[0][i]
		distance := timeAndDistances[1][i]

		numOfWaysToWin := 0
		for j := 1; j < time; j++ {
			if canWinWithPressOfLength(j, time, distance) {
				numOfWaysToWin++
			}
		}
		waysToWin = append(waysToWin, numOfWaysToWin)
	}

	totalWaysToWin := waysToWin[0]
	for i := 1; i < len(waysToWin); i++ {
		totalWaysToWin *= waysToWin[i]
	}

	fmt.Printf("Part %v: %v\n", part, totalWaysToWin)
}

func Day6(inputPath string) {
	day6(inputPath, 1)
	day6(inputPath, 2)
}
