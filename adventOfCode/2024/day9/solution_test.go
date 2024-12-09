package day1

import (
	"fmt"
	"testing"
	"time"
)

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{input: "2333133121414131402",
			res: 1928,
		},
	}

	for _, tCase := range testCase {
		if res := handlePart1(tCase.input); res != tCase.res {
			t.Fatalf("Expected: %v, got: %v", tCase.res, res)
		}
	}
}

func TestShouldReturnProperValuesForPart1(t *testing.T) {
	res := part1()
	if res != 6384282079460 {
		t.Fatalf("Got: %v", res)
	}
}

func TestShouldReturnProperValuesForPart2TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{input: "2333133121414131402",
			res: 2858,
		},
		{input: "12345",
			res: 132,
		},
	}

	for _, tCase := range testCase {
		if res := handlePart2(tCase.input); res != tCase.res {
			t.Fatalf("Expected: %v, got: %v", tCase.res, res)
		}
	}
}

func TestShouldReturnProperValuesForPart2(t *testing.T) {
	timeStart := time.Now()
	res := part2()
	timeEnd := time.Now()
	if res != 6408966547049 {
		t.Fatalf("Got: %v", res)
	}
	fmt.Println("executed in: ", timeEnd.Sub(timeStart))
}

func BenchmarkShouldReturnProperValuesForPart2(t *testing.B) {
	res := part2()
	if res != 6408966547049 {
		t.Fatalf("Got: %v", res)
	}
}
