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

func (rs *RangeStruct) getDestination2(x int) int {
	diff := x - rs.SourceStart
	return rs.Destination + diff
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

func parseSeedsPart2(line string) []RangeStruct {
	seedsRanges := make([]int, 0)

	line, _ = strings.CutPrefix(line, "seeds: ")
	split := strings.Split(line, " ")

	for _, num := range split {
		converted, err := strconv.Atoi(num)
		if err != nil {
			panic(2)
		}
		seedsRanges = append(seedsRanges, converted)
	}

	ranges := make([]RangeStruct, 0)

	for i := 0; i < len(seedsRanges); i += 2 {
		start := seedsRanges[i]
		end := start + seedsRanges[i+1]
		ranges = append(ranges, RangeStruct{SourceStart: start, SourceEnd: end})
	}

	return ranges
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

func part2(filename string) int {
	lines := readFile(filename)
	seedsRanges := parseSeedsPart2(lines[0])
	parsedRanges := parseRanges(lines[3:])

	for i := 0; i < len(parsedRanges); i++ {
		seedsRanges = applyRange(seedsRanges, parsedRanges[i])
	}

	mins := make([]int, 0)

	for _, seedRange := range seedsRanges {
		if seedRange.SourceStart == 0 {
			continue
		}
		mins = append(mins, seedRange.SourceStart)
	}
	return slices.Min(mins)
}

func applyRange(seedsRanges []RangeStruct, mapRanges []RangeStruct) []RangeStruct {
	newSeedRanges := make([]RangeStruct, 0)

	for _, seedRange := range seedsRanges {
		mapped := false
		for _, mapRange := range mapRanges {
			if seedRange.SourceStart >= mapRange.SourceStart && seedRange.SourceEnd <= mapRange.SourceEnd {
				newSeedRanges = myAppend(newSeedRanges, RangeStruct{SourceStart: mapRange.getDestination2(seedRange.SourceStart), SourceEnd: mapRange.getDestination2(seedRange.SourceEnd)})
				mapped = true
				break
			}

			if seedRange.SourceStart <= mapRange.SourceStart && seedRange.SourceEnd >= mapRange.SourceEnd {
				beforeCommonRange := RangeStruct{SourceStart: seedRange.SourceStart, SourceEnd: mapRange.SourceStart}
				commonRange := RangeStruct{SourceStart: mapRange.getDestination2(mapRange.SourceStart), SourceEnd: mapRange.getDestination2(mapRange.SourceEnd)}
				afterCommonRange := RangeStruct{SourceStart: mapRange.SourceEnd, SourceEnd: seedRange.SourceEnd}
				newSeedRanges = myAppend(newSeedRanges, beforeCommonRange)
				newSeedRanges = myAppend(newSeedRanges, commonRange)
				newSeedRanges = myAppend(newSeedRanges, afterCommonRange)
				mapped = true
				break
			}

			if seedRange.SourceStart < mapRange.SourceStart && seedRange.SourceEnd < mapRange.SourceEnd && seedRange.SourceEnd >= mapRange.SourceStart {
				beforeCommonRange := RangeStruct{SourceStart: seedRange.SourceStart, SourceEnd: mapRange.SourceStart}
				commonRange := RangeStruct{SourceStart: mapRange.getDestination2(mapRange.SourceStart), SourceEnd: mapRange.getDestination2(seedRange.SourceEnd)}
				newSeedRanges = myAppend(newSeedRanges, beforeCommonRange)
				newSeedRanges = myAppend(newSeedRanges, commonRange)
				mapped = true
				break
			}

			if seedRange.SourceStart >= mapRange.SourceStart && seedRange.SourceEnd >= mapRange.SourceStart && mapRange.SourceEnd >= seedRange.SourceStart {
				commonRange := RangeStruct{SourceStart: mapRange.getDestination2(seedRange.SourceStart), SourceEnd: mapRange.getDestination2(mapRange.SourceEnd)}
				afterCommonRange := RangeStruct{SourceStart: mapRange.SourceEnd, SourceEnd: seedRange.SourceEnd}
				newSeedRanges = myAppend(newSeedRanges, commonRange)
				newSeedRanges = myAppend(newSeedRanges, afterCommonRange)
				mapped = true
				break
			}
		}
			if !mapped {
				newSeedRanges = append(newSeedRanges, seedRange)
			}

	}

	return newSeedRanges
}

func myAppend(destination []RangeStruct, source RangeStruct) []RangeStruct {
	if source.SourceStart == source.SourceEnd {
		return append(destination, RangeStruct{SourceStart: source.SourceStart, SourceEnd: source.SourceEnd + 1})
	}
	return append(destination, source)
}

func main() {
	fileName := os.Args[1]
	//fmt.Println(part1(fileName))
	fmt.Println(part2(fileName))
}
