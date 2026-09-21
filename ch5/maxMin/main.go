// Exercise 5.15: variadic max and min functions, analogous to sum.
//
// A variadic function takes any number of trailing arguments: the "...int"
// means "zero or more ints", which arrive inside the function as a slice.
//
// The exercise asks what max/min should do with no arguments. There is no
// sensible "largest of nothing", so instead of guessing a value, these take
// one required argument (first) plus the rest, which makes a zero-argument
// call impossible: the compiler rejects it.
package main

import "fmt"

// max returns the largest of its arguments. It needs at least one argument.
func max(first int, rest ...int) int {
	biggest := first
	for _, v := range rest {
		if v > biggest {
			biggest = v
		}
	}
	return biggest
}

// min returns the smallest of its arguments. It needs at least one argument.
func min(first int, rest ...int) int {
	smallest := first
	for _, v := range rest {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
}

func main() {
	fmt.Println(max(3, 1, 4, 1, 5, 9, 2, 6)) // 9
	fmt.Println(min(3, 1, 4, 1, 5, 9, 2, 6)) // 1
	fmt.Println(max(42))                     // 42 (one argument is fine)

	// You can also spread a slice into a variadic call with "...":
	nums := []int{7, 3, 11, 5}
	fmt.Println(max(nums[0], nums[1:]...)) // 11
}
