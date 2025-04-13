package main

import (
	"fmt"
	"mymath"
)

func main() {
	hello()
	sandbox()
	pkg()
	imports()
	fmt.Println(add(42, 13))
	xs := []float64{1, 2, 3, 4}
	avg := mymath.Average(xs)
	fmt.Println((avg))
}
