package main

import (
	"fmt"
)

func main() {
	s := []byte("Hello... It's me")
	fmt.Printf("Before: %s\n", s)
	reverse(&s)
	fmt.Printf("After : %s\n", s)
}

func reverse(s *[]byte) {
	for i := range (*s)[:len(*s)/2] {
		(*s)[i], (*s)[len(*s)-1-i] = (*s)[len(*s)-1-i], (*s)[i]
	}

}
