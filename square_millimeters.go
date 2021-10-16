package units

import "fmt"

// SquareMillimeters stores an area in square millimeters.
//
// It satisfies the Area interface as well, allowing it to be used
// as part of operations.
type SquareMillimeters float64

// SquareMillimeters just returns itself.  It's used to satisfy
// the Area interface.
func (s SquareMillimeters) SquareMillimeters() SquareMillimeters {
	return s
}

// SquareCentimeters converts the SquareMillimeters into SquareCentimeters.
// It does not check to make sure that there is no overflow.
//
// Required to satisfy the Area interface.
func (s SquareMillimeters) SquareCentimeters() SquareCentimeters {
	return SquareCentimeters(float64(s) / 100.0)
}

// SquareInches converts the square millimeters into SquareInches
// and returns it.
// It does not check to make sure there is no overflow.
//
// Required to satisfy the Area interface.
func (s SquareMillimeters) SquareInches() SquareInches {
	return SquareInches(float64(s) / 645.16)
}

// String converts the SquareMillimeters to a string,
// including units.
//
// I.e. "25.34 mm^2"
func (s SquareMillimeters) String() string {
	return fmt.Sprintf("%.2f mm^2", float64(s))
}
