/*
Package main provides a word frequency counter.
It reads words from a file and prints the frequency of each word.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq(filename string) map[string]int {
	count := make(map[string]int)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		count[input.Text()]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
	}
	return count
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: wordfreq <file>")
		os.Exit(1)
	}
	counts := wordfreq(os.Args[1])
	fmt.Printf("word\tcount\n")
	for word, n := range counts {
		fmt.Printf("%q\t%d\n", word, n)
	}
}
