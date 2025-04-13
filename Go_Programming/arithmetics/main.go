package main

import "fmt"

func main() {
	var (
		a int = 10
		b int = 5
	)

	// Addition
	fmt.Printf("%d + %d = %d\n", a, b, (a + b))

	// Subtraction
	fmt.Printf("%d - %d = %d\n", a, b, (a - b))

	// Multiplication
	fmt.Printf("%d x %d = %d\n", a, b, (a * b))

	// Division
	fmt.Printf("%d / %d = %d\n", a, b, (a / b))

}
