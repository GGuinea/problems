package day1

import (
	"testing"
)

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			res: 41,
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
	if res != 5312 {
		t.Fatalf("Got: %v", res)
	}
}

func TestShouldReturnProperValuesForPart2TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{input: `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`,
			res: 6,
		},
	}

	for _, tCase := range testCase {
		if res := handlePart2(tCase.input); res != tCase.res {
			t.Fatalf("Expected: %v, got: %v", tCase.res, res)
		}
	}
}

func TestShouldReturnProperValuesForPart2(t *testing.T) {
	res := part2()
	if res != 1748 {
		t.Fatalf("Got: %v", res)
	}

}
