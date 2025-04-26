package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Printf("The value of i now is %d\n", i)
	}
	for j := 5; j > 0; j-- {
		if j == 3 {
			continue
		}
		fmt.Printf("The value of j now is %d\n", j)
	}

	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}
}
