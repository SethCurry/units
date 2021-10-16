package units

import "fmt"

type CubicInches float64

func (c CubicInches) CubicMillimeters() CubicMillimeters {
	return CubicMillimeters(float64(c) * 16387.1)
}

func (c CubicInches) CubicCentimeters() CubicCentimeters {
	return CubicCentimeters(float64(c) * 16.3871)
}

func (c CubicInches) CubicInches() CubicInches {
	return c
}

func (c CubicInches) String() string {
	return fmt.Sprintf("%.2f cu. in", float64(c))
}
