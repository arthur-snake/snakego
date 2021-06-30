package domain

type Location struct {
	X int
	Y int
}

func (l Location) Add(dir Direction) Location {
	return Location{
		X: l.X + dir.DeltaX,
		Y: l.Y + dir.DeltaY,
	}
}
