package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(s)
	fmt.Println(rotate(s, 5))
}

func rotate(s []int, n int) []int {
	var z []int
	z = make([]int, len(s))

	copy(z[:len(s)-n], s[n:])
	copy(z[len(s)-n:], s[:n])
	return z
}
