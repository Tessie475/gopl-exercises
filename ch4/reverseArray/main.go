package main

import "fmt"

// reverse reverses the array in place. len(ptr) works on an array pointer,
// and ptr[i] auto-dereferences, so no (*ptr)[i] is needed.
func reverse(ptr *[6]int) {
	for i, j := 0, len(ptr)-1; i < j; i, j = i+1, j-1 {
		ptr[i], ptr[j] = ptr[j], ptr[i]
	}
}

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	reverse(&a) // pass the address, not the array itself
	fmt.Println(a)
}
