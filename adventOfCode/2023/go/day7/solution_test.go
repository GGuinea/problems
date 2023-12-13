package main

import (
	"testing"
)

func TestGetHandTypeFiveOfAKind(t *testing.T) {
	game := CamelGame{Cards: "AAAAA"}
	handType := game.getHandType()

	if handType != FiveOfAKind {
		t.Fatalf("Wrong! current: %d, should be: %d", handType, FiveOfAKind)
	}
}

func TestGetHandTypeFourOfAKidn(t *testing.T) {
	game := CamelGame{Cards: "AA8AA"}
	handType := game.getHandType()

	if handType != FourOfAKind {
		t.Fatalf("Wrong! current: %d, should be: %d", handType, FourOfAKind)
	}
}

func TestGameHandTypeFullHouse(t *testing.T) {
	game := CamelGame{Cards: "23332"}
	handType := game.getHandType()

	if handType != FullHouse {
		t.Fatalf("Wrong! current: %d, should be: %d", handType, FullHouse)
	}
}

func TestGameThreeOfAKidn(t *testing.T) {
	game := CamelGame{Cards: "TTT98"}
	handType := game.getHandType()

	if handType != ThreeOfAKind {
		t.Fatalf("Wrong! current: %d, should be: %d", handType, ThreeOfAKind)
	}
}

func TestGameTwoPair(t *testing.T) {
	game := CamelGame{Cards: "23432"}
	handType := game.getHandType()

	if handType != TwoPair {
		t.Fatalf("Wrong! current: %d, should be: %d", handType, TwoPair)
	}
}

func TestGamePair(t *testing.T) {
	game := CamelGame{Cards: "A23A4"}
	handType := game.getHandType()

	if handType != OnePair {
		t.Fatalf("Wrong! current: %d, should be: %d", handType, OnePair)
	}
}

func TestGameHighCard(t *testing.T) {
	game := CamelGame{Cards: "23456"}
	handType := game.getHandType()

	if handType != HighCard {
		t.Fatalf("Wrong! current: %d, should be: %d", handType, HighCard)
	}
}

func TestShouldCompletePart1WithTestInput(t *testing.T) {
	testFileName := "test_input"
	expected := 6440
	res := part1(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

func TestShouldCompletePart1WithProperInput(t *testing.T) {
	testFileName := "input"
	expected := 248559379
	res := part1(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

func TestShouldCompletePart2WithTestInput(t *testing.T) {
	testFileName := "test_input"
	expected := 5905
	res := part2(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

func TestShouldCompletePart2WithProperInput(t *testing.T) {
	testFileName := "input"
	expected := 249631254
	res := part2(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}
