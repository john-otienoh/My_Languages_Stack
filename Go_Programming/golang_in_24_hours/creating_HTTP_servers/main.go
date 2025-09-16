package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Basic HTTP Server in Go
func greet(w http.ResponseWriter, r *http.Request) {
	// Adding a 404 Response
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Adding a Header to a Response
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Responding with Different Content Types
	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprintf(w, `{"message": "Hello World"}! %s`, time.Now())
	case "application/xml":
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		fmt.Fprintf(w, "Hello World! %s\n", time.Now())
	default:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "Hello World! %s\n", time.Now())
	}
	// fmt.Fprintf(w, "Hello World! %s\n", time.Now())

	// Responding to Different Types of Requests
	switch r.Method {
	case "GET":
		for k, v := range r.URL.Query() {
			fmt.Printf("%s: %s\n", k, v)
		}
		w.Write([]byte("This is a GET request\n"))
	case "POST":
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", reqBody)
		w.Write([]byte("This is a POST request\n"))
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))
	}
}

func main() {
	// Routing
	http.HandleFunc("/", greet)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
