package main

import (
	"fmt"
	"strconv"
)

func main() {
	var b bool
	fmt.Printf("Typeof variable b is: %T\n", b)
	var s string = strconv.FormatBool(true)
	fmt.Printf("Typeof variable b is: %T\n", s)

	var s1 string = "truee"
	b1, err := strconv.ParseBool(s1)
	fmt.Println(b1, err)

	s2 := strconv.FormatBool(b1)
	fmt.Println(s2)

	var age_int int16 = 22
	age_float := float64(age_int)

	fmt.Printf("Typeof variable age1 is: %T\n", age_int)
	fmt.Printf("Typeof variable age2 is: %T\n", age_float)

}
