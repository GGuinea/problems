package solution

import (
	"os"
	"strconv"
	"strings"
)

const (
	SIZE_X = 1000
	SIZE_Y = 1000
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

func handlePart1(input string) int {
	split := strings.Split(input, "\n")
	res := 0
	for _, line := range split {
		str, _ := strconv.Unquote(line)
		//fmt.Println(line, str)
		res += len(line) - len(str)
	}
	return res
}

func part2() int {
	input := readFile("input")
	return handlePart2(input)
}

func handlePart2(input string) int {
	split := strings.Split(input, "\n")
	res := 0
	for _, line := range split {
		if len(line) == 0 {
			continue
		}
		str := strconv.Quote(line)
		res += len(str) - len(line)
	}
	return res
}
