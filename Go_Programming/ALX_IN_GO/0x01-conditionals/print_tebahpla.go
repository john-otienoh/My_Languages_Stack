package main

import (
	"fmt"
)

func ReversedAlphabets() {
	/*
		Function that prints the lowercase alphabet in reversed
		followed by a new line.
	*/
	for i := 122; i >= 97; i-- {
		fmt.Printf("%c", rune(i))
	}

	fmt.Println()
}
