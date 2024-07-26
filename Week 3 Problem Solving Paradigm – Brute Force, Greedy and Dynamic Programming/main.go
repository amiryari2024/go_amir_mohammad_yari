package main

import (
	"fmt"
	"strconv"
)

// Convert a given integer into a binary number
func convertToBinary(n int) []string {
	result := []string{}
	for i := 0; i <= n; i++ {
		binary := strconv.FormatInt(int64(i), 2)
		result = append(result, binary)
	}
	return result
}

// Generate Pascal's Triangle
func generatePascalTriangle(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	triangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	return triangle
}

func main() {
	// Convert to binary example
	n := 3
	binaryNumbers := convertToBinary(n)
	fmt.Println(binaryNumbers) // Output: [0, 1, 10, 11]

	// Generate Pascal's Triangle example
	numRows := 5
	triangle := generatePascalTriangle(numRows)
	for _, row := range triangle {
		fmt.Println(row)
	}
}
