package main

import "fmt"

type Student struct {
	Name         string
	MathScore    int
	ScienceScore int
	EnglishScore int
}

func main() {
	students := []Student{
		{"jamie", 67, 88, 90},
		{"michael", 77, 77, 90},
		{"aziz", 100, 65, 88},
		{"jamal", 50, 80, 75},
		{"eric", 70, 80, 65},
		{"john", 80, 76, 68},
	}

	getInfo(students)
	// output:
	// best student in math: aziz with 100
	// best student in science: jamie with 88
	// best student in english: jamie with 90
	// best student in average: aziz with 84
}

func getInfo(students []Student) {
	var bestMath, bestScience, bestEnglish Student
	var highestMath, highestScience, highestEnglish int

	var bestAverage Student
	var highestAverage float64

	for _, student := range students {
		if student.MathScore > highestMath {
			highestMath = student.MathScore
			bestMath = student
		}
		if student.ScienceScore > highestScience {
			highestScience = student.ScienceScore
			bestScience = student
		}
		if student.EnglishScore > highestEnglish {
			highestEnglish = student.EnglishScore
			bestEnglish = student
		}

		average := float64(student.MathScore+student.ScienceScore+student.EnglishScore) / 3
		if average > highestAverage {
			highestAverage = average
			bestAverage = student
		}
	}

	fmt.Printf("best student in math: %s with %d\n", bestMath.Name, bestMath.MathScore)
	fmt.Printf("best student in science: %s with %d\n", bestScience.Name, bestScience.ScienceScore)
	fmt.Printf("best student in english: %s with %d\n", bestEnglish.Name, bestEnglish.EnglishScore)
	fmt.Printf("best student in average: %s with %.0f\n", bestAverage.Name, highestAverage)
}
