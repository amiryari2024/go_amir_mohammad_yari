package main

import (
	"fmt"
	"sort"
)

func main() {
	data1 := []float64{1, 1, 2, 5, 6, 8, 12, 4, 5, 5, 5, 8, 9}
	fmt.Printf("sum: %.2f\n", sum(data1))       // 71.00
	fmt.Printf("mean: %.2f\n", mean(data1))     // 5.46
	fmt.Printf("median: %.2f\n", median(data1)) // 5.00
	fmt.Printf("mode: %.2f\n", mode(data1))     // 5.00

	data2 := []float64{6, 7, 1, 11, 8, 12, 6, 12, 1, 7, 8, 10, 14}
	fmt.Printf("sum: %.2f\n", sum(data2))       // 103.00
	fmt.Printf("mean: %.2f\n", mean(data2))     // 7.92
	fmt.Printf("median: %.2f\n", median(data2)) // 8.00
	fmt.Printf("mode: %.2f\n", mode(data2))     // 1.00
}

func sum(data []float64) float64 {
	total := 0.0
	for _, num := range data {
		total += num
	}
	return total
}

func mean(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}
	total := sum(data)
	return total / float64(len(data))
}

func median(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}
	// Sort the data slice
	sort.Float64s(data)

	// Determine the middle index
	mid := len(data) / 2

	// If the length of data is odd, return the middle element
	if len(data)%2 == 1 {
		return data[mid]
	}

	// If the length of data is even, return the average of the two middle elements
	return (data[mid-1] + data[mid]) / 2.0
}

func mode(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}

	// Create a map to count occurrences of each element
	count := make(map[float64]int)
	for _, num := range data {
		count[num]++
	}

	// Find the element with the maximum count
	var modeValue float64
	maxCount := 0
	for num, cnt := range count {
		if cnt > maxCount {
			maxCount = cnt
			modeValue = num
		}
	}

	return modeValue
}
