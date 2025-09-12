package main

import (
	"errors"
	"fmt"
	"os"
)

// Understanding the Error Type
type error interface {
	Error() string
}

func Half(numberToHalf int) (int, error) {
	if numberToHalf%2 != 0 {
		return -1, fmt.Errorf("Cannot half %v", numberToHalf)
	}
	return numberToHalf / 2, nil
}

func main() {
	// Error Handling when Reading a File
	file, err := os.ReadFile("foo.txt")
	if err != nil {
		fmt.Println(err)
		// return
	} else {
		fmt.Println(file)
	}

	// Creating errors
	err1 := errors.New("Something went wrong")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println("hello there")
	}

	// Formatting Errors
	name, role := "Richard Jupp", "Drummer"
	err2 := fmt.Errorf("The %v %v quit", role, name)
	if err2 != nil {
		fmt.Println(err2)
	}

	// Returning an Error from a Function
	n, err3 := Half(19)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(n)

	// Using panic to halt Execution
	fmt.Println("This is executed")
	panic("Oh no. I can do no more. Goodbye.")
	fmt.Println("This is not executed")
}
