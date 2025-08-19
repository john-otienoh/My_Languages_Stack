package main

import "fmt"

// Declaring Variables with Constants
const PIE float64 = 3.14

func showMemoryAddressValues(x int) {
	// Passing Variables as Values
	fmt.Println(&x)
	return
}
func showMemoryAddressPointers(x *int) {
	// Passing Variables as Pointers
	fmt.Println(x)
	return
}

func showValue(x *int) {
	fmt.Println(*x)
	return
}
func main() {
	// Declaring a String Variable
	var s string = "foo"
	var s3 = "foo3"

	// Shorthand Variable Declaration
	var s1, t string = "foo1", "bar"

	// Assigning a Value after Initializing a Variable
	var i int
	i = 1

	// Shorthand Variable Declaration for Different Types
	var (
		s2 string = "foo2"
		i2 int    = 67
	)

	// Short Variable Declaration
	is_correct := true

	// Understanding Variables and Zero Values
	var i3 int
	var f float64
	var b bool
	var s4 string

	// Using Pointers

	fmt.Println(&i)
	showMemoryAddressValues(i)
	fmt.Println()

	fmt.Println(&i2)
	showMemoryAddressPointers(&i2)
	fmt.Println()

	fmt.Println(i)
	showValue(&i)
	fmt.Println()

	fmt.Printf("%v %v %v %q\n", i3, f, b, s4)
	fmt.Println(s, i, s1, t, s2, i2, is_correct, s3)

}
