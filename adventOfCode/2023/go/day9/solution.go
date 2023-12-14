package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func getNumbers(lines []string) map[int][]int {
	numbers := make(map[int][]int, 0)
	for k, line := range lines {
		numbers[k] = make([]int, 0)
		split := strings.Split(line, " ")
		for i := range split {
			conv, err := strconv.Atoi(split[i])
			if err != nil {
				panic(3)
			}
			numbers[k] = append(numbers[k], conv)
		}
	}
	return numbers
}

func all(numberLine []int, value int) bool {
	for _, num := range numberLine {
		if num != value {
			return false
		}
	}

	return true
}

func calculateMissingParts(numberLine []int) int {
	nextLine := make([]int, 0)

	for i := 0; i < len(numberLine)-1; i++ {
		diff := numberLine[i+1] - numberLine[i]
		nextLine = append(nextLine, diff)
	}

	if all(nextLine, 0) {
		return numberLine[len(numberLine)-1]
	} else {
		res := numberLine[len(numberLine)-1] + calculateMissingParts(nextLine)
		return res
	}
}

func part1(filename string) int {
	lines := readFile(filename)
	numbers := getNumbers(lines)
	res := 0

	for _, v := range numbers {
		sum := calculateMissingParts(v)
		res += sum
	}

	return res
}

func part2(filename string) int {
	return -1
}

func main() {
	fileName := os.Args[1]
	fmt.Println(part1(fileName))
	fmt.Println(part2(fileName))
}
