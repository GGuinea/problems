package day1

import "testing"

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input string
		blinks int
		res   int
	}{

		{
			input: `125 17`,
			blinks: 25,
			res: 55312,
		},
	}

	for _, tCase := range testCase {
		if res := handlePart1(tCase.input, tCase.blinks); res != tCase.res {
			t.Fatalf("Expected: %v, got: %v", tCase.res, res)
		}
	}
}

func TestShouldReturnProperValuesForPart1(t *testing.T) {
	res := part1()
	if res != 218956 {
		t.Fatalf("Got: %v", res)
	}
}

func TestShouldReturnProperValuesForPart2(t *testing.T) {
	res := part2()
	if res != 21790168 {
		t.Fatalf("Got: %v", res)
	}
}

// func TestShouldReturnProperValuesForPart2TestInput(t *testing.T) {
// 	testCase := []struct {
// 		input string
// 		res   int
// 	}{
//
// 		{
// 			input: `89010123
// 78121874
// 87430965
// 96549874
// 45678903
// 32019012
// 01329801
// 10456732`,
// 			res: 81,
// 		},
// 	}
//
// 	for _, tCase := range testCase {
// 		if res := handlePart2(tCase.input); res != tCase.res {
// 			t.Fatalf("Expected: %v, got: %v", tCase.res, res)
// 		}
// 	}
// }
