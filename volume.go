package units

// Volume stores the size of a volume, and allows
// converting between different units.
type Volume interface {
	Unit
	CubicMillimeters() CubicMillimeters
	CubicCentimeters() CubicCentimeters
	CubicInches() CubicInches
}
