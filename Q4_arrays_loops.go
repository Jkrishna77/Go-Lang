package main

import "fmt"

func main() {
	// Declare array
	numbers := [5]int{10, 20, 30, 40, 50}

	// Using normal for loop
	fmt.Println("Using normal for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index", i, "Value", numbers[i])
	}

	// Using range loop
	fmt.Println("\nUsing range loop:")
	for index, value := range numbers {
		fmt.Println("Index", index, "Value", value)
	}
}
