package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		firstname string = "John"
		lastname  string = "Doe"
	)
	// Concatenating strings
	fullname := firstname + " " + lastname
	fullname2 := fmt.Sprintf("%s %s", firstname, lastname)
	fmt.Println(fullname, "\n", fullname2)

	// String To Numeric
	var (
		str_value1  = "2"
		str_value2  = "4.674"
		err         error
		int_value   int64
		float_value float64
	)
	int_value, err = strconv.ParseInt(str_value1, 10, 32)
	if err == nil {
		fmt.Println(int_value)
	} else {
		fmt.Println(err)
	}
	float_value, err = strconv.ParseFloat(str_value2, 64)
	if err == nil {
		fmt.Println(float_value)
	} else {
		fmt.Println(err)
	}

	// Numeric to String
	var (
		num1     int     = 2
		float1   float64 = 5.7823
		str_num1 string
		str_num2 string
	)
	str_num1 = fmt.Sprintf("%d", num1)
	str_num2 = fmt.Sprintf("%f", float1)
	fmt.Println(str_num1, "\n", str_num2)

	// String Parser
	data := "Berlin;Amsterdam;London;Tokyo"
	cities := strings.Split(data, ";")
	for _, city := range cities {
		fmt.Println(city)
	}

	// String Length
	fmt.Printf("Length = %d\n", len(data))

	// Copy Data
	fmt.Println(data[0:len(data)])
	fmt.Println(data[:6])
	fmt.Println(data[7:16])
	fmt.Println(data[len(data)-5 : len(data)])

	// Lowercase and Uppercase
	fmt.Println(strings.ToLower(data))
	fmt.Println(strings.ToUpper(data))
}
