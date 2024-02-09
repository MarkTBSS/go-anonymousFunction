package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	// Using normal function
	result := add(3, 4)
	fmt.Println("Result from normal function:", result)

	// Anonymous function assigned to a variable
	subtract := func(a, b int) int {
		return a - b
	}

	// Using anonymous function
	result = subtract(7, 2)
	fmt.Println("Result from anonymous function:", result)

	// Passing anonymous function as an argument to another function
	funcResult := func(a, b int) int {
		return a * b
	}(5, 6) // Immediately invoked
	fmt.Println("Result from immediately invoked anonymous function:", funcResult)
}
