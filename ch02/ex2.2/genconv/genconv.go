package genconv

import (
	"fmt"
	"math"
)

type Feet float64
type Meters float64
type Pound float64
type Kilogram float64
type Grams float64
type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	metersToFeet      Meters   = 3.281
	feetToMeters      Feet     = 0.3048
	kilogramsToPounds Kilogram = 2.20462
	poundsToKilograms Pound    = 0.4536
	kilogramsToGrams  Kilogram = 1000
	gramsToKilograms  Grams    = 0.001
	AbsoluteZeroC     Celsius  = -273.15
	FreezingC         Celsius  = 0
	BoilingC          Celsius  = 100
	AbsoluteZeroK     Kelvin   = 0
	FreezingK         Kelvin   = 273.15
	BoilingK          Kelvin   = 373.15
)

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
