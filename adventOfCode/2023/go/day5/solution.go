package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type RangeStruct struct {
	SourceStart int
	SourceEnd   int
	Destination int
	Range       int
}

func (rs *RangeStruct) getDestination(x int) (bool, int) {
	if x >= rs.SourceStart && x <= rs.SourceEnd {
		diff := x - rs.SourceStart
		return true, rs.Destination + diff
	}
	return false, x
}

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(1)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	lines := make([]string, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func parseSeeds(line string) []int {
	seeds := make([]int, 0)

	line, _ = strings.CutPrefix(line, "seeds: ")
	split := strings.Split(line, " ")

	for _, num := range split {
		converted, err := strconv.Atoi(num)
		if err != nil {
			panic(2)
		}
		seeds = append(seeds, converted)
	}

	return seeds
}

func parseRanges(lines []string) map[int][]RangeStruct {
	ranges := make(map[int][]RangeStruct)
	mapPointer := 0
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if strings.Contains(line, ":") {
			mapPointer++
			continue
		}
		split := strings.Split(line, " ")
		rangeStruct := convertNumbers(split)
		ranges[mapPointer] = append(ranges[mapPointer], rangeStruct)
	}
	return ranges
}

func convertNumbers(numbersAsStrings []string) RangeStruct {
	var err error
	var source int
	var destination int
	var length int

	destination, err = strconv.Atoi(numbersAsStrings[0])
	if err != nil {
		panic(3)
	}

	source, err = strconv.Atoi(numbersAsStrings[1])
	if err != nil {
		panic(3)
	}

	length, err = strconv.Atoi(numbersAsStrings[2])
	if err != nil {
		panic(3)
	}
	return RangeStruct{SourceStart: source, Destination: destination, Range: length, SourceEnd: source + length}
}

func part1(filename string) int {
	lines := readFile(filename)
	seeds := parseSeeds(lines[0])
	fmt.Println(seeds)
	parsedRanges := parseRanges(lines[3:])

	for i := 0; i < len(parsedRanges); i++ {
		var wg sync.WaitGroup
		for j, seed := range seeds {
			wg.Add(1)
			go func(oneSeed int, seedNo int) {
				defer wg.Done()
				for k := range parsedRanges[i] {
					inRange, newVal := parsedRanges[i][k].getDestination(oneSeed)
					if inRange {
						seeds[seedNo] = newVal
					}
				}

			}(seed, j)
		}
		wg.Wait()
	}

	return slices.Min(seeds)
}

func main() {
	fileName := os.Args[1]
	fmt.Println(part1(fileName))
}
