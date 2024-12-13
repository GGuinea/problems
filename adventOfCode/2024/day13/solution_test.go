package day1

import "testing"

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{
			input: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400`,
			res: 280,
		},
		{
			input: `Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176`,
			res: 0,
		},
		{
			input: `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`,
			res: 480,
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
	if res != 35574 {
		t.Fatalf("Got: %v", res)
	}
}

func TestShouldReturnProperValuesForPart2(t *testing.T) {
	res := part2()
	if res != 80882098756071 {
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
// 			input: `AAAA
// BBCD
// BBCC
// EEEC`,
// 			res: 80,
// 		},
// 		{
// 			input: `RRRRIICCFF
// RRRRIICCCF
// VVRRRCCFFF
// VVRCCCJFFF
// VVVVCJJCFE
// VVIVCCJJEE
// VVIIICJJEE
// MIIIIIJJEE
// MIIISIJEEE
// MMMISSJEEE`,
// 			res: 1205,
// 		},
// 	}
//
// 	for _, tCase := range testCase {
// 		if res := handlePart2(tCase.input); res != tCase.res {
// 			t.Fatalf("Expected: %v, got: %v", tCase.res, res)
// 		}
// 	}
// }
