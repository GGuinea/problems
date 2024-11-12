package main

import (
	"testing"
)

func TestShouldWorksForTestInput(t *testing.T) {
	if res := part1("test_input"); -1 != res {
		t.Fatal("Returned : ", res)
	}
}

func TestShouldWorkForProperInput(t *testing.T) {
	if res := part1("input"); 232 != res {
		t.Fatal("Returned : ", res)
	}
}
