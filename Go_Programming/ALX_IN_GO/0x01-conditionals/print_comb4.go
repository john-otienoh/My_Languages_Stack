package main

import (
	"fmt"
)

func PrintTripples() {
	// Function that  prints all possible combination of tripple-digit numbers
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k < 10; k++ {
				if i < j && j < k {
					if i < 7 && j < 9 && k <= 9 {
						fmt.Printf("%d%d%d, ", i, j, k)
					} else {
						fmt.Printf("%d%d%d", i, j, k)
					}
				}
			}
		}

	}
	fmt.Println()
}
