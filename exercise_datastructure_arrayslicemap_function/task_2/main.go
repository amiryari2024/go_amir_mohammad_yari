package main

import (
	"fmt"
)

func main() {
	fmt.Println(primeSum([]int{1, 2, 4, 5, 12, 19, 30})) // 26
	fmt.Println(primeSum([]int{45, 17, 44, 67, 11}))     // 95
	fmt.Println(primeSum([]int{15, 12, 9, 10}))          // 0
}

func primeSum(numbers []int) int {
	sum := 0

	// Helper function to check if a number is prime
	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		if n <= 3 {
			return true
		}
		if n%2 == 0 || n%3 == 0 {
			return false
		}
		i := 5
		for i*i <= n {
			if n%i == 0 || n%(i+2) == 0 {
				return false
			}
			i += 6
		}
		return true
	}

	// Iterate through the numbers in the slice
	for _, num := range numbers {
		// Check if the number is prime using the helper function
		if isPrime(num) {
			sum += num
		}
	}

	return sum
}
