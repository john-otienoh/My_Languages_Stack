package main

import "fmt"

func main() {
	// Declaring Variables
	var name string = "John"
	var age int = 20
	var salary float64 = 59.0838
	/* Define multiple variables */
	var (
		name2   string  = "Jane"
		age2    int     = 21
		salary2 float64 = 8949.4984
	)
	// := syntax
	name3 := "Joe"
	age3 := 29
	salary3 := 46637.93837

	fmt.Printf("First person details:\tName:%s Age:%d Salary: %g\n", name, age, salary)
	fmt.Printf("Second person details:\tName:%s Age:%d Salary: %g\n", name2, age2, salary2)
	fmt.Printf("Third person details:\tName:%s Age:%d Salary: %g\n", name3, age3, salary3)

}
