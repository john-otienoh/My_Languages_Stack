package main

import (
	"fmt"
)

func PrintAlphabt() {
	/*
		Function that prints the alphabet in lowercase followed by a new line.
		Except q and e
	*/
	for i := 97; i < 123; i++ {
		letter := rune(i)
		if letter != 'e' && letter != 'q' {
			fmt.Printf("%c", letter)
		}
	}
	fmt.Println()
}
