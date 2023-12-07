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
	Number         int
}

type Stack struct {
	elements []int
}

func (s *Stack) Push(element int) {
	s.elements = slices.Insert(s.elements, 0, element)
}

func (s *Stack) Pop() (int, int) {
	elem := s.elements[0]
	count := s.countElements(elem)
	s.elements = s.elements[count:]
	return elem, count
}

func (s *Stack) countElements(x int) int {
	for i, elem := range s.elements {
		if elem > x {
			return i
		}
	}
	return len(s.elements)
}

func (s *Stack) Sort() {
	slices.Sort(s.elements)
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
	gamesList := make([]Game, 0)

	for i, line := range lines {
		game := parseGames(line, i)
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

func part2(filename string) int {
	lines := readFile(filename)
	gamesList := make([]Game, 0)
	cardCounter := make([]int, 0)

	for i, line := range lines {
		game := parseGames(line, i)
		gamesList = append(gamesList, game)
		cardCounter = append(cardCounter, 1)
	}

	res := 0
	for i := 0; i < len(gamesList); i++ {
		res += cardCounter[i]
		game := gamesList[i]
		if game.Result != 0 {
			for j := 1; j <= game.Result; j++ {
				cardCounter[i+j] += cardCounter[i]
			}
		}
	}
	return res
}

func parseGames(line string, i int) Game {
	_, removedCard, _ := strings.Cut(line, ": ")
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

	return Game{PlayerNumbers: yourNumbers, WinningNumbers: winningNumbers, Result: contained, Number: i}
}

func main() {
	fileName := os.Args[1]
	//fmt.Println(part1(fileName))
	fmt.Println(part2(fileName))
}
