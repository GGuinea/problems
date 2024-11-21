package solution

import (
	"testing"
)

func TestShouldReturnProperValueForTestInputPart1(t *testing.T) {
	cases := []struct {
		input    string
		response string
	}{
		{
			input:    "abcdefgh",
			response: "abcdffaa",
		},
		{
			input:    "ghijklmn",
			response: "ghjaabcc",
		},
	}

	for _, testCase := range cases {
		res := handlePart1(testCase.input)
		if testCase.response != res {
			t.Fatalf("Wrong, expected: %v, actual: %v; input: %v", testCase.response, res, testCase.input)
		}
	}

}

func TestAtLeastThreelEttersIncreasedIndRow(t *testing.T) {
	res := atLeastThreeLettersIncreasedInRow("abc")
	if res != true {
		t.Fatalf("Wrong")
	}

	res = atLeastThreeLettersIncreasedInRow("yza")
	if res != true {
		t.Fatalf("Wrong")
	}

	res = atLeastThreeLettersIncreasedInRow("abcdffaa")
	if res != true {
		t.Fatalf("wrong")
	}
}

func TestNotContainsForbiddenRule(t *testing.T) {
	res := notContainsForbidden("abc")
	if res != true {
		t.Fatalf("wrong")
	}
	res = notContainsForbidden("abo")
	if res != false {
		t.Fatalf("wrong")
	}

	res = notContainsForbidden("iol")
	if res != false {
		t.Fatalf("wrong")
	}

	res = notContainsForbidden("ghjaabcc")
	if res != true {
		t.Fatalf("wrong")
	}

	res = notContainsForbidden("abcdffaa")
	if res != true {
		t.Fatalf("wrong")
	}
}

func TestContainsTwoDifferentPairs(t *testing.T) {
	res := containsTwoDifferentPairs("aba")
	if res != false {
		t.Fatalf("wrong")
	}

	res = containsTwoDifferentPairs("aaa")
	if res != false {
		t.Fatalf("wrong")
	}

	res = containsTwoDifferentPairs("aabb")
	if res != true {
		t.Fatalf("wrong")
	}

	res = containsTwoDifferentPairs("ghjaabcc")
	if res != true {
		t.Fatalf("wrong")
	}

	res = containsTwoDifferentPairs("abcdffaa")
	if res != true {
		t.Fatalf("wrong")
	}
}

func TestVaildateRules(t *testing.T) {
	res := validateRules("ghjaabcc", []func(string) bool{atLeastThreeLettersIncreasedInRow, containsTwoDifferentPairs, notContainsForbidden})
	if res != true {
		t.Fatalf("wrong")
	}
}

func TestShouldReturnProperValueForProperInputPart1(t *testing.T) {
	res := part1("hxbxwxba")
	if res != "hxbxxyzz" {
		t.Fatal("Got:", res)
	}
}

func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
	res := part1("hxbxwxba")
	if res != "hxbxxyzz" {
		t.Fatal("Got1:", res)
	}
	res2 := part1(res)
	if res2 != "" {
		t.Fatal("Got2:", res2)
	}
}

// func TestShouldReturnProperValueForProperInputPart2(t *testing.T) {
// 	res := part2()
// 	if res != 6989950 {
// 		t.Fatal("Got:", res)
// 	}
// }
//
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
