package tempconv

// The conversions live in their own file from the type and constant
// declarations in tempconv.go. Both files declare "package tempconv", and in
// Go the package (not the file) is the unit of scope, so these functions use
// Celsius, Fahrenheit, and Kelvin directly with no import needed.
//
// Each body converts explicitly, e.g. Kelvin(c + 273.15). The arithmetic
// produces a Celsius, and the named-type conversion relabels it as a Kelvin.
// The untyped constant 273.15 adopts whichever type the expression needs.

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return Kelvin((f-32)*5/9 + 273.15) }

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return Fahrenheit((k-273.15)*9/5 + 32) }

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }
