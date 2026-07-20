// Package tempconv performs Celsius, Fahrenheit, and Kelvin conversions.
//
// Exercise 2.1: add types, constants, and functions for the Kelvin scale,
// where zero Kelvin is -273.15°C and a difference of 1K has the same
// magnitude as 1°C.
package tempconv

import "fmt"

// Each scale is its own named type, even though all three share float64 as
// their underlying type. Because Go does not convert between named types
// implicitly, the compiler rejects mixing them (a Celsius cannot be passed
// where a Kelvin is expected), which is what makes the conversions below
// worth having.
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15 // coldest possible temperature
	FreezingC     Celsius = 0       // freezing point of water
	BoilingC      Celsius = 100     // boiling point of water

	AbsoluteZeroK Kelvin = 0      // same point as AbsoluteZeroC
	FreezingK     Kelvin = 273.15 // same point as FreezingC
	BoilingK      Kelvin = 373.15 // same point as BoilingC
)

// The String methods satisfy fmt.Stringer, so fmt.Println and the %s and %v
// verbs print these values with their unit attached, e.g. "373.15K".
// Kelvin carries no degree sign by convention.
//
// Note the %g verb: using %v or %s here would call String again and recurse
// forever. Formatting float64(k) instead would also be safe.
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gK", k) }
