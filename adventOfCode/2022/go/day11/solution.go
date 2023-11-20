package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Monkey struct {
	StartingItems      []int
	Operation          string
	OperationValue     int
	OperationByItself  bool
	DestinationIfTrue  int
	DestinationIfFalse int
	TestValue          int
	InspectedItems     int
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

func part1(fileName string) {
	lines := readFile(fileName)
	monkeys := parseInput(lines)
	for j := 0; j < 20; j++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]
			for _, item := range monkey.StartingItems {
				dest, currItem := handleItem(item, monkey)
				destMonkey := monkeys[dest]
				destMonkey.StartingItems = append(destMonkey.StartingItems, currItem)
				monkeys[dest] = destMonkey
			}
			newStartingItems := make([]int, 0)
			itemsCount := len(monkey.StartingItems)
			monkey.StartingItems = newStartingItems
			monkey.InspectedItems += itemsCount
			monkeys[i] = monkey
		}
	}

	maxs := make([]int, 0)

	for i := 0; i < len(monkeys); i++ {
		maxs = append(maxs, monkeys[i].InspectedItems)
	}
	slices.Sort(maxs)
	fmt.Println(maxs[len(maxs)-1] * maxs[len(maxs)-2])
}

func part2(fileName string) {
	lines := readFile(fileName)
	monkeys := parseInput(lines)
	var lcm int = 1
	for i := 0; i < len(monkeys); i++ {
		lcm *= monkeys[i].TestValue
	}
	for j := 0; j < 10000; j++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]
			for _, item := range monkey.StartingItems {
				dest, currItem := handleItemPart2(item, monkey, lcm)
				destMonkey := monkeys[dest]
				destMonkey.StartingItems = append(destMonkey.StartingItems, currItem)
				monkeys[dest] = destMonkey
			}
			newStartingItems := make([]int, 0)
			itemsCount := len(monkey.StartingItems)
			monkey.StartingItems = newStartingItems
			monkey.InspectedItems += itemsCount
			monkeys[i] = monkey
		}
	}

	maxs := make([]int, 0)

	for i := 0; i < len(monkeys); i++ {
		maxs = append(maxs, monkeys[i].InspectedItems)
	}
	slices.Sort(maxs)
	fmt.Println(maxs[len(maxs)-1] * maxs[len(maxs)-2])
}

func parseInput(lines []string) map[int]Monkey {
	monkeys := make(map[int]Monkey)
	var currentMonkey Monkey
	monkeyNo := -1

	for _, line := range lines {
		line = strings.Trim(line, " ")
		if strings.Contains(line, "Monkey") {
			currentMonkey = Monkey{}
			monkeyNo++
		} else if strings.Contains(line, "Starting") {
			line, _ := strings.CutPrefix(line, "Starting items: ")
			splitted := strings.Split(line, ", ")
			items := make([]int, 0)
			for _, item := range splitted {
				conv, err := strconv.Atoi(item)
				if err != nil {
					panic(3)
				}
				items = append(items, conv)
			}
			currentMonkey.StartingItems = items
		} else if strings.Contains(line, "Operation") {
			line, _ := strings.CutPrefix(line, "Operation: new = old ")
			splitted := strings.Split(line, " ")
			currentMonkey.Operation = splitted[0]
			value, err := strconv.Atoi(splitted[1])
			if err != nil {
				currentMonkey.OperationByItself = true
			} else {
				currentMonkey.OperationByItself = false
				currentMonkey.OperationValue = value
			}

		} else if strings.Contains(line, "Test: ") {
			line, _ := strings.CutPrefix(line, "Test: divisible by ")
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(4)
			}
			currentMonkey.TestValue = value
		} else if strings.Contains(line, "true") {
			line, _ := strings.CutPrefix(line, "If true: throw to monkey ")
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(69)
			}
			currentMonkey.DestinationIfTrue = value
		} else if strings.Contains(line, "false") {
			line, _ := strings.CutPrefix(line, "If false: throw to monkey ")
			value, err := strconv.Atoi(line)
			if err != nil {
				panic(69)
			}
			currentMonkey.DestinationIfFalse = value
			monkeys[monkeyNo] = currentMonkey
		}
	}

	return monkeys
}

func handleItem(item int, monkey Monkey) (int, int) {
	item = makeOperation(monkey.Operation, monkey.OperationValue, item, monkey.OperationByItself)
	item = item / 3

	if item%monkey.TestValue == 0 {
		return monkey.DestinationIfTrue, item
	} else {
		return monkey.DestinationIfFalse, item
	}
}

func handleItemPart2(item int, monkey Monkey, div int) (int, int) {
	item = makeOperation(monkey.Operation, monkey.OperationValue, item, monkey.OperationByItself)
	item = item % div

	if item%monkey.TestValue == 0 {
		return monkey.DestinationIfTrue, item
	} else {
		return monkey.DestinationIfFalse, item
	}
}
func makeOperation(operation string, operand int, current int, byItself bool) int {
	switch operation {
	case "*":
		if byItself {
			return current * current
		}
		return current * operand
	case "+":
		if byItself {
			return current + current
		}
		return current + operand
	}
	fmt.Println(operation)
	panic(99)
}

func printMonkey(monkey Monkey) {
	fmt.Println("Starting items: ", monkey.StartingItems)
	fmt.Println("Operation: ", monkey.Operation)
	fmt.Println("Operation value: ", monkey.OperationValue)
	fmt.Println("Operation by itself: ", monkey.OperationByItself)
	fmt.Println("Destination if true: ", monkey.DestinationIfTrue)
	fmt.Println("Destination if false: ", monkey.DestinationIfFalse)
	fmt.Println("Test value: ", monkey.TestValue)
	fmt.Println()
}

func main() {
	fileName := os.Args[1]
	//part1(fileName)
	part2(fileName)
}
