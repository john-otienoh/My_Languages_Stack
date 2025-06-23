package main

import (
	"fmt"
	"ioutil"
)

func main() {
	writeFile("Heloo World")
}

func writeFile(message string) {
	bytes := []byte(message)
	ioutil.WriteFile("./file_operations/test.txt", bytes, 0644)
	fmt.Println("File Created")
}
