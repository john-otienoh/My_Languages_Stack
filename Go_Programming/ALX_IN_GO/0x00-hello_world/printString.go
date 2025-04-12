package main

import (
	"fmt"
	"strings"
)

func PrintStringAndChars() {
	/*
		Function to print 3 times a string stored in the variable str,
		followed by its first 9 characters.
	*/
	str := "Holberton School "
	fmt.Println(strings.Repeat(str, 3))
	fmt.Println(str[:9])
}
