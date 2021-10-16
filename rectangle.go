package units

type Rectangle struct {
	Height Length
	Width  Length
}

func (r Rectangle) Area() Area {
	return SquareMillimeters(
		float64(r.Height.Millimeters()) * float64(r.Width.Millimeters()),
	)
}
