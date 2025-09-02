package main

import (
	"fmt"
	"reflect"
)

// Declaring a Struct
type Movie struct {
	Name   string
	Rating float32
}

type Actor struct {
	Name  string
	Age   int
	Movie Movie
}

type Alarm struct {
	Time  string
	Sound string
}

// Understanding Public and Private Values
type Values struct {
	privateValue string
	PublicValue  string
}

// Setting Default Values Using a Constructor Function
func NewAlarm(time string) Alarm {
	a := Alarm{
		Time:  time,
		Sound: "Klaxon",
	}
	return a
}
func main() {
	// Initializing a struct
	m := Movie{
		Name:   "Citizen Kane",
		Rating: 10,
	}

	// Intializing a struct using a variable
	var m1 Movie
	m1.Name = "Metropolis"
	m1.Rating = 0.99

	// Initializing a struct using new
	m2 := new(Movie)
	m2.Name = "Metropolis"
	m2.Rating = 0.99

	// create a struct using short variable assignment
	m3 := Movie{"Metropolis", 0.99}

	// Initializing a Nested Struct Using Shorthand Variable Assignment
	m4 := Actor{
		Name: "john reece",
		Age:  29,
		Movie: Movie{
			Name:   "Person of Interest",
			Rating: 9.67,
		},
	}

	// Comparing Structs
	if m1 == m3 {
		fmt.Println("They are the same")
	} else {
		fmt.Println("They are not same")
	}

	// Copying a Struct Using a Value Reference
	m5 := m3
	m5.Rating = 0.874
	fmt.Println(m5)
	fmt.Printf("%p\n", &m5)
	fmt.Printf("%p\n", &m3)

	// Copying a Struct Using a Pointer Reference
	m6 := &m1
	m6.Rating = 2.4664

	// Checking type of struct
	fmt.Printf("%T\n", m)
	fmt.Println(reflect.TypeOf(m))

	fmt.Printf("%+v\n", m2)
	fmt.Printf("%+v\n", m)
	fmt.Printf("%+v\n", m3)
	fmt.Printf("%+v\n", m4)
	fmt.Println(m4.Movie.Name)
	fmt.Printf("%+v\n", NewAlarm("07:00"))
	fmt.Printf("%+v\n", m1)
	fmt.Printf("%+v\n", *m6)
	fmt.Printf("%p\n", m6)
	fmt.Printf("%p\n", &m1)
}
