package main

import "fmt"

func main() {
	ex_one()
	ex_two()
	ex_three()
}

func ex_one() {
	/*
		Function that defines a variable named greetings of type slice of strings with the following values:
		"Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
		Create a subslice containing the first two values;
		a second subslice with the second, third, and fourth values; and a third subslice with the fourth and fifth values.
		Print out all four slices.
	*/
	var greetings = []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	firstTwoValues := greetings[:2]
	secondThirdAndForthValues := greetings[1:4]
	forthAndFifthValues := greetings[3:]
	fmt.Println(firstTwoValues, secondThirdAndForthValues, forthAndFifthValues)
}

func ex_two() {
	/*
		Function that defines a string variable called message with the value "Hi and " and
		prints the fourth rune in it as a character, not a number.
	*/
	message := "Hi 👩 and 👨"
	fourthRune := []rune(message)[3]
	fmt.Println(string(fourthRune))

}

func ex_three() {
	/*
		Function that defines a struct called Employee with three fields: firstName, lastName, and id.
		The first two fields are of type string, and the last field (id) is of type int.
		Create three instances of this struct using whatever values you’d like.
		Initialize the first one using the struct literal style without names,
		the second using the struct literal style with names,
		and the third with a var declaration.
		Use dot notation to populate the fields in the third struct.
		Print out all three structs.
	*/
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	firstEmployee := Employee{"Joe", "Doe", 10334}
	secondEmployee := Employee{
		"John",
		"Doe",
		20334,
	}
	var thirdEmployee Employee
	thirdEmployee.firstName = "Jane"
	thirdEmployee.lastName = "Doe"
	thirdEmployee.id = 25334

	fmt.Println(firstEmployee, secondEmployee, thirdEmployee)
}
