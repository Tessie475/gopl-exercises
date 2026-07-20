// Command ex2.2 is a general-purpose unit converter. Each number given on the
// command line is treated as a value in every supported unit at once, and the
// equivalent in the paired unit is printed: temperature, length, and weight.
//
//	$ go run . 100
//
// Still to do for the exercise: when no arguments are supplied it should read
// numbers from standard input instead of doing nothing. See the review notes.
package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl-exercises/ch02/ex2.2/genconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		// ParseFloat turns the text argument into a float64; the second
		// argument, 64, asks for float64 rather than float32 precision.
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		// The same number reinterpreted as each unit. These are type
		// conversions, not calculations: t is just relabelled.
		f := genconv.Fahrenheit(t)
		c := genconv.Celsius(t)
		m := genconv.Meters(t)
		ft := genconv.Feet(t)
		p := genconv.Pound(t)
		k := genconv.Kilogram(t)
		g := genconv.Grams(t)
		// Each %s calls the value's String method, which supplies the unit.
		fmt.Printf("%s = %s, %s = %s\n", f, genconv.FToC(f), c, genconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", m, genconv.MToF(m), ft, genconv.FToM(ft))
		fmt.Printf("%s = %s, %s = %s\n", p, genconv.PToK(p), k, genconv.KToP(k))
		fmt.Printf("%s = %s, %s = %s\n", k, genconv.KToG(k), g, genconv.GToK(g))
	}
}
