package main

import "fmt"

// rotate rotates s left by n places, in place. n may exceed len(s), and a
// negative n rotates to the right.
func rotate(s []int, n int) {
	if len(s) == 0 {
		return
	}
	n %= len(s) // handle rotations larger than the slice
	if n < 0 {
		n += len(s) // negative n means rotate right
	}
	tmp := make([]int, 0, len(s))
	tmp = append(tmp, s[n:]...) // the part that moves to the front
	tmp = append(tmp, s[:n]...) // followed by the wrapped-around front
	copy(s, tmp)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 2)
	fmt.Println(s) // [2 3 4 5 0 1]
}
