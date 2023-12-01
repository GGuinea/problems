package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func part2(fileName string) {
	lines := readFile(fileName)

	var numbersAsString []string
	numbersAsString = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	res := 0
	for _, line := range lines {
		keysFirst, keysLast := findAndReadNumberFromLine(line, numbersAsString)
		res += getNumber(keysFirst[0], 10)
		res += getNumber(keysLast[len(keysLast)-1], 1)
	}
	fmt.Println(res)
}

func getNumber(key string, multiplyBy int) int {
	num := 0
	switch key {
	case "one":
		num += 1 * multiplyBy
	case "two":
		num += 2 * multiplyBy
	case "three":
		num += 3 * multiplyBy
	case "four":
		num += 4 * multiplyBy
	case "five":
		num += 5 * multiplyBy
	case "six":
		num += 6 * multiplyBy
	case "seven":
		num += 7 * multiplyBy
	case "eight":
		num += 8 * multiplyBy
	case "nine":
		num += 9 * multiplyBy
		break
	default:
		conveted, err := strconv.Atoi(key)
		if err != nil {
			panic(3)
		}
		num += conveted * multiplyBy
	}
	return num
}

func findAndReadNumberFromLine(line string, numbersAsLett []string) ([]string, []string) {
	firstOccurance := make(map[string]int)

	for i := range numbersAsLett {
		index := strings.Index(line, numbersAsLett[i])
		if index != -1 {
			firstOccurance[numbersAsLett[i]] = index
		}
	}

	firstKeys := make([]string, 0, len(firstOccurance))

	for key := range firstOccurance {
		firstKeys = append(firstKeys, key)
	}

	sort.SliceStable(firstKeys, func(i, j int) bool {
		return firstOccurance[firstKeys[i]] < firstOccurance[firstKeys[j]]
	})

	lastOccurance := make(map[string]int)

	for i := range numbersAsLett {
		index := strings.LastIndex(line, numbersAsLett[i])
		if index != -1 {
			lastOccurance[numbersAsLett[i]] = index
		}
	}

	lastKeys := make([]string, 0, len(lastOccurance))

	for key := range lastOccurance {
		lastKeys = append(lastKeys, key)
	}

	sort.SliceStable(lastKeys, func(i, j int) bool {
		return lastOccurance[lastKeys[i]] < lastOccurance[lastKeys[j]]
	})

	return firstKeys, lastKeys
}

func main() {
	fileName := os.Args[1]
	//part1(fileName)
	part2(fileName)
}
