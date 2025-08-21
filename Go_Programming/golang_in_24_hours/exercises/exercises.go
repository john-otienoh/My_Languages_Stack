package main

import (
	"fmt"
	"strconv"
)

func tempConvertor(c float64) float64 {
	// function to convert Celsius to Kelvin
	return c + 273.15
}

func iCallMyself(i int) int {
	// function that calls itself ten times, then exits.
	if i >= 10 {
		fmt.Printf("I've called myself %d times now exiting.\n", i)
		return i
	}
	fmt.Printf("%d time calling myself\n", i)
	return iCallMyself(i + 1)
}

func takeLessReturnMore(x, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func main() {
	// function that converts a string type variable to an int type variable and back again.
	var i int = 22
	var s string = "1234"

	fmt.Printf("Typeof variable b(%d) is: %T\n", i, i)
	fmt.Printf("Typeof variable s(%s) is: %T\n", s, s)

	var s2 string = strconv.Itoa(i)
	fmt.Printf("Typeof variable s(%s) is: %T\n", s2, s2)

	i2, err := strconv.Atoi(s)
	fmt.Println(i2, err)

	fmt.Println(tempConvertor(1))

	fmt.Println(iCallMyself(0))

	add, sub, mul := takeLessReturnMore(5, 3)
	fmt.Printf("Addition of %d and %d is equal to %d\n", 5, 3, add)
	fmt.Printf("Subtraction of %d and %d is equal to %d\n", 5, 3, sub)
	fmt.Printf("Multiplication of %d and %d is equal to %d\n", 5, 3, mul)
}
