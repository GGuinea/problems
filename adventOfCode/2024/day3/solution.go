package day1

import (
	"fmt"
	"os"
	"regexp"
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
	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	check(err)

	res := 0
	muls := r.FindAllString(input, -1)
	for _, one := range muls {
		res += doMul(one)
	}
	return res
}

func handlePart2(input string) int {
	r, err := regexp.Compile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
	check(err)

	res := 0
	muls := r.FindAllString(input, -1)
	enabled := true
	for _, one := range muls {
		if one == "do()" {
			enabled = true
			continue
		} else if one == "don't()" {
			enabled = false
			continue
		}
		if !enabled {
			continue
		}
		res += doMul(one)
	}
	return res
}

func doMul(one string) int {
	var x, y int
	n, err := fmt.Sscanf(one, "mul(%d,%d)", &x, &y)
	check(err)
	if n != 2 {
		panic("not expected number of args")
	}
	return x * y
}
