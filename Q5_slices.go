package main

import "fmt"

func main() {
	// Create a slice
	languages := []string{"Go", "Python", "Java"}

	// Append an element
	languages = append(languages, "Rust")

	// Print all elements
	fmt.Println("Programming Languages:")
	for index, lang := range languages {
		fmt.Println("Index", index, "Value", lang)
	}
}
