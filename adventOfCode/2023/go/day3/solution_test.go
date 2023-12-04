package main

import (
	"testing"
)

func TestShouldCompletePart1WithTestInput(t *testing.T) {
	testFileName := "test_input"
	expected := 4361 
	res := part1(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

func TestShouldCompletePart1WithProperInput(t *testing.T) {
	testFileName := "input"
	expected := 531561
	res := part1(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

func TestShouldReturnTrueIfStringInputIsAString(t *testing.T) {
	res := isSybmol("%")

	if res != true {
		t.Fatalf("Should return true")
	}

	res = isSybmol("&")

	if res != true {
		t.Fatalf("Should return true")
	}

	res = isSybmol("0")

	if res != false{
		t.Fatalf("Should return false")
	}

	res = isSybmol(".")

	if res != false{
		t.Fatalf("Should return false")
	}
}

// func TestShouldCompletePart2WithTestInput(t *testing.T) {
// 	testFileName := "test_input"
// 	expected := 2286
// 	res := part2(testFileName)
//
// 	if res != expected {
// 		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
// 	}
// }
//
// func TestShouldCompletePart2WithProperInput(t *testing.T) {
// 	testFileName := "input"
// 	expected := 72596
// 	res := part2(testFileName)
//
// 	if res != expected {
// 		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
// 	}
// }
