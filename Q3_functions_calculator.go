package main

import "fmt"

// Function to add two numbers
func add(a int, b int) int {
	return a + b
}

// Function to subtract
func subtract(a int, b int) int {
	return a - b
}

// Function to multiply
func multiply(a int, b int) int {
	return a * b
}

// Function to divide
func divide(a int, b int) float64 {
	// convert to float64 to avoid integer division
	return float64(a) / float64(b)
}

func main() {
	a, b := 10, 5

	fmt.Println("Addition:", add(a, b))
	fmt.Println("Subtraction:", subtract(a, b))
	fmt.Println("Multiplication:", multiply(a, b))
	fmt.Println("Division:", divide(a, b))
}
