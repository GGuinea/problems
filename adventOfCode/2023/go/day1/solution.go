package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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


func part1(fileName string) {
	lines := readFile(fileName)
	res := 0
	for _, line := range lines {
		numberAsString := readNumberFromLine(line)
		firstAndLast := numberAsString[0]
		firstAndLast += numberAsString[len(numberAsString)-1]
		converted, err := strconv.Atoi(firstAndLast)
		if err != nil {
			panic(2)
		}
		res += converted
	}
	fmt.Println(res)
}

func readNumberFromLine(line string) []string {
	res := make([]string, 0)
	for _, char := range line {
		_, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}
		res = append(res, string(char))
	}
	return res
}


func main() {
	fileName := os.Args[1]
	part1(fileName)
}
