package main

import (
	"fmt"
	// "github.com/GoesToEleven/GolangTraining/tree/master/02_package/stringutil"
	"math"
	"strings"
	// "temperature"

	"main.go/temperature"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(math.Pi)
	fmt.Println(strings.ToLower("STOP SHOUTING!"))

	// Using a third party Package
	// s := "ti esrever dna ti pilf nwod gniht ym tup I"
	// fmt.Println(stringutil.Reverse(s))

	// Creating a Package
	fmt.Println(temperature.CtoF(34))
}
