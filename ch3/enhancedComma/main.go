package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// enhanced comma inserts commas into a string of numbers
func comma(s string) string {

	//handling numbers with + or -
	prefix := ""

	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		prefix = string(s[0])
		s = s[1:]
	}

	//handling decimal numbers
	intPart := s
	fracPart := ""

	dotidx := strings.Index(s, ".")
	if dotidx != -1 {
		intPart = s[:dotidx]
		fracPart = s[dotidx:]
	}

	var buf bytes.Buffer
	for i := 0; i < len(intPart); i++ {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(intPart[i]) // WriteByte writes the character itself
	}
	return prefix + buf.String() + fracPart
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
	}
}
