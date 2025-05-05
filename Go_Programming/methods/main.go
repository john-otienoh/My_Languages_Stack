package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y int
	r    float64
}

func main() {
	circle := Circle{10, 5, 2.8}
	circle.display()
	fmt.Printf("Area = %2.f\n", circle.area())
	fmt.Printf("Circumference = %2.f\n", circle.circumference())
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c Circle) circumference() float64 {
	return 2 * math.Pi * c.r
}

func (c Circle) display() {
	fmt.Printf("x=%d,y=%d,r=%.2f\n", c.x, c.y, c.r)
}
