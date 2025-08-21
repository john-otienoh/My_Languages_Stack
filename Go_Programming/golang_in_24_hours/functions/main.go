package main

import "fmt"

// simple function
func add(x int, y int) int {
	return x + y
}

// Returning a Single Result
func isEven(x int) bool {
	return x%2 == 0
}

// Returning Multiple Values
func getPrize() (int, string) {
	i := 2
	s := "goldfish"
	return i, s
}

// Defining Variadic Functions
func sumNumbers(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Using Named Return Values
func sayHi() (x, y string) {
	x = "hello"
	y = "world"
	return
}

// Using Recursive Functions
func feedMe(portion, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("I'm full! I've eaten %d\n", eaten)
		return eaten
	}
	fmt.Printf("I'm still hungry! I've eaten %d\n", eaten)

	return feedMe(portion, eaten)
}

// Passing Functions as Values
var fn = func() {
	fmt.Println("Passing function as values")
}

// Passing a Function as an Argument
func anotherFunction(f func() string) string {
	return f()
}
func main() {
	fmt.Println(sayHi())
	fmt.Println(add(2, 3))
	fmt.Println(isEven(3))
	fmt.Println(feedMe(1, 0))

	fn()
	fn := func() string {
		return "Passing function as arguments"
	}
	fmt.Println(anotherFunction(fn))

	quantity, prize := getPrize()
	fmt.Printf("You won %v %v\n", quantity, prize)

	result := sumNumbers(1, 2, 3, 4)
	fmt.Printf("The result is %v\n", result)
}
