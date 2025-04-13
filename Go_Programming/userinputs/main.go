package main

import "fmt"

func main() {
	fmt.Println("Hello There")
	fmt.Print("Please enter your name: ")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello %s welcome to our site\n", name)
}
