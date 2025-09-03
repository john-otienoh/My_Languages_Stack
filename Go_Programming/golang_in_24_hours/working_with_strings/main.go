package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Creating a String Literal
	s := "I am an interpreted string literal"
	fmt.Println(s)

	// Understanding Rune litearals
	r := '😊'

	fmt.Printf("Rune: %c, Unicode: %U\n", r, r)

	// Understanding Raw String Literals
	s2 := `After a backslash, certain single character escapes
	represent
	special values\nn is a line feed or new line \n\t t is a tab`

	fmt.Println(s2)

	// Concatenating Strings using + and +=
	s3 := "Oh sweet ignition" + " be my fuse"
	s4 := "Oh sweet ignition"
	s4 += "\n be my fuse"
	fmt.Println(s4)
	fmt.Println(s3)

	// Converting type to a string
	var i int = 1
	var s5 string = " egg"
	intToString := strconv.Itoa(i)
	var breakfast string = intToString + s5
	fmt.Println(breakfast)

	// Concatenating Strings with a Buffer

	var buffer bytes.Buffer

	for i := 0; i < 20; i++ {
		buffer.WriteString("z")
	}
	fmt.Println(buffer.String())

	// String Length
	fmt.Println(len(breakfast))

	// Understanding What a String Is
	s6 := "hello"
	fmt.Printf("%d\n", s6[0])
	fmt.Printf("%q\n", s6[0])
	fmt.Printf("%b\n", s6[0])

	// Converting String to lowercase
	fmt.Println(strings.ToLower("VERY IMPORTANT MESSAGE"))

	// Searching for a Substring in a String
	fmt.Println(strings.Index("surface", "face"))
	fmt.Println(strings.Index("moon", "aer"))

	// Trimming Space from a String
	fmt.Println(strings.TrimSpace("		I don't need all this space		"))

}
