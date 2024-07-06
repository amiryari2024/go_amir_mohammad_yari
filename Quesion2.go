package main

import (
	"fmt"
)

func calculateGrade(score int) string {
	if score < 0 || score > 100 {
		return "Invalid score"
	} else if score >= 85 {
		return "A"
	} else if score >= 70 {
		return "B"
	} else if score >= 55 {
		return "C"
	} else if score >= 40 {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	var score int
	fmt.Print("Enter the score: ")
	fmt.Scanln(&score)

	grade := calculateGrade(score)
	fmt.Printf("The grade for score %d is: %s\n", score, grade)
}
