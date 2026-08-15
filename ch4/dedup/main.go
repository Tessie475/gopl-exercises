package main

import "fmt"

// dedup removes adjacent duplicate strings in place and returns the
// compacted slice.
func dedup(s []string) []string {
	i := 0 // index where the next kept element goes
	for _, x := range s {
		if i == 0 || s[i-1] != x { // different from the last kept element?
			s[i] = x
			i++
		}
	}
	return s[:i]
}

func main() {
	s := []string{"a", "a", "b", "b", "b", "c", "a", "a"}
	fmt.Printf("%q\n", dedup(s)) // ["a" "b" "c" "a"]
}
