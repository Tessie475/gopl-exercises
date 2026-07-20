package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl-exercises/ch02/ex2.2/genconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := genconv.Fahrenheit(t)
		c := genconv.Celsius(t)
		m := genconv.Meters(t)
		ft := genconv.Feet(t)
		p := genconv.Pound(t)
		k := genconv.Kilogram(t)
		g := genconv.Grams(t)
		fmt.Printf("%s = %s, %s = %s\n", f, genconv.FToC(f), c, genconv.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", m, genconv.MToF(m), ft, genconv.FToM(ft))
		fmt.Printf("%s = %s, %s = %s\n", p, genconv.PToK(p), k, genconv.KToP(k))
		fmt.Printf("%s = %s, %s = %s\n", k, genconv.KToG(k), g, genconv.GToK(g))
	}
}
