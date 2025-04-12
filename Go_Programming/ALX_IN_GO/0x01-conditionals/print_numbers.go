package main

import (
	"fmt"
)

func PrintNumbers() {
	/* Function that  prints all single digit numbers of base 10 starting from 0
	followed by a new line.
	*/
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()
}
