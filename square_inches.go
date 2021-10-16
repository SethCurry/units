package units

import "fmt"

var _ Area = SquareInches(0.0)

// SquareInches stores an area as a float32 number of square inches.
// Satisfies the Area interface.
type SquareInches float64

// SquareMillimeters converts the square inches into square millimeters
// and returns the value.
func (s SquareInches) SquareMillimeters() SquareMillimeters {
	return SquareMillimeters(float64(s) * 645.16)
}

// SquareCentimeters converts the squre inches into square centimeters
// and returns the value.
func (s SquareInches) SquareCentimeters() SquareCentimeters {
	return SquareCentimeters(float64(s) * 6.4516)
}

// SquareInches just returns a copy of itself.
// Required to satisfy the Area interface.
func (s SquareInches) SquareInches() SquareInches {
	return s
}

// String converts the SquareInches into a string with
// 2 decimals and appends the units.
//
// I.e. "29.23 sq. in"
func (s SquareInches) String() string {
	return fmt.Sprintf("%.2f sq. in", float64(s))
}
