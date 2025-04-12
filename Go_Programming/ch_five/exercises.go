package main

import (
	"errors"
	"fmt"
	"strconv"
)

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		if p2 == 0 && op == "/" {
			fmt.Println("division by zero")
			continue
		}
		result, err := opFunc(p1, p2)
		fmt.Println(result)
	}
	// Exercise 3
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
func add(i int, j int) (int, error) { return i + j, nil }
func sub(i int, j int) (int, error) { return i - j, nil }
func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

// func fileLen(s string) (int, error) {
// 	/*
// 		function called fileLen that has an input parameter of type string and returns an int and an error.
// 		The function takes in a filename and returns the number of bytes in the file.
// 		If there is an error reading the file, return the error.
// 		Use defer to make sure the file is closed properly.
// 	*/
// 	return 0, nil
// }

func prefixer(s string) func(string) string {
	/*
		function called prefixer that has an input parameter of type string
		and returns a function that has an input parameter of type string and returns a string.
		The returned function should prefix its input with the string passed into prefixer.
		Use the following main function to test prefixer:
	*/
	return func(s2 string) string {
		return s + " " + s2
	}
}
