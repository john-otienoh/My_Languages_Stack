package main

import "fmt"

func reverse(s *string) {
	s_runes := []rune(*s)
	for i, j := 0, len(s_runes)-1; i < j; i, j = i+1, j-1 {
		s_runes[i], s_runes[j] = s_runes[j], s_runes[i]
	}
	fmt.Println(string(s_runes))
}
