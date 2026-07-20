// Package genconv converts between units of length, weight, and temperature.
//
// Exercise 2.2: a general-purpose unit-conversion program analogous to cf,
// reading numbers from its command-line arguments or from standard input
// when no arguments are given.
//
// The layout mirrors the book's tempconv package: types, constants, and
// String methods here; the conversion functions in conv.go.
package genconv

import (
	"fmt"
	"math"
)

// One named type per unit. They all share float64 underneath, but Go treats
// them as distinct types, so the compiler will not let a Meters be used where
// a Feet is wanted. That is the whole point of declaring them separately.
type Feet float64
type Meters float64
type Pound float64
type Kilogram float64
type Grams float64
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	// Conversion factors. Lowercase means unexported: usable anywhere inside
	// this package, invisible to code that imports it.
	metersToFeet      Meters   = 3.281
	feetToMeters      Feet     = 0.3048
	kilogramsToPounds Kilogram = 2.20462
	poundsToKilograms Pound    = 0.4536
	kilogramsToGrams  Kilogram = 1000
	gramsToKilograms  Grams    = 0.001

	// Named temperature reference points, exported for callers to use.
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	AbsoluteZeroK Kelvin  = 0
	FreezingK     Kelvin  = 273.15
	BoilingK      Kelvin  = 373.15
)

// The String methods satisfy fmt.Stringer, so %s and %v print each value with
// its unit attached.
//
// math.Round is applied first, so these print whole numbers only. See the
// caveat noted in the review: fractional results such as 0.3048 m display
// as "0 m".
func (m Meters) String() string {
	return fmt.Sprintf("%g m", math.Round(float64(m)))
}

func (f Feet) String() string {
	return fmt.Sprintf("%g Ft", math.Round(float64(f)))
}

func (p Pound) String() string {
	return fmt.Sprintf("%g Pound", math.Round(float64(p)))
}

func (k Kilogram) String() string {
	return fmt.Sprintf("%g Kg", math.Round(float64(k)))
}

func (g Grams) String() string {
	return fmt.Sprintf("%g g", math.Round(float64(g)))
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", math.Round(float64(c)))
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", math.Round(float64(f)))
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g K", math.Round(float64(k)))
}
