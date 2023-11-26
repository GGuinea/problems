package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type list struct {
	Value  int
	Inner  []*list
	Parent *list
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
		if len(line) == 0 {
			continue
		}
		lines = append(lines, line)
	}

	return lines
}

func part1(fileName string) {
	lines := readFile(fileName)
	total := 0
	index := 1
	for i := 0; i < len(lines); i += 2 {
		l := parseLine(lines[i])
		r := parseLine(lines[i+1])

		res := compare(l, r)
		if res == 1 {
			total += index
		}

		index++
	}
	fmt.Println(total)
}

func compare(left, right list) int {
	switch {
	case len(left.Inner) == 0 && len(right.Inner) == 0:
		if left.Value > right.Value {
			return -1
		} else if left.Value == right.Value {
			return 0
		}
		return 1

	case left.Value >= 0:
		return compare(list{-1, []*list{&left}, nil}, right)
	case right.Value >= 0:
		return compare(left, list{-1, []*list{&right}, nil})
	default:
		var i int
		for i = 0; i < len(left.Inner) && i < len(right.Inner); i++ {
			isOrder := compare(*left.Inner[i], *right.Inner[i])
			if isOrder != 0 {
				return isOrder
			}
		}

		if i < len(left.Inner) {
			return -1
		} else if i < len(right.Inner) {
			return 1
		}

	}
	return 0
}

func parseLine(line string) list {
	root := list{-1, []*list{}, nil}
	tmp := &root

	var currentNumber string
	for _, lett := range line {
		switch lett {
		case '[':
			newTree := list{-1, []*list{}, tmp}
			tmp.Inner = append(tmp.Inner, &newTree)
			tmp = &newTree
			break
		case ']':
			if len(currentNumber) > 0 {
				number, err := strconv.Atoi(currentNumber)
				if err != nil {
					panic(2)
				}
				tmp.Value = number
				currentNumber = ""
			}
			tmp = tmp.Parent
			break
		case ',':
			if len(currentNumber) > 0 {
				number, err := strconv.Atoi(currentNumber)
				if err != nil {
					panic(3)
				}
				tmp.Value = number
				currentNumber = ""
			}
			tmp = tmp.Parent
			newTree := list{-1, []*list{}, tmp}
			tmp.Inner = append(tmp.Inner, &newTree)
			tmp = &newTree
		default:
			currentNumber += string(lett)
		}
	}

	return root
}

func main() {
	fileName := os.Args[1]
	part1(fileName)
}
