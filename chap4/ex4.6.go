// Removes Duplicate adjacent entries in an array

package main

import (
	"fmt"
)

func main() {
	s := []byte{65, 70, 90, 32, 32, 32, 90, 32, 32, 80}
	fmt.Printf("Before: %s \n", s)
	fmt.Printf("After : %s \n", remDupSpaceInplace(s))
}

func remDupSpaceInplace(s []byte) []byte {
	sliceLength := len(s)
	i := 0
	for i < sliceLength {
		for true {
			if i >= sliceLength-1 {
				break
			}
			if s[i] == byte(32) && s[i] == s[i+1] {
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

func remove(slice []byte, i int) []byte {
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
