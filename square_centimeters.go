package units

import "fmt"

var _ Area = SquareCentimeters(0.0)

// SquareCentimeters stores an area as a number of square centimeters.
// Can be used in calculations or converted to other units.
type SquareCentimeters float64

// SquareMillimeters converts the square centimeters into square millimeters
// and returns the result.
func (s SquareCentimeters) SquareMillimeters() SquareMillimeters {
	return SquareMillimeters(float64(s) * 100.0)
}

// SquareCentimeters just returns a copy of itself.
// Required to implement the Area interface.
func (s SquareCentimeters) SquareCentimeters() SquareCentimeters {
	return s
}

// SquareInches converts the square centimeters into square inches
// and returns the result.
func (s SquareCentimeters) SquareInches() SquareInches {
	return SquareInches(float64(s) / 6.4516)
}

// String converts the square centimeters into a string with 2
// decimals and appends the units.
//
// I.e. 18.92 cm^2
func (s SquareCentimeters) String() string {
	return fmt.Sprintf("%.2f cm^2", float64(s))
}
