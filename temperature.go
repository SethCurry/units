package units

import "fmt"

// Temperature defines values that can be used as temperatures
// in calculations.
type Temperature interface {
	Unit
	Fahrenheit() Fahrenheit
	Celsius() Celsius
}

// Celsius stores a temperature in Celsius as a float32.
// Can be converted to Farenheit via the Farenheit() method.
type Celsius float32

// Fahrenheit converts the Celsius temperature to Fahrenheit
// and returns it.
func (c Celsius) Fahrenheit() Fahrenheit {
	return Fahrenheit((float32(c) * 1.8) + 32)
}

// Celsius simple returns the instance of Celsius.
// Required to implement the Temperature interface.
func (c Celsius) Celsius() Celsius {
	return c
}

// String converts the temperature into a string, using 2 significant
// digits and C for the units.
//
// I.e. "12.54 C"
func (c Celsius) String() string {
	return fmt.Sprintf("%.2f C", float32(c))
}

// Fahrenheit stores a temperature in Fahrenheit as a float32.
// Can be converted to Celsius via the Celsius() method.
type Fahrenheit float32

// Celsius converts the Fahrenheit temperature to Celsius
// and returns it.
func (f Fahrenheit) Celsius() Celsius {
	return Celsius((float32(f) - 32) * float32(5) / float32(9))
}

// Fahrenheit just returns a copy of itself.  Used to satisfy the
// Temperature interface.
func (f Fahrenheit) Fahrenheit() Fahrenheit {
	return f
}

// String converts the temperature into a string, using 2 significant
// digits and F for the units.
//
// I.e.  "72.18 F"
func (f Fahrenheit) String() string {
	return fmt.Sprintf("%.2f F", float32(f))
}
