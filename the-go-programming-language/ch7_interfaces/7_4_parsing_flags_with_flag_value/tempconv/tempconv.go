package tempconv

import (
	"flag"
	"fmt"
	"gopl.io/ch2_program_structure/2_6_packages_and_files/exercise_2_1"
	"gopl.io/ch2_program_structure/2_6_packages_and_files/tempconv"
)

type Celsius = tempconv.Celsius

type celsiusFlag struct {
	Celsius
}

// Set : Impl `type Value interface {}` //
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	_, _ = fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "c", "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "f", "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	case "k", "K":
		f.Celsius = tempconv.KToC(exercise_2_1.Kelvin(value))
		return nil
	}

	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
