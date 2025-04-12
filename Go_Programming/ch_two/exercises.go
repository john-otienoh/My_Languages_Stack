package main

import (
	"fmt"
)

const value = 10

func main() {
	exercise_one()
	exercise_two()
	exercise_three()
}

func exercise_one() {
	/*
		Function that declares an integer variable called i with the value 20.
		Assign i to a floating-point variable named f.
		Print out i and f.
	*/
	var i int = 20
	var f float64 = float64(i)
	fmt.Println(i, f)
}

func exercise_two() {
	/*
		Function that declares a constant called value that can be assigned to both an integer and a floating-point variable.
		Assign it to an integer called i and a floating-point variable called f.
		Print out i and f..
	*/
	i := value
	var f float64 = float64(value)
	fmt.Println(f, i)
}

func exercise_three() {
	/*
		Function with three variables, one named b of type byte, one named smallI of type int32, and one named bigI of type uint64.
		Assign each variable the maximum legal value for its type; then add 1 to each variable.
		Print out their values.
	*/

	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615
	b += 1
	smallI += 1
	bigI += 1
	fmt.Println(b, smallI, bigI)
}
