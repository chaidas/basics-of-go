package main

import "fmt"

func functionPlayground() {
	fmt.Println("Adding two numbers: ", add(10, 20))
}

func add(a, b int) int {
	return a + b
}

// Example of a function returning _two_ values
func addAndSubtract(a int, b int) (int, int) {
	return a + b, a - b
}

// Passing a reference
