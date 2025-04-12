package main

import "fmt"

func WordSlicer() {
	/*
		Function to print
		word_first_3 should contain the first 3 letters of the variable word
		word_last_2 should contain the last 2 letters of the variable word
		middle_word should contain the value of the variable word without the first and last letters
	*/
	word := "Holberton"
	word_first_3, word_last_2, middle_world := word[:3], word[7:], word[1:8]
	fmt.Printf("First 3 letters: %s.\n", word_first_3)
	fmt.Printf("Last 2 letters: %s.\n", word_last_2)
	fmt.Printf("Middle word: %s.\n", middle_world)

}
