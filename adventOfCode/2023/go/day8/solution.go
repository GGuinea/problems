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

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
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

func doesAllHaveEndingSuffix(keys []string) bool {
	for _, k := range keys {
		if !strings.HasSuffix(k, "Z") {
			return false
		}
	}
	return true
}

func getNewNodes(currentNodes []string, nodes map[string]Nodes, direction string) []string {
	newCurrentNodes := make([]string, 0)
	for _, k := range currentNodes {
		if string(direction) == "L" {
			newCurrentNodes = append(newCurrentNodes, nodes[k].Left)
		} else {
			newCurrentNodes = append(newCurrentNodes, nodes[k].Right)
		}
	}
	return newCurrentNodes
}

func part2(filename string) int {
	lines := readFile(filename)
	moves := lines[0]

	nodes := parseNodes(lines[2:])

	startingKeys := make([]string, 0)

	for k := range nodes {
		if strings.HasSuffix(k, "A") {
			startingKeys = append(startingKeys, k)
		}
	}

	allResults := make([]int, 0)
	fmt.Println(startingKeys)
	for _, k := range startingKeys {
		currentKey := k
		res := 0
		movePointer := 0
		for {
			direction := moves[movePointer]
			movePointer++

			if movePointer == len(moves) {
				movePointer = 0
			}

			if strings.HasSuffix(currentKey, "Z") {
				allResults = append(allResults, res)
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
	}

	return LCM(allResults[0], allResults[1], allResults[2:]...)
}

func main() {
	fileName := os.Args[1]
	fmt.Println(part1(fileName))
	fmt.Println(part2(fileName))
}
