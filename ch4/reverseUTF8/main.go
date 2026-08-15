package main

import (
	"fmt"
	"unicode/utf8"
)

// reverseBytes reverses a byte slice in place.
func reverseBytes(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

// reverseUTF8 reverses the runes of a UTF-8-encoded byte slice in place.
func reverseUTF8(b []byte) {
	for i := 0; i < len(b); { // step 1: reverse each rune's bytes
		_, size := utf8.DecodeRune(b[i:])
		reverseBytes(b[i : i+size])
		i += size
	}
	reverseBytes(b) // step 2: reverse the whole slice
}

func main() {
	b := []byte("Hello, 世界")
	reverseUTF8(b)
	fmt.Printf("%q\n", string(b)) // "界世 ,olleH"
}
