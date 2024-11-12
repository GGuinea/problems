package main

import (
	"testing"
)

func TestShouldWorksForTestInputPart1(t *testing.T) {
	testInput := "())"
	if res := part1Solution(testInput); -1 != res {
		t.Fatal("Returned : ", res)
	}
}

func TestShouldWorkForProperInputPart1(t *testing.T) {
	if res := handlePart1("input"); 232 != res {
		t.Fatal("Returned : ", res)
	}
}

func TestShouldWorksForTestInputPart2(t *testing.T) {
	if res := part2Solution("()())"); 5 != res {
		t.Fatal("Returned : ", res)
	}
}

func TestShouldWorkForProperInputPart2(t *testing.T) {
	if res := handlePart2("input"); 232 != res {
		t.Fatal("Returned : ", res)
	}
}
