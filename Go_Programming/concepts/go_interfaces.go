package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	circum() float64
}

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circum() float64 {
	return s.length * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circum() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfos(s shape) {
	fmt.Printf("The area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("The circumference of %T is: %0.2f \n", s, s.circum())
}
func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 5.6},
	}

	for _, v := range shapes {
		printShapeInfos(v)
		fmt.Println("--------")
	}
}
