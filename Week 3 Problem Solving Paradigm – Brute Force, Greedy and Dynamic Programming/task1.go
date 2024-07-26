package main

import "fmt"

// Helper function to calculate Fibonacci with memoization
func fibonacciHelper(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}

	if val, ok := memo[n]; ok {
		return val
	}

	memo[n] = fibonacciHelper(n-1, memo) + fibonacciHelper(n-2, memo)
	return memo[n]
}

// Fibonacci function that initializes the memo map
func fibonacci(number int) int {
	memo := make(map[int]int)
	return fibonacciHelper(number, memo)
}

func main() {
	fmt.Println(fibonacci(0))  // 0
	fmt.Println(fibonacci(2))  // 1
	fmt.Println(fibonacci(9))  // 34
	fmt.Println(fibonacci(10)) // 55
	fmt.Println(fibonacci(12)) // 144
}
