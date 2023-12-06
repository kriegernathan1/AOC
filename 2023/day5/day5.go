package day5

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type seedBounds struct {
	lower int
	upper int
}

func getSeedBounds(line string) []seedBounds {
	bounds := []seedBounds{}

	seedInfo := strings.Fields(strings.Split(line, ":")[1])
	for i := 0; i < len(seedInfo); i += 2 {
		rangeStart, err := strconv.Atoi(seedInfo[i])
		check(err)

		range_, err := strconv.Atoi(seedInfo[i+1])
		check(err)

		bounds = append(bounds, seedBounds{lower: rangeStart, upper: rangeStart + range_})
	}

	return bounds
}

func getSeeds(line string) []string {
	return strings.Fields(strings.Split(line, ":")[1])
}

func day5Part1(inputPath string) {
	fd, error := os.Open(inputPath)
	check(error)
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Scan()
	intermediateValues := getSeeds(string(scanner.Text()))

	for scanner.Scan() { // for each map
		line := string(scanner.Text())

		if strings.Contains(line, ":") || line == "" {
			continue
		}

		// read each line of map
		mapValues := []string{}
		for {
			line := string(scanner.Text())
			if line == "" {
				break
			}

			mapValues = append(mapValues, strings.Fields(line)...)
			scanner.Scan()
		}

		processedIndicies := make(map[int]bool)
		for i := 0; i < len(mapValues); i += 3 {
			// map values
			dstRangeStart, err := strconv.Atoi(mapValues[i])
			check(err)

			srcRangeStart, err := strconv.Atoi(mapValues[i+1])
			check(err)

			range_, err := strconv.Atoi(mapValues[i+2])
			check(err)

			for j := 0; j < len(intermediateValues); j++ {
				valueToBeMapped, err := strconv.Atoi(intermediateValues[j])
				check(err)

				if processedIndicies[j] {
					continue
				}

				if valueToBeMapped >= srcRangeStart && valueToBeMapped < srcRangeStart+range_ {
					intermediateValues[j] = strconv.Itoa(dstRangeStart + (valueToBeMapped - srcRangeStart))
					processedIndicies[j] = true
				}
			}

		}
	}

	location := math.MaxInt
	for i := 0; i < len(intermediateValues); i++ {
		currVal, _ := strconv.Atoi(intermediateValues[i])
		if currVal < location {
			location = currVal
		}
	}

	fmt.Printf("Part 1: %v\n", location)
}

func day5Part2(inputPath string) {
	fd, error := os.Open(inputPath)
	check(error)
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	scanner.Scan()
	bounds := getSeedBounds(string(scanner.Text()))

	for scanner.Scan() { // for each map
		line := string(scanner.Text())

		if strings.Contains(line, ":") || line == "" {
			continue
		}

		// read each line of map
		mapValues := []string{}
		for {
			line := string(scanner.Text())
			if line == "" {
				break
			}

			mapValues = append(mapValues, strings.Fields(line)...)
			scanner.Scan()
		}

		for i := 0; i < len(bounds); i++ {
			currBound := bounds[i]

			for j := 0; j < len(mapValues); j += 3 {
				dstStart, _ := strconv.Atoi(mapValues[j])
				srcStart, _ := strconv.Atoi(mapValues[j+1])
				range_, _ := strconv.Atoi(mapValues[j+2])

				mapUpperBounds := srcStart + range_
				isOverlap := (currBound.lower >= srcStart && currBound.lower <= mapUpperBounds) || (currBound.upper <= mapUpperBounds && currBound.upper >= srcStart)
				if isOverlap {
					newLowerBound := dstStart + (currBound.lower - srcStart)

					if currBound.upper > mapUpperBounds && currBound.lower > srcStart {
						processedBound := seedBounds{lower: newLowerBound, upper: dstStart + range_ - 1}
						bounds[i] = processedBound

						unprocessedBound := seedBounds{lower: mapUpperBounds + 1, upper: currBound.upper}
						bounds = append(bounds, unprocessedBound)
					} else if currBound.lower < srcStart && currBound.upper < mapUpperBounds {
						processedBound := seedBounds{lower: dstStart, upper: dstStart + (currBound.upper - srcStart)}
						bounds[i] = processedBound

						unprocessedBound := seedBounds{lower: currBound.lower, upper: srcStart - 1}
						bounds = append(bounds, unprocessedBound)
					} else if currBound.lower < srcStart && currBound.upper > mapUpperBounds {
						processedBound := seedBounds{lower: dstStart, upper: dstStart + range_ - 1}
						bounds[i] = processedBound

						leftUnprocessedBound := seedBounds{lower: currBound.lower, upper: srcStart - 1}
						rightUnprocessedBound := seedBounds{lower: srcStart + range_ + 1, upper: currBound.upper}
						bounds = append(bounds, leftUnprocessedBound, rightUnprocessedBound)
					} else { // map covers entire range of seeds
						bounds[i] = seedBounds{lower: newLowerBound, upper: dstStart + (currBound.upper - srcStart)}
					}
				}
			}
		}
	}

	location := math.MaxInt
	for i := 0; i < len(bounds); i++ {
		currVal := bounds[i].lower
		if currVal < location {
			location = currVal
		}
	}

	fmt.Printf("Part 2: %v\n", location) // It's off by one. Not sure why
}

func Day5(inputPath string) {
	day5Part1(inputPath)
	start := time.Now()
	day5Part2(inputPath)
	elapsed := time.Since(start)
	fmt.Printf("Part 2 took %v\n", elapsed)
}
