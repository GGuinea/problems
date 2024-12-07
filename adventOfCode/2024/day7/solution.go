package day1

import (
	"fmt"
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
		vals := strings.Split(line, ": ")
		target, err := strconv.Atoi(vals[0])
		check(err)
		possibleInts := strings.Split(vals[1], " ")
		nums := getIntArr(possibleInts)
		if canCombine(target, nums[0], nums[1:]) {
			res += target
		}
	}
	return res
}

func handlePart2(input string) int {
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		vals := strings.Split(line, ": ")
		target, err := strconv.Atoi(vals[0])
		check(err)
		possibleInts := strings.Split(vals[1], " ")
		nums := getIntArr(possibleInts)
		if canCombine2(target, nums[0], nums[1:]) {
			res += target
		}
	}
	return res
}

func canCombine2(target int, currValue int, restValues []int) bool {
	if currValue == target && len(restValues) == 0 {
		return true
	}
	if currValue > target {
		return false
	}

	if len(restValues) == 0 {
		return false
	}

	return canCombine2(target, currValue+restValues[0], restValues[1:]) ||
		canCombine2(target, currValue*restValues[0], restValues[1:]) ||
		canCombine2(target, concat(currValue, restValues[0]), restValues[1:])
}

func canCombine(target int, currValue int, restValues []int) bool {
	if currValue == target && len(restValues) == 0 {
		return true
	}
	if currValue > target {
		return false
	}

	if len(restValues) == 0 {
		return false
	}

	return canCombine(target, currValue+restValues[0], restValues[1:]) || canCombine(target, currValue*restValues[0], restValues[1:])
}

func concat(v1, v2 int) int {
	val := fmt.Sprintf("%v%v", v1, v2)
	num, err := strconv.Atoi(val)
	check(err)
	return num
}

func getIntArr(input []string) []int {
	arr := make([]int, 0, len(input))
	for _, v := range input {
		num, err := strconv.Atoi(v)
		check(err)
		arr = append(arr, num)
	}
	return arr
}
