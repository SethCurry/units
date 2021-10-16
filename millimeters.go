package units

import "fmt"

// Ensure that Millimeters satisfies the Length interface.
var _ Length = Millimeters(0)

// Millimeters stores a length in Millimeters.
// It complies with the Length interface.
type Millimeters float64

// Millimeters simply returns a copy of itself.
// Required to implement the Length interface.
func (m Millimeters) Millimeters() Millimeters {
	return m
}

// Centimeters converts the millimeters to centimeters
// and returns the value.
func (m Millimeters) Centimeters() Centimeters {
	return Centimeters(float64(m) / 10.0)
}

// Meters converts the millimeters to meters
// and returns the value.
func (m Millimeters) Meters() Meters {
	return Meters(float64(m) / 1000.0)
}

// Inches converts the millimeters to inches
// and returns the value.
func (m Millimeters) Inches() Inches {
	return Inches(float64(m) / 25.4)
}

// Feet converts the millimeters to feet
// and returns the value.
func (m Millimeters) Feet() Feet {
	return Feet(float64(m) / 305)
}

// Yards converts the millimeters to yards
// and returns the value.
func (m Millimeters) Yards() Yards {
	return Yards(float64(m) / 914)
}

// String converts the millimeters to a string
// with 2 decimals and append the units.
//
// I.e. 53.27 mm
func (m Millimeters) String() string {
	return fmt.Sprintf("%.2f mm", float64(m))
}
