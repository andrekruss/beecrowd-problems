package main

import "fmt"

func main() {

	// Defining variables
	var a int
	var b int
	var soma int

	// reading user input
	fmt.Scan(&a)
	fmt.Scan(&b)

	// Calculating and Outputting sum
	soma = a + b
	fmt.Printf("SOMA = %v\n", soma)
}
