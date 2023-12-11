package main

import (
	"testing"
)

func TestGetDestination(t *testing.T) {
	rangeStruct := RangeStruct{
		SourceStart: 98,
		SourceEnd:   100,
		Destination: 50,
		Range:       2,
	}

	inRange, val := rangeStruct.getDestination(98)

	if inRange == false {
		t.Fatalf("Should be in range!")
	}

	if val != 50 {
		t.Fatalf("Should be 50, acutal: %d", val)
	}
}

func TestGetDestination2(t *testing.T) {
	rangeStruct := RangeStruct{
		SourceStart: 50,
		SourceEnd:   98,
		Destination: 52,
		Range:       48,
	}

	inRange, val := rangeStruct.getDestination(50)

	if inRange == false {
		t.Fatalf("Should be in range!")
	}

	if val != 52 {
		t.Fatalf("Should be 50, acutal: %d", val)
	}
}

func TestGetDestination3(t *testing.T) {
	rangeStruct := RangeStruct{
		SourceStart: 50,
		SourceEnd:   98,
		Destination: 52,
		Range:       48,
	}

	inRange, val := rangeStruct.getDestination(1)

	if inRange == true {
		t.Fatalf("Should not be in range!")
	}

	if val != 1 {
		t.Fatalf("Should be 1, acutal: %d", val)
	}
}

func TestShouldCompletePart1WithTestInput(t *testing.T) {
	testFileName := "test_input"
	expected := 35
	res := part1(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

func TestShouldCompletePart1WithProperInput(t *testing.T) {
	testFileName := "input"
	expected := 662197086
	res := part1(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

// func TestShouldCompletePart2WithTestInput(t *testing.T) {
// 	testFileName := "test_input"
// 	expected := 30
// 	res := part2(testFileName)
//
// 	if res != expected {
// 		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
// 	}
// }
//
// func TestShouldCompletePart2WithProperInput(t *testing.T) {
// 	testFileName := "input"
// 	expected := 6420979
// 	res := part2(testFileName)
//
// 	if res != expected {
// 		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
// 	}
// }
