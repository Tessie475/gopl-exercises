package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// squash collapses runs of Unicode spaces into a single ' ' in place and
// returns the compacted slice.
func squash(b []byte) []byte {
	out := 0
	inSpace := false
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if unicode.IsSpace(r) {
			if !inSpace {
				b[out] = ' '
				out++
				inSpace = true
			}
		} else {
			// copy is memmove-safe, so this overlapping in-place move is fine.
			copy(b[out:out+size], b[i:i+size])
			out += size
			inSpace = false
		}
		i += size
	}
	return b[:out]
}

func main() {
	b := []byte("  the\t\tquick   brown   fox  ")
	fmt.Printf("%q\n", string(squash(b))) // " the quick brown fox "
}
