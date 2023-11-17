package main

import (
	"bufio"
	"fmt"
	"os"
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

func part1() {
	lines := readFile("input")
	res := findUniqueSlidingWindow(lines[0], 4)
	fmt.Println(res)
}

func findUniqueSlidingWindow(line string, size int) int {
	splitted := strings.Split(line, "")
	for i := 0; i < len(line); i++ {
		if isUnique(splitted[i : i+size]) {
			return i + size
		}
	}
	return -1
}

func isUnique(window []string) bool {
	occuranceMap := make(map[string]int)
	for _, lett := range window {
		_, ok := occuranceMap[lett]
		if ok {
			return false
		}
		occuranceMap[lett] = 1
	}
	return true
}

func part2() {
	lines := readFile("input")
	res := findUniqueSlidingWindow(lines[0], 14)
	fmt.Println(res)
}

func main() {
	part1()
	part2()
}
