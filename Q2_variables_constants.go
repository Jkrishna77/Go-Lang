package main

import (
	"fmt"
)

func main() {
	var name string = "Jk"
	var age int = 23

	num := 07

	const pi = 3.14

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("num:", num)

	fmt.Println("pi:", pi)
	// pi += pi // This will cause an error because pi is a constant
}
