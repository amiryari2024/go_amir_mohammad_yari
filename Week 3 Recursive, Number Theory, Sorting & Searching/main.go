package main

import "fmt"

// Function to calculate the Fibonacci number at position n using recursion
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Function to sum all Fibonacci numbers up to the nth Fibonacci number
func fibX(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += fibonacci(i)
	}
	return sum
}

func main() {
	fmt.Println(fibX(5))  // Output: 12
	fmt.Println(fibX(10)) // Output: 143
}
