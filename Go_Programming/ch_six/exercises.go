package main

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstname, lastname string, age int) Person {
	return Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}
}

func MakePersonPointer(firstname, lastname string, age int) *Person {
	return &Person{
		FirstName: firstname,
		LastName:  lastname,
		Age:       age,
	}
}

func main() {
	p := MakePerson("John", "Doe", 20)
	fmt.Println(p)
	p2 := MakePersonPointer("John", "Doe", 20)
	fmt.Println(p2)
}
