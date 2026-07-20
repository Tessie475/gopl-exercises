package genconv

// Conversion functions for every unit pair the package supports. They live in
// a separate file from the declarations in genconv.go; both say "package
// genconv", and scope in Go is per-package rather than per-file, so the types
// and factor constants are visible here without any import.

// Length.

// MToF converts a length in meters to feet.
func MToF(m Meters) Feet {
	return Feet(m * metersToFeet)
}

// FToM converts a length in feet to meters.
func FToM(f Feet) Meters {
	return Meters(f * feetToMeters)
}

// Weight. Note that K here means Kilogram, whereas in the temperature
// functions below K means Kelvin. See the naming caveat in the review.

// PToK converts a weight in pounds to kilograms.
func PToK(p Pound) Kilogram {
	return Kilogram(p * poundsToKilograms)
}

// KToP converts a weight in kilograms to pounds.
func KToP(k Kilogram) Pound {
	return Pound(k * kilogramsToPounds)
}

// KToG converts a weight in kilograms to grams.
func KToG(k Kilogram) Grams {
	return Grams(k * kilogramsToGrams)
}

// GToK converts a weight in grams to kilograms.
func GToK(g Grams) Kilogram {
	return Kilogram(g * gramsToKilograms)
}

// Temperature.

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

// FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin {
	return Kelvin((f-32)*5/9 + 273.15)
}

// KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit {
	return Fahrenheit((k-273.15)*9/5 + 32)
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
