package units

import "fmt"

// CubicCentimeters stores a volume measurement as a float64
// of cubic centimeters.
type CubicCentimeters float64

// CubicMillimeters converts the cubic centimeters measurement
// to cubic millimeters and returns it.
func (c CubicCentimeters) CubicMillimeters() CubicMillimeters {
	return CubicMillimeters(float64(c) * 1000.0)
}

// CubicCentimeters simply returns a copy of itself.
// Required to implement the Volume interface.
func (c CubicCentimeters) CubicCentimeters() CubicCentimeters {
	return c
}

// CubicInches converts the cubic centimeters to cubic inches
// and returns them.
func (c CubicCentimeters) CubicInches() CubicInches {
	return CubicInches(float64(c) / 16.3871)
}

// String converts the cubic centimeters to a human-readable string
// with 2 decimal places.
//
// I.e. "56.78 cm^3"
func (c CubicCentimeters) String() string {
	return fmt.Sprintf("%.2f cm^3", float64(c))
}
