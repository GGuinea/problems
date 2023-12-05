package main

import (
	"testing"
)

func TestShouldCompletePart1WithTestInput(t *testing.T) {
	testFileName := "test_input"
	expected := 13
	res := part1(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

func TestShouldCompletePart1WithProperInput(t *testing.T) {
	testFileName := "input"
	expected := 25174
	res := part1(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

// func TestShouldCompletePart2WithTestInput(t *testing.T) {
// 	testFileName := "test_input"
// 	expected := 467835
// 	res := part2(testFileName)
//
// 	if res != expected {
// 		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
// 	}
// }
//
// func TestShouldCompletePart2WithProperInput(t *testing.T) {
// 	testFileName := "input"
// 	expected := 83279367
// 	res := part2(testFileName)
//
// 	if res != expected {
// 		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
// 	}
// }
