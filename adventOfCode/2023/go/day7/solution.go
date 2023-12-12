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

type HandType int

const (
	HighCard = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type CamelGame struct {
	Cards        string
	CardsAsValue []int
	BID          int
	Rank         int
	HandType     HandType
}

func getMapValues(source map[string]int) []int {
	values := make([]int, 0)

	for _, v := range source {
		values = append(values, v)
	}

	slices.Sort(values)
	return values
}

func isFiveOfAKind(cardOccurances map[string]int) bool {
	if len(cardOccurances) != 1 {
		return false
	}
	return true
}

func isFourOfAKind(cardOccurances map[string]int) bool {
	if len(cardOccurances) != 2 {
		return false
	}

	values := getMapValues(cardOccurances)

	if values[0] != 1 {
		return false
	}

	if values[1] != 4 {
		return false
	}
	return true
}

func isFullHouse(cardOccurances map[string]int) bool {
	if len(cardOccurances) != 2 {
		return false
	}

	values := getMapValues(cardOccurances)

	if values[0] != 2 {
		return false
	}

	if values[1] != 3 {
		return false
	}
	return true
}

func isThreeOfAKind(cardOccurances map[string]int) bool {
	if len(cardOccurances) != 3 {
		return false
	}

	values := getMapValues(cardOccurances)

	if values[0] != 1 {
		return false
	}

	if values[1] != 1 {
		return false
	}

	if values[2] != 3 {
		return false
	}
	return true
}

func isTwoPair(cardOccurances map[string]int) bool {
	if len(cardOccurances) != 3 {
		return false
	}

	values := getMapValues(cardOccurances)

	if values[0] != 1 {
		return false
	}

	if values[1] != 2 {
		return false
	}

	if values[2] != 2 {
		return false
	}
	return true
}

func isOnePair(cardOccurances map[string]int) bool {
	if len(cardOccurances) != 4 {
		return false
	}

	values := getMapValues(cardOccurances)

	if values[0] != 1 {
		return false
	}

	if values[1] != 1 {
		return false
	}

	if values[2] != 1 {
		return false
	}

	if values[3] != 2 {
		return false
	}
	return true
}

func (cg *CamelGame) getHandType() HandType {
	split := strings.Split(cg.Cards, "")
	occurances := make(map[string]int, 0)
	for _, lett := range split {
		occurances[lett]++
	}

	if isFiveOfAKind(occurances) {
		return FiveOfAKind
	} else if isFourOfAKind(occurances) {
		return FourOfAKind
	} else if isFullHouse(occurances) {
		return FullHouse
	} else if isThreeOfAKind(occurances) {
		return ThreeOfAKind
	} else if isTwoPair(occurances) {
		return TwoPair
	} else if isOnePair(occurances) {
		return OnePair
	}

	return HighCard
}

func newHand(cards, bid string) CamelGame {
	var game CamelGame
	bidConv, err := strconv.Atoi(bid)
	if err != nil {
		panic(2)
	}
	game.Cards = cards
	game.BID = bidConv
	game.HandType = game.getHandType()

	split := strings.Split(cards, "")
	cardsAsValue := make([]int, 0)
	for i := range split {
		if split[i] == "A" {
			cardsAsValue = append(cardsAsValue, 14)
		} else if split[i] == "K" {
			cardsAsValue = append(cardsAsValue, 13)
		} else if split[i] == "Q" {
			cardsAsValue = append(cardsAsValue, 12)
		} else if split[i] == "J" {
			cardsAsValue = append(cardsAsValue, 11)
		} else if split[i] == "T" {
			cardsAsValue = append(cardsAsValue, 10)
		} else {
			conv, err := strconv.Atoi(split[i])
			if err != nil {
				panic(3)
			}
			cardsAsValue = append(cardsAsValue, conv)
		}
	}

	game.CardsAsValue = cardsAsValue

	return game
}

func parseGames(lines []string) []CamelGame {
	games := make([]CamelGame, 0)
	for _, line := range lines {
		split := strings.Split(line, " ")
		games = append(games, newHand(split[0], split[1]))
	}
	return games
}

func compareCardByCard(a, b CamelGame) int {
	cardsA := a.CardsAsValue
	cardsB := b.CardsAsValue

	for i := range cardsA {
		if cardsA[i] > cardsB[i] {
			return 1
		} else if cardsA[i] < cardsB[i] {
			return -1
		}
	}
	return 0
}

func part1(filename string) int {
	lines := readFile(filename)
	games := parseGames(lines)
	slices.SortFunc(games, func(a, b CamelGame) int {
		if a.HandType > b.HandType {
			return 1
		} else if a.HandType < b.HandType {
			return -1
		} else {
			return compareCardByCard(a, b)
		}
	})
	res := 0
	for i := 1; i <= len(games); i++ {
		res += i * games[i-1].BID
	}

	return res
}

func part2(filename string) int {
	return -1
}

func main() {
	fileName := os.Args[1]
	fmt.Println(part1(fileName))
	// fmt.Println(part2(fileName))
}
