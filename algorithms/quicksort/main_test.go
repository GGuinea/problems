package main

import (
	"slices"
	"testing"
)

func TestShouldSort(t *testing.T) {
	arr := []int{4, 1, 5, 10, 2, 1}
	output := qs(arr, 0, len(arr)-1)
	if !slices.IsSorted(output) {
		t.Fatalf("Should be sorted! %v", output)
	}
}
