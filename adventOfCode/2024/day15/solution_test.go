package day1

import "testing"

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{
		{
			input: `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`,
			res:   2028,
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

// func TestShouldReturnProperValuesForPart2(t *testing.T) {
// 	res := part2()
// 	if res != 80882098756071 {
// 		t.Fatalf("Got: %v", res)
// 	}
// }

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
