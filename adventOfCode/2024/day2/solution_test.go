package day1

import "testing"

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			res: 2,
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
	if res != 224 {
		t.Fatalf("Got: %v", res)
	}
}

func TestShouldReturnProperValuesForPart2TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{
			input: `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			res: 4,
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
	if res != 293 {
		t.Fatalf("Got: %v", res)
	}
}
