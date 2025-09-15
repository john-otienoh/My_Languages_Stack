package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func echo(s string) {
	fmt.Println(s)
	return
}
func main() {
	// Logging in Go
	log.Printf("This is a log message")

	// Writing Logs to a File
	f, err := os.OpenFile("server_logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	log.SetOutput(f)
	for i := 1; i <= 5; i++ {
		log.Printf("Log Iteration %d", i)
	}

	// Sample Logging Program
	for i := 1; i <= 5; i++ {
		log.Printf("Log iteration %d", i)
	}
	// Simple Program to Demonstrate Using Delve
	s := "Hello World"
	t := "Goodbye Cruel World"
	echo(s)
	echo(t)
	// Logging a Fatal Error
	err1 := errors.New("Server Crashing")
	log.Fatal(err1)
}
