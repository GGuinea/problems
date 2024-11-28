package fib

import (
	"testing"
)

func BenchmarkFibWithoutDP(t *testing.B)  {
	fib(40)
}

func BenchmarkFibWithDP(t *testing.B)  {
	memo := make(map[int]int)
	fibDP(40, memo)
}
