package domain

type FieldSize struct {
	SizeX int
	SizeY int
}

func (f FieldSize) Move(l Location, d Direction) Location {
	return f.Fit(l.Add(d))
}

func (f FieldSize) Fit(l Location) Location {
	return Location{floorMod(l.X, f.SizeX), floorMod(l.Y, f.SizeY)}
}

func (f FieldSize) IsInside(l Location) bool {
	return l.X >= 0 && l.Y >= 0 && l.X < f.SizeX && l.Y <= f.SizeY
}

func floorMod(a, b int) int {
	a %= b
	if a < 0 {
		a += b
	}
	return a
}
