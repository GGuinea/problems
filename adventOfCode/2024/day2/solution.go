package day1

import (
	"math"
	"os"
	"strconv"
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

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, " ")

		if isGood(nums) {
			res++
		}
	}

	return res
}

func isGood(arr []string) bool {
	prev := -1
	var direction int
	for i, num := range arr {
		asInt, err := strconv.Atoi(num)
		check(err)
		if i == 0 {
			prev = asInt
			continue
		}
		if i == 1 {
			if prev-asInt < 0 {
				direction = 1
			} else if prev-asInt > 0 {
				direction = -1
			}
		}
		diffAbs := math.Abs(float64(asInt - prev))
		if diffAbs > 3 || diffAbs == 0 {
			return false

		}

		if (direction == -1 && asInt-prev > 0) || (direction == 1 && asInt-prev < 0) {
			return false
		}

		prev = asInt
	}
	return true
}

func handlePart2(input string) int {
	lines := strings.Split(input, "\n")
	res := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		nums := strings.Split(line, " ")
		var good bool
		for i := range len(nums) {
			newNums := make([]string, 0)
			newNums = append(newNums, nums[:i]...)
			newNums = append(newNums, nums[i+1:]...)
			if isGood(newNums) {
				good = true
			}
		}
		if good {
			res++
		}
	}

	return res
}
