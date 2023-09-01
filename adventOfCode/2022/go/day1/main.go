package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func readFile(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
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

func day1(lines []string) {
	globalMax := 0
	localMax := 0
	for i := 0; i < len(lines); i++ {

		if len(lines[i]) == 0 {
			if localMax >= globalMax {
				globalMax = localMax
			}
			localMax = 0
		} else {
			number, _ := strconv.Atoi(lines[i])
			localMax += number
		}
	}
	fmt.Println(globalMax)
}

func day2(lines []string) {
	globalMax := make([]int, 0)
	localMax := 0
	for i := 0; i < len(lines); i++ {

		if len(lines[i]) == 0 {
			globalMax = append(globalMax, localMax)
			localMax = 0
		} else {
			number, _ := strconv.Atoi(lines[i])
			localMax += number
		}
	}
	sort.Ints(globalMax)
	fmt.Println(sum(globalMax[len(globalMax)-3:]))
}

func sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func main() {
	numbers := readFile("input")
	day1(numbers)
	day2(numbers)
}
