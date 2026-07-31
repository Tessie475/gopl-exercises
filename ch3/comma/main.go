package main

import (
	"bytes"
	"fmt"
	"os"
)

// comma inserts commas into a non-negative string of digits.
func comma(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		if i > 0 && (len(s)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i]) // WriteByte writes the character itself
	}
	return buf.String()
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}
