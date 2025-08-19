package main

import (
	"fmt"
	"reflect"
)

func types() {
	var is_correct bool
	var name string = "John"
	var age int = 22
	var salary float64 = 120000.87

	fmt.Printf("My name is %s, %d years old earning %f monthly.\nVariable is_correct is of type %T \n", name, age, salary, is_correct)
	fmt.Println(reflect.TypeOf(is_correct))
}
