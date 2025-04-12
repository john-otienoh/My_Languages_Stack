package main

import (
	"fmt"
)

func PrintSingle() {
	// Function that  prints all possible combination of single-digit numbers
	for i := 0; i < 10; i++ {
		if i < 9 {
			fmt.Printf("%d, ", i)
		} else {
			fmt.Printf("%d", i)
		}
	}
	fmt.Println()
}
