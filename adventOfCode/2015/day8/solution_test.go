package solution

import "testing"

func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
	cases := []struct {
		input    string
		cable    string
		response int
	}{
		{
			input: `""
"abc"
"aaa\"aaa"
"\x27"`,
			response: 12,
		},
	}

	for _, testCase := range cases {
		res := handlePart1(testCase.input)
		if testCase.response != res {
			t.Fatalf("Wrong, expected: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
		}
	}

}
func TestShouldReturnProperValueForProperInputPart1(t *testing.T) {
	res := part1()
	if res != 1342 {
		t.Fatal("Got:", res)
	}
}

func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
	res := part2()
	if res != 14134 {
		t.Fatal("Got:", res)
	}
}

func TestShouldReturnProperValueForTestInputPart2(t *testing.T) {
	cases := []struct {
		input    string
		cable    string
		response int
	}{
		{
			input: `""
"abc"
"aaa\"aaa"
"\x27"`,
			response: 19,
		},
	}
	for _, testCase := range cases {
		res := handlePart2(testCase.input)

		if testCase.response != res {
			t.Fatalf("Wrong, expected: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
		}
	}

}
