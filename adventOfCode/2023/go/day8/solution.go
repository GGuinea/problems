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

type Nodes struct {
	Left  string
	Right string
}

func parseNodes(lines []string) map[string]Nodes {
	res := make(map[string]Nodes, 0)

	for _, line := range lines {
		split := strings.Split(line, " = ")
		key := split[0]
		left, right := getLeftRight(split[1])
		res[key] = Nodes{
			Left:  left,
			Right: right,
		}
	}

	return res
}

func getLeftRight(nodes string) (string, string) {
	nodes, _ = strings.CutPrefix(nodes, "(")
	nodes, _ = strings.CutSuffix(nodes, ")")
	split := strings.Split(nodes, ", ")
	return split[0], split[1]
}

func part1(filename string) int {
	lines := readFile(filename)
	moves := lines[0]

	nodes := parseNodes(lines[2:])
	res := 0
	currentKey := "AAA"
	movePointer := 0

	for {
		direction := moves[movePointer]
		movePointer++

		if movePointer == len(moves) {
			movePointer = 0
		}
		if currentKey == "ZZZ" {
			break
		}
		nodes := nodes[currentKey]
		if string(direction) == "L" {
			currentKey = nodes.Left
		} else {
			currentKey = nodes.Right
		}
		res++
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
