package day1

import "testing"

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input   string
		seconds int
		res     int
		tilesX  int
		tilesY  int
	}{
		{
			seconds: 100,
			tilesX:  7,
			tilesY:  11,
			input: `p=0,4 v=3,-3
p=6,3 v=-1,-3
p=10,3 v=-1,2
p=2,0 v=2,-1
p=0,0 v=1,3
p=3,0 v=-2,-2
p=7,6 v=-1,-3
p=3,0 v=-1,-2
p=9,3 v=2,3
p=7,3 v=-1,2
p=2,4 v=2,-3
p=9,5 v=-3,-3`,
			res: 12,
		},
	}

	for _, tCase := range testCase {
		if res := handlePart1(tCase.input, tCase.seconds, tCase.tilesX, tCase.tilesY); res != tCase.res {
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
//
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
