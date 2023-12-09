package main

import (
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"aoc/day7"
	"aoc/day8"
	"fmt"
	"os"
)

var days = map[string]func(path string){
	"2": day2.Day2,
	"3": day3.Day3,
	"4": day4.Day4,
	"5": day5.Day5,
	"6": day6.Day6,
	"7": day7.Day7,
	"8": day8.Day8,
}

func main() {
	args := os.Args[1:] // don't include path

	if len(args) != 2 {
		panic("Incorrect number of arguments. Example go run main.go 2 day2/input.txt")
	}

	day := args[0]
	path := args[1]
	if days[day] != nil {
		days[day](path)
	} else {
		panic(fmt.Sprintf("Unable to find entry point for day %v.", day))
	}
}
