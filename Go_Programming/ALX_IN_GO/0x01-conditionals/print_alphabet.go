package main

import (
	"fmt"
)

func PrintAlphabet() {
	// Function that prints the alphabet in lowercase followed by a new line.
	for i := 97; i < 123; i++ {
		fmt.Printf("%c", rune(i))
	}
	fmt.Println()
}
