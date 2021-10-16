package units

import "fmt"

type CubicMillimeters float64

func (c CubicMillimeters) CubicMillimeters() CubicMillimeters {
	return c
}

func (c CubicMillimeters) CubicCentimeters() CubicCentimeters {
	return CubicCentimeters(float64(c) / 1000.0)
}

func (c CubicMillimeters) CubicInches() CubicInches {
	return CubicInches(float64(c) / 16387.1)
}

func (c CubicMillimeters) String() string {
	return fmt.Sprintf("%.2f mm^3", float64(c))
}
