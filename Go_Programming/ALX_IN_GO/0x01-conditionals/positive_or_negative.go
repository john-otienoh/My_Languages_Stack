package main

import (
	"fmt"
	"math/rand"
)

func PositiveOrNegative() {
	min, max := -1000, 1000
	n := rand.Intn(max-min+1) + min
	if n > 0 {
		fmt.Printf("%d is positive\n", n)
	} else if n < 0 {
		fmt.Printf("%d is negative\n", n)
	} else {
		fmt.Printf("%d is zero\n", n)
	}
}
