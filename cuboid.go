package units

type Cuboid struct {
	Length Length
	Width  Length
	Height Length
}

func (c Cuboid) Volume() Volume {
	return CubicMillimeters(
		float64(c.Length.Millimeters()) * float64(c.Height.Millimeters()) * float64(c.Width.Millimeters()),
	)
}
