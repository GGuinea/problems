package solution

import "testing"

func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
	cases := []struct {
		input      string
		cable      string
		response   string
		iterations int
	}{
		{
			input:      "1",
			response:   "11",
			iterations: 1,
		},
		{
			input:      "1",
			response:   "21",
			iterations: 2,
		},
		{
			input:      "1",
			response:   "1211",
			iterations: 3,
		},
		{
			input:      "1",
			response:   "111221",
			iterations: 4,
		},
		{
			input:      "1",
			response:   "312211",
			iterations: 5,
		},
	}

	for _, testCase := range cases {
		res := handlePart1(testCase.input, testCase.iterations)
		if testCase.response != res {
			t.Fatalf("Wrong, expected: %v, actual: %v; input: %v; iter: %v", testCase.response, res, testCase.input, testCase.iterations)
		}
	}

}

func TestShouldReturnProperValueForProperInputPart1(t *testing.T) {
	res := part1()
	if res != 492982 {
		t.Fatal("Got:", res)
	}
}

func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
	res := part2()
	if res != 6989950 {
		t.Fatal("Got:", res)
	}
}

//
// func TestShouldReturnProperValueForTestInputPart2(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		cable    string
// 		response int
// 	}{
// 		{
// 			input: `""
// "abc"
// "aaa\"aaa"
// "\x27"`,
// 			response: 19,
// 		},
// 	}
// 	for _, testCase := range cases {
// 		res := handlePart2(testCase.input)
//
// 		if testCase.response != res {
// 			t.Fatalf("Wrong, expected: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
// 		}
// 	}
//
// }
