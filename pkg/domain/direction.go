package domain

type Direction struct {
	ID     int
	DeltaX int
	DeltaY int
	Name   string
}

func (d Direction) Negate() Direction {
	return Directions[d.ID^2]
}

func GetDirection(id int) Direction {
	return Directions[id]
}

// TODO: add base directions
var Directions = []Direction{
	{0, 0, -1, "UP"},
	{1, -1, 0, "LEFT"},
	{2, 0, 1, "DOWN"},
	{3, 1, 0, "RIGHT"},
}

func ParseDirection(str string) (Direction, bool) {
	for _, dir := range Directions {
		if dir.Name == str {
			return dir, true
		}
	}

	return Direction{}, false
}
