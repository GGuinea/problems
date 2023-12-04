package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Symbol struct {
	X     int
	Y     int
	Value string
}

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

func part1(filename string) int {
	lines := readFile(filename)

	inputArray := make([][]Symbol, 0)
	for i := range lines {
		parsedLine := []Symbol{}
		for j := range lines[i] {
			parsedLine = append(parsedLine, Symbol{X: i, Y: j, Value: string(lines[i][j])})
		}
		inputArray = append(inputArray, parsedLine)
	}

	res := 0

	for i := range inputArray {
		numberAsSlice := make([]string, 0)
		adjacent := false

		for j := range inputArray[i] {
			if inputArray[i][j].Value == "." || isSybmol(inputArray[i][j].Value) {
				if adjacent {
					numberAsString := strings.Join(numberAsSlice, "")
					number, err := strconv.Atoi(numberAsString)
					if err != nil {
						panic(5)
					}
					res += number
				}
				numberAsSlice = []string{}
				adjacent = false
			} else {
				numberAsSlice = append(numberAsSlice, inputArray[i][j].Value)
				if !adjacent {
					adjacent = isAdjacentToSymbol(i, j, inputArray)
				}

				if adjacent && j == len(inputArray[i])-1 {
					numberAsString := strings.Join(numberAsSlice, "")
					number, err := strconv.Atoi(numberAsString)
					if err != nil {
						panic(5)
					}
					res += number
				}
			}

		}
	}

	return res
}

func isAdjacentToSymbol(i, j int, array [][]Symbol) bool {
	if i == 0 && j == 0 {
		return isSybmol(array[i+1][j].Value) || isSybmol(array[i][j+1].Value) || isSybmol(array[i+1][j+1].Value)
	}

	if i == 0 && j == len(array[i])-1 {
		return isSybmol(array[i+1][j].Value) || isSybmol(array[i][j-1].Value) || isSybmol(array[i+1][j-1].Value)
	}

	if i == len(array)-1 && j == 0 {
		return isSybmol(array[i-1][j].Value) || isSybmol(array[i][j+1].Value) || isSybmol(array[i-1][j+1].Value)
	}

	if i == len(array)-1 && j == len(array[i])-1 {
		return isSybmol(array[i-1][j].Value) || isSybmol(array[i][j-1].Value) || isSybmol(array[i-1][j-1].Value)
	}

	if i != 0 && j == 0 {
		return isSybmol(array[i+1][j].Value) || isSybmol(array[i-1][j].Value) || isSybmol(array[i][j+1].Value) || isSybmol(array[i+1][j+1].Value) || isSybmol(array[i-1][j+1].Value)
	}

	if i == len(array)-1 {
		return isSybmol(array[i-1][j].Value) || isSybmol(array[i][j-1].Value) || isSybmol(array[i][j+1].Value) || isSybmol(array[i-1][j-1].Value) || isSybmol(array[i-1][j+1].Value)
	}

	if j == len(array[i])-1 {
		return isSybmol(array[i+1][j].Value) || isSybmol(array[i-1][j].Value) || isSybmol(array[i][j-1].Value) || isSybmol(array[i-1][j-1].Value) || isSybmol(array[i-1][j-1].Value)
	}

	if i == 0 && j != 0 {
		return isSybmol(array[i+1][j].Value) || isSybmol(array[i][j-1].Value) || isSybmol(array[i][j+1].Value) || isSybmol(array[i+1][j+1].Value) ||
			isSybmol(array[i+1][j-1].Value) || isSybmol(array[i+1][j+1].Value)
	}

	if i != 0 && j != 0 {
		return isSybmol(array[i+1][j].Value) || isSybmol(array[i][j-1].Value) || isSybmol(array[i][j+1].Value) || isSybmol(array[i-1][j].Value) ||
			isSybmol(array[i-1][j-1].Value) || isSybmol(array[i-1][j+1].Value) || isSybmol(array[i+1][j-1].Value) || isSybmol(array[i+1][j+1].Value)
	}

	panic(420)
}

func isSybmol(value string) bool {
	if value == "." {
		return false
	}

	if value[0] > 57 || value[0] < 48 {
		return true
	}
	return false
}

func main() {
	fileName := os.Args[1]
	fmt.Println(part1(fileName))
}
