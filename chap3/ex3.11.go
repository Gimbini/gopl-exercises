// handles floating points, but not 'optional sign'
package main

import (
	"fmt"
	"os"
	"strings"
)

// comma inserts commas in a non-negative decimal integer string.

func main() {
	args := os.Args[1:]
	for _, num := range args {
		if !strings.Contains(num, ".") {
			fmt.Println(comma(num))
			continue
		} else {
			dotIndex := strings.Index(num, ".")
			quotient := num[:dotIndex]
			decimals := num[dotIndex:]
			fmt.Println(comma(quotient) + decimals)
		}
	}
}

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
