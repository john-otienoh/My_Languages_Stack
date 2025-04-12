package main

func strlen(s *string) int {
	count := 0
	for _ = range *s {
		count++
	}
	return count
}
