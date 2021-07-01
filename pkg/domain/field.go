package domain

type FieldSize struct {
	SizeX int
	SizeY int
}

func (f FieldSize) Move(loc Pair, dir Pair) Pair {
	return f.Fit(loc.Add(dir))
}

func (f FieldSize) Fit(loc Pair) Pair {
	return Pair{floorMod(loc.X, f.SizeX), floorMod(loc.Y, f.SizeY)}
}

func (f FieldSize) IsInside(loc Pair) bool {
	return loc.X >= 0 && loc.Y >= 0 && loc.X < f.SizeX && loc.Y <= f.SizeY
}

func floorMod(a, b int) int {
	a %= b
	if a < 0 {
		a += b
	}
	return a
}
