package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Maxes struct {
	Blue                int
	Green               int
	Red                 int
	TotalMultipliedTurn int
}

func (m *Maxes) updateTotal() {
	m.TotalMultipliedTurn += m.Blue * m.Red * m.Green
	m.Blue = 0
	m.Red = 0
	m.Green = 0
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

func part1(fileName string) int {
	lines := readFile(fileName)
	redLimit := 12
	greenThreshold := 13
	blueThreshold := 14

	res := 0

	for gameId, line := range lines {
		if checkIfGameInLimit(redLimit, greenThreshold, blueThreshold, line, gameId+1) {
			res += gameId + 1
		}
	}

	return res
}

func checkIfGameInLimit(redLimit, greenLimit, blueLimit int, line string, gameId int) bool {
	line, _ = strings.CutPrefix(line, fmt.Sprintf("Game %d: ", gameId))
	turns := strings.Split(line, "; ")
	for i := range turns {
		if !checkIfTurnInLimit(redLimit, greenLimit, blueLimit, turns[i]) {
			return false
		}
	}
	return true
}

func checkIfTurnInLimit(redLimit, greenLimit, blueLimit int, turnString string) bool {
	splitted := strings.Split(turnString, ", ")

	for i := range splitted {
		if !checkIfWithinLimit(redLimit, greenLimit, blueLimit, splitted[i]) {
			return false
		}
	}

	return true
}

func checkIfWithinLimit(redLimit, greenLimit, blueLimit int, oneTurn string) bool {
	oneTurnSplitted := strings.Split(oneTurn, " ")
	convNum, err := strconv.Atoi(oneTurnSplitted[0])
	if err != nil {
		panic(4)
	}

	switch oneTurnSplitted[1] {
	case "green":
		return convNum <= greenLimit
	case "blue":
		return convNum <= blueLimit
	case "red":
		return convNum <= redLimit
	default:
		panic(3)
	}
}

func part2(fileName string) int {
	lines := readFile(fileName)

	maxes := &Maxes{}

	for gameId, line := range lines {
		calculateMaxs(maxes, line, gameId+1)
		maxes.updateTotal()
	}

	return maxes.TotalMultipliedTurn
}

func calculateMaxs(maxes *Maxes, line string, gameId int) {
	line, _ = strings.CutPrefix(line, fmt.Sprintf("Game %d: ", gameId))
	turns := strings.Split(line, "; ")
	for i := range turns {
		calculateMaxesInTurn(maxes, turns[i])
	}
}

func calculateMaxesInTurn(maxes *Maxes, turnString string) {
	splitted := strings.Split(turnString, ", ")

	for i := range splitted {
		checkMaxes(maxes, splitted[i])
	}
}

func checkMaxes(maxes *Maxes, oneTurn string) {
	oneTurnSplitted := strings.Split(oneTurn, " ")
	convNum, err := strconv.Atoi(oneTurnSplitted[0])
	if err != nil {
		panic(4)
	}

	switch oneTurnSplitted[1] {
	case "green":
		if maxes.Green < convNum {
			maxes.Green = convNum
		}
	case "blue":
		if maxes.Blue < convNum {
			maxes.Blue = convNum
		}
	case "red":
		if maxes.Red < convNum {
			maxes.Red = convNum
		}
	default:
		panic(3)
	}
}

func main() {
	fileName := os.Args[1]
	// res := part1(fileName)
	// fmt.Println(res)
	res := part2(fileName)
	fmt.Println(res)
}
