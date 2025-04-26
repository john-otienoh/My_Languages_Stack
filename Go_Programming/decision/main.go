package main

import "fmt"

func main() {
	var (
		a = 5
		b = 8
	)
	if a > b {
		fmt.Printf("Truely %d is greater than %d\n", a, b)
	} else if a == b {
		fmt.Printf("Truly %d is equal to %d\n", a, b)
	} else {
		fmt.Printf("Truely %d is greater than %d\n", b, a)
	}
}
