package main

import (
	"fmt"
	"math/rand"
)

func main() {
	products := make(map[string]int)
	customers := make(map[string]int)
	for i := 0; i < 3; i++ {
		products[fmt.Sprintf("Product %d", i)] = rand.Intn(100)
		customers[fmt.Sprintf("Customer %d", i)] = rand.Intn(100)
	}
	fmt.Println(products)
	fmt.Println(customers)
}
