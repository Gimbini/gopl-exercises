// Removes Duplicate adjacent entries in an array

package main

import (
	"fmt"
)

func main() {
	s := []int{1, 1, 1, 2, 2, 3, 2, 4, 5, 5, 5, 5, 5}
	fmt.Printf("Before: %v \n", s)
	fmt.Printf("After : %v \n", remDupInplace(s))
}

func remDupInplace(s []int) []int {
	sliceLength := len(s)
	i := 0
	for i < sliceLength {
		for true {
			if i >= sliceLength-1 {
				break
			}
			if s[i] == s[i+1] {
				s = remove(s, i+1)
				sliceLength--
			} else {
				break
			}

		}
		i++
	}
	return s
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// func remDupInplace(s []int) []int {
// 	sliceLength := len(s)
// 	i := 0
// 	for i < sliceLength {
// 		for true {
// 			if i >= sliceLength {
// 				break
// 			}
// 			if s[i] == s[i+1] {
// 				s = remove(s, i)
// 				sliceLength--
// 			}
// 		}
// 		i++
// 	}
// 	return s
// }
