package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCompletePart1WithTestInput(t *testing.T) {
	testCases := []struct {
		input []string
		res   int
	}{{
		input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		},
		res: 374,
	}}

	for _, tc := range testCases {
		res := handlePart1(tc.input, 1)
		assert.Equal(t, tc.res, res)
	}
}

func TestShouldCompletePart1WithProperInput(t *testing.T) {
	res := part1("input")

	if res != 9734203 {
		t.Fatalf("Wrong number, should be %d, current: %d", 0, res)
	}
}

func TestShouldCompletePart2WithTestInput(t *testing.T) {
	testCases := []struct {
		input []string
		res   int
	}{{
		input: []string{
			"...#......",
			".......#..",
			"#.........",
			"..........",
			"......#...",
			".#........",
			".........#",
			"..........",
			".......#..",
			"#...#.....",
		},
		res: 1030,
	}}

	for _, tc := range testCases {
		res := handlePart2(tc.input, 9)
		assert.Equal(t, tc.res, res)
	}
}

func TestShouldCompletePart2WithProperInput(t *testing.T) {
	res := part2("input")

	if res != 568914596391 {
		t.Fatalf("Wrong number, should be %d, current: %d", 0, res)
	}
}
