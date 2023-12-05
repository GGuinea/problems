package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Game struct {
	PlayerNumbers  []int
	WinningNumbers []int
	Result         int
}

var PowerOfTwo = map[int]int{}

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
	gamesList := make([]Game, 0)

	for i, line := range lines {
		game := parseGames(line, i+1)
		gamesList = append(gamesList, game)
	}

	res := 0
	for _, game := range gamesList {
		if game.Result != 0 {
			res += int(math.Pow(2, float64(game.Result)-1))
		}
	}

	return res
}

func parseGames(line string, i int) Game {
	 _, removedCard,_:= strings.Cut(line, ": ")
	winningAndYours := strings.Split(removedCard, " | ")
	winningNumbersAsString := strings.Split(winningAndYours[0], " ")
	yoursNumbersAsString := strings.Split(winningAndYours[1], " ")
	winningNumbers := []int{}
	yourNumbers := []int{}

	for i := range winningNumbersAsString {
		if winningNumbersAsString[i] == "" {
			continue
		}
		num, err := strconv.Atoi(winningNumbersAsString[i])

		if err != nil {
			fmt.Println(winningNumbersAsString[i])
			panic(2)
		}

		winningNumbers = append(winningNumbers, num)
	}

	for i := range yoursNumbersAsString {
		if yoursNumbersAsString[i] == "" {
			continue
		}

		num, err := strconv.Atoi(yoursNumbersAsString[i])
		if err != nil {
			fmt.Println(yoursNumbersAsString[i])
			panic(3)
		}

		yourNumbers = append(yourNumbers, num)
	}

	contained := 0
	for _, num := range yourNumbers {
		if slices.Contains(winningNumbers, num) {
			contained++
		}
	}

	return Game{PlayerNumbers: yourNumbers, WinningNumbers: winningNumbers, Result: contained}
}

func main() {
	fileName := os.Args[1]
	fmt.Println(part1(fileName))
}
