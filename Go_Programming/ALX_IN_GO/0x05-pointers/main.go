package main

import (
	"fmt"
)

func main() {
	fmt.Println("----------------------Exercise One---------------------")
	a := 402
	b := 90
	fmt.Println(a)
	// reset_to_98(&a)
	fmt.Println(a)
	fmt.Println("----------------------Exercise Two---------------------")
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Println("----------------------Exercise Three---------------------")
	word := "Hello John"
	fmt.Println(strlen(&word))
	fmt.Println("----------------------Exercise Two---------------------")
	reverse(&word)
	fmt.Println("----------------------Exercise Two---------------------")
	fmt.Println("----------------------Exercise Two---------------------")
	fmt.Println("----------------------Exercise Two---------------------")
	fmt.Println("----------------------Exercise Two---------------------")
}
