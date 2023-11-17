package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	stacks := make(map[int][]string)
	for _, line := range lines {
		if strings.Contains(line, "move") {
			break
		}
		parseStacks(line, stacks)
	}

	for _, line := range lines {
		if !strings.Contains(line, "move") {
			continue
		}
		from, to, count := parseMoves(line)
		makeMove(from, to, count, stacks)
	}

	var res string
	for i := 1; i < len(stacks)+1; i++ {
		res += strings.Join(stacks[i][:1], "")
	}
	fmt.Println(res)
}

func parseStacks(line string, stacks map[int][]string) {
	stackPointer := 0
	for i := 0; i < len(line); i++ {
		if i%4 == 0 {
			stackPointer++
		}

		character := line[i]
		if character >= 'A' && character <= 'Z' {
			stacks[stackPointer] = append(stacks[stackPointer], string(character))
		}
	}
}

func makeMove(from, to, count int, stacks map[int][]string) {
	movedObjects := make([]string, count)
	copy(movedObjects, stacks[from][:count])
	slices.Reverse(movedObjects)
	stacks[to] = append(movedObjects, stacks[to]...)
	stacks[from] = stacks[from][count:]
}

func parseMoves(line string) (int, int, int) {
	line = strings.Replace(line, "move ", "", 1)
	line = strings.Replace(line, " from ", ",", 1)
	line = strings.Replace(line, " to ", ",", 1)
	splitted := strings.Split(line, ",")
	count, _ := strconv.Atoi(splitted[0])
	from, _ := strconv.Atoi(splitted[1])
	to, _ := strconv.Atoi(splitted[2])
	return from, to, count
}

func part2() {
	lines := readFile("input")
	stacks := make(map[int][]string)
	for _, line := range lines {
		if strings.Contains(line, "move") {
			break
		}
		parseStacks(line, stacks)
	}

	for _, line := range lines {
		if !strings.Contains(line, "move") {
			continue
		}
		from, to, count := parseMoves(line)
		makeMoveNewCrateMover(from, to, count, stacks)
	}

	var res string
	for i := 1; i < len(stacks)+1; i++ {
		res += strings.Join(stacks[i][:1], "")
	}
	fmt.Println(res)
}

func makeMoveNewCrateMover(from, to, count int, stacks map[int][]string) {
	movedObjects := make([]string, count)
	copy(movedObjects, stacks[from][:count])
	stacks[to] = append(movedObjects, stacks[to]...)
	stacks[from] = stacks[from][count:]
}

func main() {
	//part1()
	part2()
}
