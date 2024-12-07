package day1

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	s := []int{1, 2, 3, 4}
	fmt.Println(s[0])
	fmt.Println(s[1])
	fmt.Println(s[2:])
}

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{input: `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`,
			res: 3749,
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
	if res != 42283209483350 {
		t.Fatalf("Got: %v", res)
	}
}

func TestShouldReturnProperValuesForPart2TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{input: `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`,
			res: 11387,
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
	if res != 1026766857276279 {
		t.Fatalf("Got: %v", res)
	}
}
