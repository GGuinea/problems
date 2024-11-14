package solution

import "testing"

func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
	cases := []struct {
		input    string
		response int
	}{
		{
			input:    "turn on 0,0 through 999,999",
			response: 1000000,
		},
		{
			input:    "toggle 0,0 through 999,0",
			response: 1000,
		},
		{
			input:    "turn off 499,499 through 500,500",
			response: 0,
		},
	}

	for _, testCase := range cases {
		res := handlePart1(testCase.input)
		if testCase.response != res {
			t.Fatalf("Wrong, expeced: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
		}
	}

}
func TestShouldReturnProperValueForProperInputPart1(t *testing.T) {
	res := part1()
	if res != 400410 {
		t.Fatal("Got:", res)
	}
}

// func TestShouldReturnProperValueForTestInputPart2(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		response int
// 	}{
// 		{
// 			input:    "qjhvhtzxzqqjkmpb",
// 			response: 1,
// 		},
// 		{
// 			input:    "aaa",
// 			response: 0,
// 		},
// 		{
// 			input:    "uurcxstgmygtbstg",
// 			response: 0,
// 		},
// 		{
// 			input:    "ieodomkazucvgmuy",
// 			response: 0,
// 		},
// 	}
//
// 	for _, testCase := range cases {
// 		res := handlePart2(testCase.input)
//
// 		if testCase.response != res {
// 			t.Fatalf("Wrong, expeced: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
// 		}
// 	}
//
// }

func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
	res := part2()
	if res != 15343601 {
		t.Fatal("Got:", res)
	}
}
