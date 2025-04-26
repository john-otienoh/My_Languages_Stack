package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numbers []int
	numbers = make([]int, 5)
	// matrix := make([]int, 3*3)

	for i := 0; i < len(numbers); i++ {
		numbers[i] = rand.Intn(100)
	}
	// for i := 0; i < 3; i++ {
	// 	matrix[i] = make([]int, 3)
	// 	for j := 0; j < 3; j++ {
	// 		matrix[i][j] = rand.Intn(100)
	// 	}
	// }

	fmt.Println(">>>>>display slice data")
	for j := 0; j < 5; j++ {
		fmt.Println(numbers[j])
	}
	// fmt.Println(">>>>>display slice 2D data")
	// for i := 0; i < 3; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Println(matrix[i][j])
	// 	}
	// }
}
