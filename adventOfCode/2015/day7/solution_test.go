package solution

import "testing"

func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
	cases := []struct {
		input    string
		cable    string
		response uint16
	}{
		{
			cable: "d",
			input: `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`,
			response: 72,
		},

		{
			cable: "e",
			input: `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`,
			response: 507,
		},
		{
			cable: "f",
			input: `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`,
			response: 492,
		},
		{
			cable: "g",
			input: `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`,
			response: 114,
		},
		{
			cable: "h",
			input: `123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i`,
			response: 65412,
		},
	}

	for _, testCase := range cases {
		res := handlePart1(testCase.input, testCase.cable)
		if testCase.response != res {
			t.Fatalf("Wrong, expected: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
		}
	}

}
func TestShouldReturnProperValueForProperInputPart1(t *testing.T) {
	res := part1()
	if res != 46065 {
		t.Fatal("Got:", res)
	}
}

func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
	res := part2()
	if res != 14134 {
		t.Fatal("Got:", res)
	}
}

// // func TestShouldReturnProperValueForTestInputPart2(t *testing.T) {
// // 	cases := []struct {
// // 		input    string
// // 		response int
// // 	}{
// // 		{
// // 			input:    "qjhvhtzxzqqjkmpb",
// // 			response: 1,
// // 		},
// // 		{
// // 			input:    "aaa",
// // 			response: 0,
// // 		},
// // 		{
// // 			input:    "uurcxstgmygtbstg",
// // 			response: 0,
// // 		},
// // 		{
// // 			input:    "ieodomkazucvgmuy",
// // 			response: 0,
// // 		},
// // 	}
// //
// // 	for _, testCase := range cases {
// // 		res := handlePart2(testCase.input)
// //
// // 		if testCase.response != res {
// // 			t.Fatalf("Wrong, expected: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
// // 		}
// // 	}
// //
// // }
//
// func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
// 	res := part2()
// 	if res != 15343601 {
// 		t.Fatal("Got:", res)
// 	}
// }
