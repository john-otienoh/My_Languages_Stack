package main

import "fmt"

func main() {
	// array to include (a) cheese(s) that you like
	var cheeses [2]string
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"
	fmt.Println(cheeses)
	// Create a slice with four elements.
	var s = make([]string, 4)
	s[0] = "zero"
	s[1] = "one"
	s[2] = "two"
	s[3] = "three"
	fmt.Println(s)
	// Create a new slice and copy the third and fourth elements only into it.
	var s2 = make([]string, 2)
	copy(s2, s[2:])
	fmt.Println(s2)
	// Create a map from the following list of HTML elements.
	var htmlElements = make(map[string]string)
	htmlElements["p"] = "Paragraph"
	htmlElements["img"] = "Image"
	htmlElements["h1"] = "Heading One"
	htmlElements["h2"] = "Heading Two"
	fmt.Println(htmlElements)
}
