package main

import (
	"fmt"
)

func main() {
	// Arrays
	var cheeses [2]string
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	fmt.Println(cheeses)

	// Slices
	var cheese = make([]string, 2)
	cheese[0] = "Mariolles"
	cheese[1] = "Époisses de Bourgogne"

	// Adding element to slice
	cheese = append(cheese, "Camebert", "Reblochon", "Picodon")
	fmt.Println(cheese)

	// Deleting element from a slice
	cheese = append(cheese[:2], cheese[2+1:]...)
	fmt.Println(cheese)

	// Copying Elements from a Slice
	var smellyCheeses = make([]string, 2)
	copy(smellyCheeses, cheese)
	fmt.Println(smellyCheeses)

	// Maps
	var players = make(map[string]int)
	players["cook"] = 32
	players["bairstow"] = 27
	players["stokes"] = 26
	fmt.Println(players)
	// Deleting Elements from a Map
	delete(players, "cook")
	fmt.Println(players)
}
