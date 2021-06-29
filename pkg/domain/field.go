package domain

type FieldSize struct {
	SizeX int
	SizeY int
}

func (f FieldSize) Move(l Location, d Direction) Location {
	return f.Fit(Location{l.X + d.DeltaX, l.Y + d.DeltaY})
}

func (f FieldSize) Fit(l Location) Location {
	return Location{floorMod(l.X, f.SizeX), floorMod(l.Y, f.SizeY)}
}

func floorMod(a, b int) int {
	a %= b
	if a < 0 {
		a += b
	}
	return a
}
