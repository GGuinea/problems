package solution

import "testing"

// func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		cable    string
// 		response int
// 	}{
// 		{
// 			input: `London to Dublin = 464
// London to Belfast = 518
// Dublin to Belfast = 141`,
// 			response: 605,
// 		},
// 	}
//
// 	for _, testCase := range cases {
// 		res := handlePart1(testCase.input)
// 		if testCase.response != res {
// 			t.Fatalf("Wrong, expected: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
// 		}
// 	}
// }
// func TestShouldReturnProperValueForProperInputPart1(t *testing.T) {
// 	res := part1()
// 	if res != 117 {
// 		t.Fatal("Got:", res)
// 	}
// }
//
// func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
// 	res := part2()
// 	if res != 909 {
// 		t.Fatal("Got:", res)
// 	}
// }
//
// func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
// 	cases := []struct {
// 		input string
// 		cable string
// 		min   int
// 		max   int
// 	}{
// 		{
// 			input: `London to Dublin = 464
// London to Belfast = 518
// Dublin to Belfast = 141`,
// 			min: 605,
// 			max: 982,
// 		},
// 	}
//
// 	for _, testCase := range cases {
// 		min, max := handlePart1_dfs(testCase.input)
// 		if testCase.min != min || testCase.max != max {
// 			t.Fatalf("Wrong, expected: %v, actual: %v; input: %v", testCase.min, min, testCase.input)
// 		}
// 	}
// }

func BenchmarkPart1Old(t *testing.B) {
	min := part1()
	if min != 117 {
		t.Fatal("Problem with min:", min)
	}
}

func BenchmarkPart1DFS(t *testing.B) {
	min, _:= part1_dfs()
	if min != 117 {
		t.Fatal("Problem with min:", min)
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
