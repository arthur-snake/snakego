package domain

type Direction struct {
	ID     int
	DeltaX int
	DeltaY int
}

func (d Direction) Negate() Direction {
	return Directions[d.ID^2]
}

func GetDirection(id int) Direction {
	return Directions[id]
}

var Directions = []Direction{
	{0, 0, -1},
	{1, -1, 0},
	{2, 0, 1},
	{3, 1, 0},
}
