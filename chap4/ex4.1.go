package main

import (
	"crypto/sha256"
)

// // pc[i] is the population count of i.
// var pc [256]byte

// func init() {
// 	for i := range pc {
// 		pc[i] = pc[i/2] + byte(i&1)
// 	}
// }

// // PopCount returns the population count (number of set bits) of x.
// func PopCount(x uint64) int {
// 	return int(pc[byte(x>>(0*8))] +
// 		pc[byte(x>>(1*8))] +
// 		pc[byte(x>>(2*8))] +
// 		pc[byte(x>>(3*8))] +
// 		pc[byte(x>>(4*8))] +
// 		pc[byte(x>>(5*8))] +
// 		pc[byte(x>>(6*8))] +
// 		pc[byte(x>>(7*8))])
// }

func CountDifferentBits(hash1 [32]byte, hash2 [32]byte) {
	totalDiff := 0
	for i := range hash1 {
		for j := 0; j < 8; j++ {
			shift1 := hash1[i] >> j & 1
			shift2 := hash2[i] >> j & 1
			if shift1 != shift2 {
				totalDiff++
			}
		}
	}

}

func main() {
	CountDifferentBits(sha256.Sum256([]byte("a")), sha256.Sum256([]byte("b")))
}
