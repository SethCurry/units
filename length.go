package units

type Centimeters float64
type Inches float64
type Feet float64
type Yards float64
type Meters float64

// Length defines an interface for objects that can be used
// as lengths.
type Length interface {
	Unit
	Millimeters() Millimeters
	Centimeters() Centimeters
	Inches() Inches
	Feet() Feet
	Yards() Yards
	Meters() Meters
}
