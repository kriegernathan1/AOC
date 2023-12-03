package main

import (
	"fmt"
	"os"
	"strconv"
)

var NEWLINE int = 10
var PERIOD int = 46
var STAR int = 42

func isSymbol(data byte) bool {
	return !isNumber(data) && data != byte(PERIOD)
}

func isGear(data byte) bool {
	return data == byte(STAR)
}

func isNumber(data byte) bool {
	return data >= 48 && data <= 57 // UTF8
}

func isValidIndicies(row int, col int, maxRow int, maxCol int) bool {
	isNotNegative := row >= 0 && col >= 0
	isNotOutOfRange := row <= maxRow && col < maxCol

	return isNotNegative && isNotOutOfRange
}

func getNumberFromIndicesAndRemoveFromSourceArray(row int, col int, schematic [][]byte) int {
	left := col
	right := col
	MIN_COL := 0
	MAX_COL := len(schematic[row]) - 1
	// find left side of number
	for {
		if left < MIN_COL || !isNumber(schematic[row][left]) {
			left = left + 1
			break
		}
		left--
	}

	// go right until you reach non-number OR end of row
	for {
		if right > MAX_COL || !isNumber(schematic[row][right]) {
			break
		}
		right++
	}

	// create slice with left and right
	numberByteArr := schematic[row][left:right]

	// convert to int
	number, err := strconv.Atoi(string(numberByteArr))
	if err != nil {
		panic("Error converting string to int" + err.Error())
	}

	// remove number from source array
	for left < right && left < len(schematic[row]) {
		schematic[row][left] = byte(PERIOD)
		left++
	}

	return number
}

func part1(schematic [][]byte) {
	partNumbers := []int{}

	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[row]); col++ {
			if isSymbol(schematic[row][col]) {
				ADJACENT_INDICES := [][]int{
					{row - 1, col - 1}, // upper left
					{row - 1, col},     // upper
					{row - 1, col + 1}, // upper right
					{row, col - 1},     // left
					{row, col + 1},     // right
					{row + 1, col - 1}, // lower left
					{row + 1, col},     // lower
					{row + 1, col + 1}, // lower right
				}

				for i := 0; i < len(ADJACENT_INDICES); i++ {
					row := ADJACENT_INDICES[i][0]
					col := ADJACENT_INDICES[i][1]

					if isValidIndicies(row, col, len(schematic), len(schematic[row])) && isNumber(schematic[row][col]) {
						partNumbers = append(partNumbers, getNumberFromIndicesAndRemoveFromSourceArray(row, col, schematic))
					}
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(partNumbers); i++ {
		sum += partNumbers[i]
	}

	fmt.Println(sum)
}

func part2(schematic [][]byte) {
	gearRatios := []int{}

	for row := 0; row < len(schematic); row++ {
		for col := 0; col < len(schematic[row]); col++ {
			if isGear(schematic[row][col]) {
				partNumbers := []int{}
				ADJACENT_INDICES := [][]int{
					{row - 1, col - 1}, // upper left
					{row - 1, col},     // upper
					{row - 1, col + 1}, // upper right
					{row, col - 1},     // left
					{row, col + 1},     // right
					{row + 1, col - 1}, // lower left
					{row + 1, col},     // lower
					{row + 1, col + 1}, // lower right
				}

				for i := 0; i < len(ADJACENT_INDICES); i++ {
					row := ADJACENT_INDICES[i][0]
					col := ADJACENT_INDICES[i][1]

					if isValidIndicies(row, col, len(schematic), len(schematic[row])) && isNumber(schematic[row][col]) {
						partNumbers = append(partNumbers, getNumberFromIndicesAndRemoveFromSourceArray(row, col, schematic))
					}
				}

				if len(partNumbers) == 2 {
					gearRatios = append(gearRatios, partNumbers[0]*partNumbers[1])
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(gearRatios); i++ {
		sum += gearRatios[i]
	}

	fmt.Println(sum)
}

func main() {
	args := os.Args[1:] // don't include path

	if len(args) != 1 {
		panic("Two many arguments. Only 1 argument (input path) is allowed")
	}

	path := args[0]
	data, err := os.ReadFile(path)
	if err != nil {
		panic("unable to open input file.")
	}

	schematic := [][]byte{}

	left := 0
	for right := 0; right < len(data); right++ {
		if data[right] == byte(NEWLINE) || right+1 == len(data) {
			schematic = append(schematic, data[left:right])
			left = right + 1
		}
	}

	// Note: Only will be correct if run with one commented; Both mutate the schematic array.
	// part1(schematic)
	part2(schematic)
}
