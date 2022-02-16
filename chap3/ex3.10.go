// non-recursive comma
package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args[1:]
	for _, num := range list {
		comma(num)
	}
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) {
	n := len(s)
	j := 3
	if n > 3 {
		for i := 0; i < n-3; {
			s = s[:j] + "," + s[j:]
			j += 4
			i += 3
		}
	}
	fmt.Println(s)
}
