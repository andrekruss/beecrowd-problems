package main

import "fmt"

func main() {
	// defining variables
	const pi = 3.14159
	var r float64

	// reading user input
	fmt.Scan(&r)

	// defining and calculating area
	a := pi * r * r

	// Outputing area with 4 decimal places
	fmt.Printf("A=%.4f\n", a)
}
