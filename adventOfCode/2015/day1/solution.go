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

func part1(filename string) int {
	dataFromFile := readFile(filename)
	currentFloor := 0
	for _, v := range dataFromFile {
		if v == '(' {
			currentFloor++
		} else if v == ')' {
			currentFloor--
		}
	}
	return currentFloor
}
