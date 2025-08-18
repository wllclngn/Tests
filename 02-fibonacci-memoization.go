package main

import "fmt"

func refinedExpensiveFib(n int, cache map[int]int) int {
	if n < 2 {
		cache[n] = n
		return n
	}

	if _, ok := cache[n-1]; !ok {
		cache[n-1] = refinedExpensiveFib(n-1, cache)
	}

	if _, ok := cache[n-2]; !ok {
		cache[n-2] = refinedExpensiveFib(n-2, cache)
	}

	return cache[n-1] + cache[n-2]
}

func memoize(n int) int {
	cache := make(map[int]int)
	result := make([]int, n)

	for i := 1; i <= n; i++ {
		result[i-1] = refinedExpensiveFib(i, cache)
	}

	return result[n-1]
}

func main() {
	input := 100

	fmt.Println(memoize(input))
}
