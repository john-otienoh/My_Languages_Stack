package main

import (
	"fmt"
)

func Concat() {
	/*
		Function to print Welcome to Holberton School!
		You have to use the variables str1 and str2 in your new line of code
	*/
	str1, str2 := "Holberton", "School"
	str3 := str1 + " " + str2
	fmt.Printf("Welcome to %s !\n", str3)
}
