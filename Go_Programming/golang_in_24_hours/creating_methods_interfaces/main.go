package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Using Methods
type Movie struct {
	Name   string
	Rating float64
}

// Declaring a Method
func (m *Movie) summary() string {
	r := strconv.FormatFloat(m.Rating, 'f', 1, 64)
	return m.Name + ", " + r
}

// Creating Method Sets
type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}
func (s *Sphere) Volume() float64 {
	radiusCubed := s.Radius * s.Radius * s.Radius
	return (float64(4) / float64(3)) * math.Pi * radiusCubed
}

// Working with Methods and Pointers
type Triangle struct {
	base   float64
	height float64
}

// Passing a Pointer Reference to a Method
func (t *Triangle) area() float64 {
	return 0.5 * (t.base * t.height)
}

// Passing a Value Reference to a Method
func (t Triangle) changeBase(f float64) {
	t.base = f
	return
}

// Golang Interfaces

type Robot interface {
	PowerOn() error
}

type T850 struct {
	Name string
}

func (a *T850) PowerOn() error {
	return nil
}

type R2D2 struct {
	Broken bool
}

func (r *R2D2) PowerOn() error {
	if r.Broken {
		return errors.New("R2D2 is broken")
	} else {
		return nil
	}
}

func Boot(r Robot) error {
	return r.PowerOn()
}
func main() {
	// Calling a Method
	m := Movie{
		Name:   "Spiderman",
		Rating: 3.2,
	}
	fmt.Println(m.summary())

	s := Sphere{
		Radius: 5,
	}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())

	t := Triangle{base: 3, height: 1}
	fmt.Println(t.area())
	t.changeBase(4)
	fmt.Println(t.base)

	t2 := T850{
		Name: "The Terminator",
	}
	r := R2D2{
		Broken: true,
	}
	err := Boot(&r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}
	err = Boot(&t2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Robot is powered on!")
	}
}
