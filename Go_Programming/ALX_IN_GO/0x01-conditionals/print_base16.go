package main

import (
	"fmt"
)

func PrintHexadecimals() {
	/* Function that prints all the numbers of base 16 in lowercase.
	followed by a new line.
	*/
	for i := 0; i < 16; i++ {
		fmt.Printf("%x", i)
	}
	fmt.Println()
}
