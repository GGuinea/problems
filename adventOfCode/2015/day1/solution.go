package main

import "os"

func main() {
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) string {
	data, err := os.ReadFile(filename)
	check(err)
	return string(data)
}

func part1Solution(input string) int {
	currentFloor := 0
	for _, v := range input {
		if v == '(' {
			currentFloor++
		} else if v == ')' {
			currentFloor--
		}
	}
	return currentFloor
}

func handlePart1(filename string) int {
	dataFromFile := readFile(filename)
	return part1Solution(dataFromFile)
}

func part2Solution(input string) int {
	currentFloor := 0
	for idx, v := range input {
		if v == '(' {
			currentFloor++
		} else if v == ')' {
			currentFloor--
		}
		if currentFloor == -1 {
			return idx + 1
		}
	}
	return -1
}

func handlePart2(filename string) int {
	dataFromFile := readFile(filename)
	return part2Solution(dataFromFile)
}
