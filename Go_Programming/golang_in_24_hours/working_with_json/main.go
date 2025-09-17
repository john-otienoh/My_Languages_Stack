package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Creating a Struct
type Person struct {
	Name    string
	Age     int
	Hobbies []string
}

// Using struct tags
type Person2 struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

// Omitting Empty Struct Fields
type Person1 struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
}
type User struct {
	Name string `json:"name"`
	Blog string `json:"blog"`
}

func main() {
	hobbies := []string{"Cycling", "Cheese", "Techno"}
	p := Person{
		Name:    "John",
		Age:     22,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p)
	// Encoding a Struct as JSON
	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)

	// Using Struct tags
	p2 := Person2{
		Name:    "John",
		Age:     22,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p2)
	// Encoding a Struct as JSON
	jsonByteData1, err := json.Marshal(p2)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData1 := string(jsonByteData1)
	fmt.Println(jsonStringData1)

	// Omitting Empty Struct Fields
	p1 := Person1{
		Name:    "John",
		Age:     22,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p1)
	// Encoding a Struct as JSON
	jsonByteData2, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData2 := string(jsonByteData2)
	fmt.Println(jsonStringData2)

	// Decoding a struct into JSON
	jsonStringData3 := `{"name":"George","age":40,"hobbies":["Cycling", "Cheese","Techno"]}`
	jsonByteData3 := []byte(jsonStringData3)
	p3 := Person2{}
	err1 := json.Unmarshal(jsonByteData3, &p3)
	if err1 != nil {
		log.Fatal(err1)
	}
	fmt.Printf("%+v\n", p3)

	// JSON Data Types in go
	// Boolean	bool
	// Number	float64
	// String	string
	// Array	[]interface{}
	// Object	map[string]interface{}
	// Null	nil

	// Fetching JSON over HTTP
	var u User
	res, err2 := http.Get("https://api.github.com/users/shapeshed")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer res.Body.Close()
	err2 = json.NewDecoder(res.Body).Decode(&u)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Printf("%+v\n", u)

}
