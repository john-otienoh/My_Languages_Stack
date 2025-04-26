package main

import (
	"fmt"
	"math"
)

func main() {
	foo()
	circle_area(7)
	calculate(3, 4, 5)
	result := multiplier(3, 4, 5)
	r1, r2, r3 := compute(3, 4, 5, "John")
	r4 := add(3, 5, 5)
	fmt.Printf("Results = %.2f \n", result)
	fmt.Printf("Multiplication results = %.2f\nAddition results = %.2f\n%s\n", r1, r2, r3)
	fmt.Printf("Addition Results = %d \n", r4)
	closure_func("Testing")
	fib := fibonacci(15)
	fmt.Printf("fibonacci() = %d \n", fib)
}

// Creating A Simple Function
func foo() {
	fmt.Println("foo() was called")
}

// Function with Parameters
func circle_area(r float64) {
	area := math.Pi * math.Pow(r, 2)
	fmt.Printf("Circle Area (r=%.2f) = %.2f\n", r, area)
}

// Function with many parameters
func calculate(a, b, c float64) {
	result := a * b * c
	fmt.Printf("Results if a=%.2f, b=%.2f and c=%.2f = %.2f \n", a, b, c, result)
}

// Function with Returning Value
func multiplier(a, b, c float64) float64 {
	result := a * b * c
	return result
}

// Function with Multiple Returning Values
func compute(a, b, c float64, name string) (float64, float64, string) {
	res1 := a * b * c
	res2 := a + b + c
	res3 := res1 + res2
	newName := fmt.Sprintf("%s's value = %.2f", name, res3)

	return res1, res2, newName

}

// Function with Multiple Parameters and Returning Value
func add(numbers ...int) int {
	result := 0
	for _, i := range numbers {
		result += i
	}
	return result
}

// Closure Functions
func closure_func(name string) {
	hoo := func(a, b int) {
		result := a * b
		fmt.Printf("Hoo() = %d\n", result)
	}
	joo := func(a, b int) int {
		return a*b + a
	}
	fmt.Printf("closure_func(%s) was called\n", name)
	hoo(2, 3)
	val := joo(5, 8)
	fmt.Printf("val from joo() = %d \n", val)
}

// Recursion Function
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
