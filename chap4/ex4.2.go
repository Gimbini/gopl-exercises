package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {
	algorithm := "sha256"
	if len(os.Args) > 1 {
		algorithm = os.Args[1]
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if algorithm == "sha256" {
			fmt.Printf("%x\n", sha256.Sum256([]byte(scanner.Text())))
		}
		if algorithm == "sha384" {
			fmt.Printf("%x\n", sha512.Sum384([]byte(scanner.Text())))
		}
		if algorithm == "sha512" {
			fmt.Printf("%x\n", sha512.Sum512([]byte(scanner.Text())))
		}
	}
}
