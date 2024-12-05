package day1

import (
	"testing"
)

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{input: `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`,
			res: 143,
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
	if res != 5639 {
		t.Fatalf("Got: %v", res)
	}
}

func TestShouldReturnProperValuesForPart2TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{input: `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`,
			res: 123,
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
	if res != 2046 {
		t.Fatalf("Got: %v", res)
	}

}
