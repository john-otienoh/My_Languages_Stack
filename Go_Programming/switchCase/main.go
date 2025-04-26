package main

import "fmt"

func main() {
	choice := 2
	switch choice {
	case 0:
		fmt.Println("Choice = 0")
	case 1:
		fmt.Println("Choice = 1")
	case 2:
		fmt.Println("Choice = 2")
	case 3:
		fmt.Println("Choice = 3")
	default:
		fmt.Println("Choice = Other")
	}

}
