package fib

func fib(n int) int {
	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fibDP(n int, memo map[int]int) int {
	res, found := memo[n]
	if found {
		return res
	}

	if n <= 2 {
		return 1
	}

	fibOfN := fibDP(n-1, memo) + fibDP(n-2, memo)
	memo[n] = fibOfN

	return fibOfN
}
