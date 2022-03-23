package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	firstWord := string(os.Args[1])
	secondWord := string(os.Args[2])
	popWord := secondWord

	for _, alphabet := range firstWord {
		popWord = strings.Replace(popWord, string(alphabet), "", 1)
	}

	isAnagram := popWord == ""

	if isAnagram {
		fmt.Printf("%s and %s are anagrams \n", firstWord, secondWord)
	} else {
		fmt.Printf("%s and %s are not anagrams \n", firstWord, secondWord)
	}

}
