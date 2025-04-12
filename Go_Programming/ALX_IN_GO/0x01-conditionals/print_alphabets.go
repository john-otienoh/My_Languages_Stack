package main

import (
	"fmt"
)

func PrintAlphabets() {
	/*
		Function that prints the alphabet in
		lowercase, and then in uppercase,
		followed by a new line.
	*/
	for i := 97; i < 123; i++ {
		fmt.Printf("%c", rune(i))
	}
	for i := 65; i < 91; i++ {
		fmt.Printf("%c", rune(i))
	}
	fmt.Println()
}
