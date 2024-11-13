package solution

import "testing"

func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
	cases := []struct {
		input    string
		response int
	}{
		{
			input:    ">",
			response: 2,
		},
		{
			input:    "^>v<",
			response: 4,
		},
		{
			input:    "^v^v^v^v^v",
			response: 2,
		},
	}

	for _, testCase := range cases {
		res := handlePart1(testCase.input)
		if testCase.response != res {
			t.Fatalf("Wrong, expeced: %v, actual: %v", testCase.response, res)
		}
	}

}
func TestShouldReturnProperValueForProperInputPart1(t *testing.T) {
	res := part1()
	if res != 0 {
		t.Fatal("Got:", res)
	}
}

func TestShouldReturnProperValueForTestPart1(t *testing.T) {
	cases := []struct {
		input    string
		response int
	}{
		{
			input:    "^v",
			response: 3,
		},
		{
			input:    "^>v<",
			response: 3,
		},
		{
			input:    "^v^v^v^v^v",
			response: 11,
		},
	}

	for _, testCase := range cases {
		res := handlePart2(testCase.input)
		if testCase.response != res {
			t.Fatalf("Wrong, expeced: %v, actual: %v", testCase.response, res)
		}
	}

}

func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
	res := part2()
	if res != 0 {
		t.Fatal("Got: ", res)
	}
}
