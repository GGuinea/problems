package day1

import (
	"regexp"
	"testing"
)

func TestShouldReturnProperValuesForPart1TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{
			input: "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			res:   161,
		},
	}

	for _, tCase := range testCase {
		if res := handlePart1(tCase.input); res != tCase.res {
			t.Fatalf("Expected: %v, got: %v", tCase.res, res)
		}
	}
}

func TestShouldReturnProperValuesForPart1(t *testing.T) {
	res := part1()
	if res != 170778545 {
		t.Fatalf("Got: %v", res)
	}
}

func TestRegexpTest(t *testing.T) {
	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	res := r.FindAllString("mul(10,15)", -1)
	if res == nil {
		t.Fatalf("not found")
	}

	if res[0] != "mul(10,15)" {
		t.Fatalf("not found expected, found: %v", res)
	}
}

func TestRegexpTestDoDont(t *testing.T) {
	r, err := regexp.Compile(`do\(\)|don't\(\)|mul\(\d{1,3},\d{1,3}\)`)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	res := r.FindAllString("do()mul(10,15)", -1)
	if res == nil {
		t.Fatalf("not found")
	}

	if res[0] != "do()" && res[1] != "mul(10,15)" {
		t.Fatalf("not found expected, found: %v", res)
	}
}

func TestShouldReturnProperValuesForPart2TestInput(t *testing.T) {
	testCase := []struct {
		input string
		res   int
	}{

		{
			input: "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			res:   48,
		},
	}

	for _, tCase := range testCase {
		if res := handlePart2(tCase.input); res != tCase.res {
			t.Fatalf("Expected: %v, got: %v", tCase.res, res)
		}
	}
}

func TestShouldReturnProperValuesForPart2(t *testing.T) {
	res := part2()
	if res != 82868252 {
		t.Fatalf("Got: %v", res)
	}

}
