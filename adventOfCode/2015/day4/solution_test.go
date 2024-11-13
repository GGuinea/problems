package solution

import "testing"

func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
	cases := []struct {
		input    string
		response int
	}{
		{
			input:    "abcdef",
			response: 609043,
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
	res := part1("yzbqklnj")
	if res != 282749 {
		t.Fatal("Got:", res)
	}
}

func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
	res := part2("yzbqklnj")
	if res != 9962624 {
		t.Fatal("Got:", res)
	}
}
