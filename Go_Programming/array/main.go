package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers [5]int
	var names [5]string
	var matrix [3][3]int
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i
		names[i] = fmt.Sprintf("Name %d", i)
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = rand.Intn(100)
		}
	}
	fmt.Println(">>>>>display array data")
	for j := 0; j < 5; j++ {
		fmt.Println(numbers[j])
		fmt.Println(names[j])
	}
	fmt.Println(">>>>>display matrix data")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
