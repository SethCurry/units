package units

// Area represents the measurement of an area.
type Area interface {
	Unit
	SquareMillimeters() SquareMillimeters
	SquareCentimeters() SquareCentimeters
	SquareInches() SquareInches
}
