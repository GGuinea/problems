package day1

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) string {
	dat, err := os.ReadFile(filename)
	check(err)
	return string(dat)
}

func part1() int {
	input := readFile("input")
	return handlePart1(input)
}

func part2() int {
	input := readFile("input")
	return handlePart2(input)
}

func handlePart1(input string) int {
	lines := strings.Split(input, "\n")
	res := 0
	arrL := make([]int, 0)
	arrR := make([]int, 0)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var left, right int
		n, err := fmt.Sscanf(line, "%d   %d", &left, &right)
		check(err)
		if n != 2 {
			panic("Wrong number of read inputs")
		}

		arrL = append(arrL, left)
		arrR = append(arrR, right)
	}

	slices.Sort(arrL)
	slices.Sort(arrR)

	for i := range len(arrL) {
		res += int(math.Abs(float64(arrL[i] - arrR[i])))
	}

	return res
}

func handlePart2(input string) int {
	lines := strings.Split(input, "\n")
	res := 0
	arrL := make([]int, 0)
	occurrenceMap := make(map[int]int)
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		var left, right int
		n, err := fmt.Sscanf(line, "%d   %d", &left, &right)
		check(err)
		if n != 2 {
			panic("Wrong number of read inputs")
		}

		arrL = append(arrL, left)
		occurrenceMap[right]++
	}

	for _, val := range arrL {
		res += val * occurrenceMap[val]
	}

	return res
}
