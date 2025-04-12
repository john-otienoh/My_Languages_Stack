package main

import (
	"fmt"
)

func PrintFloat() {
	/*
		Function to print the float stored in the variable number
		with a precision of 2 digits.
	*/
	number := 3.14159
	fmt.Printf("Float: %.2f\n", number)
}
