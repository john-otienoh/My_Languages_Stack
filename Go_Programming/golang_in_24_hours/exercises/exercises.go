package main

import (
	"fmt"
	"strconv"
)

func main() {
	// function that converts a string type variable to an int type variable and back again.
	var i int = 22
	var s string = "1234.746"
	fmt.Print(s)
	s = "hello"
	fmt.Print(s)

	fmt.Printf("Typeof variable b(%d) is: %T\n", i, i)
	fmt.Printf("Typeof variable s(%s) is: %T\n", s, s)

	var s2 string = strconv.Itoa(i)
	fmt.Printf("Typeof variable s(%s) is: %T\n", s2, s2)

	i2, err := strconv.Atoi(s)
	fmt.Println(i2, err)
}
