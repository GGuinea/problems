package main

import (
	"fmt"
	"reflect"
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

func TestParseSeedsPart2(t *testing.T) {
	inputLine := "seeds: 79 14 55 13"
	seeds := parseSeedsPart2(inputLine)
	fmt.Println(seeds)
	if len(seeds) != 27 {
		t.Fatalf("Should be 27 seeds, current: %d", len(seeds))
	}
}

func TestApplyRangesSeedRangeIsOverlapedByMapRange(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 10,
			SourceEnd:   20,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 3,
			SourceEnd:   30,
			Destination: 10,
			Range:       27,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 17,
			SourceEnd:   27,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangesSeedRangeIsOverlapedByMapRange2(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 1,
			SourceEnd:   2,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 1,
			SourceEnd:   10,
			Destination: 1,
			Range:       9,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 1,
			SourceEnd:   2,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangesSeedRangeIsOverlapedByMapRange3(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 10,
			SourceEnd:   20,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 1,
			SourceEnd:   20,
			Destination: 1,
			Range:       19,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 10,
			SourceEnd:   20,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangesSeedRangeOverlapingMapRange(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 2,
			SourceEnd:   40,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 3,
			SourceEnd:   30,
			Destination: 10,
			Range:       27,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 2,
			SourceEnd:   3,
		},
		{
			SourceStart: 10,
			SourceEnd:   37,
		},
		{
			SourceStart: 30,
			SourceEnd:   40,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangesSeedRangeOverlapingMapRange2(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 2,
			SourceEnd:   40,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 2,
			SourceEnd:   30,
			Destination: 10,
			Range:       27,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 2,
			SourceEnd:   3,
		},
		{
			SourceStart: 10,
			SourceEnd:   38,
		},
		{
			SourceStart: 30,
			SourceEnd:   40,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangesSeedRangeOverlapingMapRange3(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 2,
			SourceEnd:   40,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 30,
			SourceEnd:   40,
			Destination: 1,
			Range:       10,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 2,
			SourceEnd:   30,
		},
		{
			SourceStart: 1,
			SourceEnd:   11,
		},
		{
			SourceStart: 40,
			SourceEnd:   41,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangesSeedRangeHasCommonBeginningWithMapRange(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 0,
			SourceEnd:   10,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 3,
			SourceEnd:   30,
			Destination: 10,
			Range:       27,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 0,
			SourceEnd:   3,
		},
		{
			SourceStart: 10,
			SourceEnd:   17,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangesSeedRangeHasCommonBeginningWithMapRange2(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 0,
			SourceEnd:   10,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 10,
			SourceEnd:   20,
			Destination: 1,
			Range:       10,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 0,
			SourceEnd:   10,
		},
		{
			SourceStart: 1,
			SourceEnd:   2,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangesSeedRangeHasCommonEndWithMapRange(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 20,
			SourceEnd:   40,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 3,
			SourceEnd:   30,
			Destination: 10,
			Range:       27,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 27,
			SourceEnd:   37,
		},
		{
			SourceStart: 30,
			SourceEnd:   40,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangesSeedRangeHasCommonEndWithMapRange2(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 20,
			SourceEnd:   30,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 1,
			SourceEnd:   20,
			Destination: 1,
			Range:       19,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 20,
			SourceEnd:   21,
		},
		{
			SourceStart: 20,
			SourceEnd:   30,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangeSeedAndMapRangeDoesNotHaveCommonPart(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 0,
			SourceEnd:   10,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 20,
			SourceEnd:   30,
			Destination: 10,
			Range:       27,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 0,
			SourceEnd:   10,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestApplyRangeSeedAndMapRangeDoesNotHaveCommonPart2(t *testing.T) {
	seedsRanges := []RangeStruct{
		{
			SourceStart: 10,
			SourceEnd:   20,
			Destination: 0,
			Range:       0,
		},
	}

	mapRanges := []RangeStruct{
		{
			SourceStart: 1,
			SourceEnd:   9,
			Destination: 10,
			Range:       8,
		},
	}

	expected := []RangeStruct{
		{
			SourceStart: 10,
			SourceEnd:   20,
		},
	}

	newSeedRanges := applyRange(seedsRanges, mapRanges)
	res := reflect.DeepEqual(newSeedRanges, expected)

	if res != true {
		fmt.Println(newSeedRanges)
		t.Fatalf("Should be equal!")
	}
}

func TestShouldCompletePart2WithTestInput(t *testing.T) {
	testFileName := "test_input"
	expected := 46
	res := part2(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}

func TestShouldCompletePart2WithProperInput(t *testing.T) {
	testFileName := "input"
	expected := 52510809
	res := part2(testFileName)

	if res != expected {
		t.Fatalf("Wrong number, should be %d, current: %d", expected, res)
	}
}
