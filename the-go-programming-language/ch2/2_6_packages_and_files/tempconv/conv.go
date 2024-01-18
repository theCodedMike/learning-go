// Package tempconv
//
// T(°F) = t(°C) * 1.8 + 32
//
// T(K) = t(°C) + 273.15
package tempconv

import "gopl.io/ch2/2_6_packages_and_files/exercise_2_1"

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) exercise_2_1.Kelvin {
	return exercise_2_1.Kelvin(float64(c) - float64(AbsoluteZeroC))
}

// KToC converts a Kelvin temperature to Celsius.
func KToC(k exercise_2_1.Kelvin) Celsius {
	return Celsius(float64(k) + float64(AbsoluteZeroC))
}
