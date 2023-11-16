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

func part1() {
	lines := readFile("input")
	res := 0
	for _, line := range lines {
		splittedByComma := strings.Split(line, ",")
		xSplitted := strings.Split(splittedByComma[0], "-")
		ySplitted := strings.Split(splittedByComma[1], "-")
		x1, _ := strconv.Atoi(xSplitted[0])
		x2, _ := strconv.Atoi(xSplitted[1])
		y1, _ := strconv.Atoi(ySplitted[0])
		y2, _ := strconv.Atoi(ySplitted[1])
		if contains(x1, x2, y1, y2) {
			res++
		}
	}
	fmt.Println(res)
}

func contains(x1, x2, y1, y2 int) bool {
	if x1 <= y1 && x2 >= y2 {
		return true
	} else if x1 >= y1 && x2 <= y2 {
		return true
	}

	return false
}

func main() {
	part1()
}
