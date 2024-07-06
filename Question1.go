package main

import (
	"fmt"
)

func main() {
	// Declare the required variables
	const phi = 3.14
	var r, t float64

	// Prompt the user to enter the radius and height
	fmt.Print(" 7 ")
	fmt.Scanln(&r)
	fmt.Print("5 ")
	fmt.Scanln(&t)

	// Calculate the volume of the cylinder
	v := phi * r * r * t

	// Display the output
	fmt.Printf("The volume of the cylinder is: %.2f\n", v)
}
