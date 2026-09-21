// Exercise 5.16: a variadic version of strings.Join.
//
// The real strings.Join takes a slice and a separator: Join(elems, sep). This
// variadic version takes the separator first, then any number of strings, so
// you can call join(", ", "a", "b", "c") without building a slice yourself.
package main

import "fmt"

// join concatenates elems, placing sep between each pair. It is the same idea
// as the echo program from chapter 1: add the separator before every element
// except the first.
func join(sep string, elems ...string) string {
	result := ""
	for i, s := range elems {
		if i > 0 {
			result += sep
		}
		result += s
	}
	return result
}

func main() {
	fmt.Println(join(", ", "apple", "banana", "cherry")) // apple, banana, cherry
	fmt.Println(join("-", "2026", "09", "21"))           // 2026-09-21
	fmt.Println(join(", "))                              // "" (no elements)
	fmt.Println(join(", ", "solo"))                      // solo (no separator added)
}
