package main

import "fmt"

func main() {
	fmt.Println("test")
	arr := []int{4, 1, 5, 10, 2, 1}
	output := qs(arr, 0, len(arr)-1)
	fmt.Println(output)
}

func qs(input []int, lo, hi int) []int {
	if lo >= hi {
		return input
	}

	pivotIdx := partition(input, lo, hi)
	qs(input, lo, pivotIdx-1)
	qs(input, pivotIdx+1, hi)
	return input
}

func partition(input []int , lo, hi int) int {
	pivot := input[hi]
	idx := lo - 1
	for i := lo; i < hi; i++ {
		if input[i] <= pivot {
			idx++
			tmp := input[i]
			input[i] = input[idx]
			input[idx] = tmp
		}
	}
	idx++
	input[hi] = input[idx]
	input[idx] = pivot
	return idx
}
