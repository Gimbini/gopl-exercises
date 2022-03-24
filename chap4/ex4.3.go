package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4}
	fmt.Println(s)
	reverse(&s)
	fmt.Println(s)
}

// reverse reverses a slice of ints in place.
func reverse(s *[]int) {
	for i := range (*s)[:len(*s)/2] {
		(*s)[i], (*s)[len(*s)-1-i] = (*s)[len(*s)-1-i], (*s)[i]
	}

	// func reverse(s *[]int) {
	// 	var slice []int
	// 	slice = *s

	// 	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
	// 		(*s)[i], slice[j] = slice[j], slice[i]
	// 	}

	// for i := range s {
	// 	s[i], s[len(s)-i] = s[len(s)-i], s[i]
	// }
}
