// Exercise 4.8: modify charcount to count letters, digits, and so on in their
// Unicode categories, using functions like unicode.IsLetter.
//
// It reads UTF-8 text from standard input one rune at a time and tallies each
// valid rune into a Unicode category (letter, digit, space, ...), then prints
// the totals.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	categories := make(map[string]int) // counts per Unicode category
	invalid := 0                       // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // rune, byte count, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charCategories: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		// Classify the rune. The cases are checked top to bottom, so the
		// first match wins; order them from most specific to the catch-all.
		switch {
		case unicode.IsLetter(r):
			categories["letter"]++
		case unicode.IsDigit(r):
			categories["digit"]++
		case unicode.IsSpace(r):
			categories["space"]++
		case unicode.IsPunct(r):
			categories["punct"]++
		case unicode.IsSymbol(r):
			categories["symbol"]++
		case unicode.IsControl(r):
			categories["control"]++
		default:
			categories["other"]++
		}
	}

	fmt.Printf("category\tcount\n")
	for cat, n := range categories {
		fmt.Printf("%s\t%d\n", cat, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
