package main

import "fmt"

func main() {
	// Defining variables
	var A int
	var B int
	var X int

	// Reading User Input
	fmt.Scan(&A)
	fmt.Scan(&B)

	// Print X value
	X = A + B
	fmt.Printf("X = %v\n", X)
}
