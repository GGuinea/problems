package main

import (
	"bufio"
	"fmt"
	"os"
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

func part1(fileName string) {
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

	fmt.Println(res)

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
			return  false
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

func main() {
	fileName := os.Args[1]
	part1(fileName)
}
