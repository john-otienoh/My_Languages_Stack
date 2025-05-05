package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(&x)
	var num *int
	val := new(int)
	num = new(int)
	*num = x
	// set value
	val = &x
	// set address
	fmt.Println("===pointer num===")
	fmt.Println(*num)
	fmt.Println(num)
	fmt.Println("===pointer val===")
	fmt.Println(*val)
	fmt.Println(val)
}
