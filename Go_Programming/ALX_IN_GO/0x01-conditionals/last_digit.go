package main

import (
	"fmt"
	"math/rand"
)

func LastDigit() {
	/*
		Function that will assign a random number to the variable n each time it is executed.
		The output of the program should be:
		The string "Last digit of", followed by
		n, followed by
		the string is, followed by
		if the last digit of n is greater than 5: the string and is greater than 5
		if the last digit of n is 0: the string and is 0
		if the last digit of n is less than 6 and not 0: the string and is less than 6 and not 0
		followed by a new line
	*/
	min, max := -1000, 1000
	n := rand.Intn(max-min+1) + min
	lastDigit, prefix := n%10, "Last digit of"

	if lastDigit > 5 {
		fmt.Printf("%s %d is %d and is greater than 5\n", prefix, n, lastDigit)
	} else if lastDigit != 0 && lastDigit < 6 {
		fmt.Printf("%s %d is %d and is less than 6 and not 0\n", prefix, n, lastDigit)
	} else {
		fmt.Printf("%s %d is %d and is 0\n", prefix, n, lastDigit)
	}
}
