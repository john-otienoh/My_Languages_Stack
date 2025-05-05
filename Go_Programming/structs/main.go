package main

import (
	"fmt"
	"time"
)

type Employee struct {
	name    string
	age     int
	salary  float64
	created time.Time
}

func main() {
	emp := Employee{
		name:    "John Doe",
		age:     20,
		salary:  45437.98,
		created: time.Now(),
	}
	newEmployee := new(Employee)
	newEmployee.name = "Jane Doe"
	newEmployee.age = 20
	newEmployee.salary = 74848.84
	newEmployee.created = time.Now()

	fmt.Println(newEmployee)
	fmt.Println(emp)
}
