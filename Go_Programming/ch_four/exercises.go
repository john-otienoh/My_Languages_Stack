package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ex_one()
	ex_two()
	ex_three()
}

func ex_one() {
	/*
		For loop that puts 100 random numbers between 0 and 100 into an int slice.
	*/
	randomInt := []int{}
	for i := 0; i < 101; i++ {
		randomInt = append(randomInt, rand.Intn(100))
	}
	fmt.Println(randomInt)
}

func ex_two() {
	/*
		Loop over the slice you created in exercise 1. For each value in the slice, apply the following rules:
		a. If the value is divisible by 2, print “Two!”
		b. If the value is divisible by 3, print “Three!”
		c. If the value is divisible by 2 and 3, print “Six!”. Don’t print anything else.
		d. Otherwise, print “Never mind”.
	*/
	randomInt := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		randomInt = append(randomInt, rand.Intn(100))
	}
	for _, v := range randomInt {
		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six!")
		} else if v%2 == 0 {
			fmt.Println("Two!")
		} else if v%2 == 0 {
			fmt.Println("Two!")
		} else {
			fmt.Println("Never Mind")
		}
	}

}

func ex_three() {
	/*
		Declare an int variable called total.
		Write a for loop that uses a variable named i to iterate from 0 (inclusive) to 10 (exclusive).
		total := total + i
		fmt.Println(total)
		After the for loop, print out the value of total.
		What is printed out? What is the likely bug in this code?
	*/
	var total int
	for i := 0; i < 10; i++ {
		total = total + 1
		fmt.Println(total)
	}
	fmt.Println("Overall total: ", total)
}
