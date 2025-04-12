package main

import (
	"fmt"
)

func PrintDoubles() {
	// Function that  prints all possible combination of double-digit numbers
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i < j {
				if i < 8 && j <= 9 {
					fmt.Printf("%d%d, ", i, j)
				} else {
					fmt.Printf("%d%d", i, j)
				}
			}
		}
	}
	fmt.Println()
}
