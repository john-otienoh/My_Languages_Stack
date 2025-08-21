package main

import (
	"fmt"
)

func main() {
	// If, else...if and else
	b := 4
	if b > 5 {
		fmt.Println("b is greater than 5")
	} else if b < 5 {
		fmt.Println("b is less than 5")
	} else {
		fmt.Println("b is equal to 5")
	}

	// Switch case
	s := "bs"
	switch s {
	case "a":
		fmt.Println("The letter a")
	case "b":
		fmt.Println("The letter b")
	default:
		fmt.Println("Unrecognized letter!!!!!!")
	}

	// For loop
	i := 0
	for i < 3 {
		fmt.Println("i is", i)
		i++
	}
	// for Statement with init and post Statements
	for i := 0; i < 5; i++ {
		fmt.Println("i is", i)
	}

	// for Statements with a range Clause
	numbers := []int{1, 2, 3, 4}
	for k, v := range numbers {
		fmt.Println("The index of the loop is", k)
		fmt.Println("The value from the array is", v)
	}

	// defer Statements
	defer fmt.Println("I am the first defer statement")
	defer fmt.Println("I am the second defer statement")
	defer fmt.Println("I am the third defer statement")
	fmt.Println("Hello World!")
}
