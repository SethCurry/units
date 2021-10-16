package units_test

import (
	"testing"

	"github.com/SethCurry/units"
	"github.com/stretchr/testify/assert"
)

func Test_Celsius_String(t *testing.T) {
	c := units.Celsius(12.4)

	assert.Equal(t, "12.40 C", c.String())
}

func Test_Fahrenheit_String(t *testing.T) {
	f := units.Fahrenheit(32.12)

	assert.Equal(t, "32.12 F", f.String())
}

func Test_Celsius_Fahrenheit(t *testing.T) {
	c := units.Celsius(87.12)

	f := c.Fahrenheit()

	assert.Equal(t, 118.816, float32(f))
}
