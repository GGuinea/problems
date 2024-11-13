package solution

import "testing"

func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
	cases := []struct {
		input    string
		response int
	}{
		{
			input:    "ugknbfddgicrmopn",
			response: 1,
		},
		{
			input:    "aaa",
			response: 1,
		},
		{
			input:    "jchzalrnumimnmhp",
			response: 0,
		},
		{
			input:    "haegwjzuvuyypxyu",
			response: 0,
		},
		{
			input:    "dvszwmarrgswjxmb",
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
	if res != 238 {
		t.Fatal("Got:", res)
	}
}

func TestShouldReturnProperValueForTestInputPart2(t *testing.T) {
	cases := []struct {
		input    string
		response int
	}{
		{
			input:    "qjhvhtzxzqqjkmpb",
			response: 1,
		},
		{
			input:    "aaa",
			response: 0,
		},
		{
			input:    "uurcxstgmygtbstg",
			response: 0,
		},
		{
			input:    "ieodomkazucvgmuy",
			response: 0,
		},
	}

	for _, testCase := range cases {
		res := handlePart2(testCase.input)

		if testCase.response != res {
			t.Fatalf("Wrong, expeced: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
		}
	}

}

func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
 	res := part2()
 	if res != 9962624 {
 		t.Fatal("Got:", res)
 	}
 }
