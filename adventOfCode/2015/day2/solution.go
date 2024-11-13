package solution

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readFile(filename string) string {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(res)
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
		curves := strings.Split(line, "x")
		if len(curves) == 1 {
			continue
		}
		l := getInt(curves[0])
		w := getInt(curves[1])
		h := getInt(curves[2])

		firstSide := l * w
		secondSide := l * h
		thirdSide := w * h
		smallest := getSmallest(firstSide, secondSide, thirdSide)
		fmt.Println("smallest:", smallest)
		res += 2*(firstSide+secondSide+thirdSide) + smallest
		fmt.Println(res)
	}
	return res
}

func handlePart2(input string) int {
	lines := strings.Split(input, "\n")
	res := 0
	for _, line := range lines {
		curves := strings.Split(line, "x")
		if len(curves) == 1 {
			continue
		}
		l := getInt(curves[0])
		w := getInt(curves[1])
		h := getInt(curves[2])
		sides := []int{l, w, h}
		slices.Sort(sides)
		res += 2*sides[0] + 2*sides[1] + (l * w * h)
	}
	return res
}

func getInt(almostInt string) int {
	res, _ := strconv.Atoi(almostInt)
	return res
}

func getSmallest(a, b, c int) int {
	fmt.Println(a, b, c)
	if a <= b && a <= c {
		return a
	}

	if b <= a && b <= c {
		return b
	}

	return c
}
