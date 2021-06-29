package domain

type GameConfig struct {
	Name      string
	Size      FieldSize
	TickTime  int
	FoodCells int
	Fillers   []string // TODO
}

var DefaultGame = GameConfig{
	Name:      "main",
	Size:      FieldSize{SizeX: 10, SizeY: 10},
	TickTime:  500,
	FoodCells: 4,
	Fillers:   nil,
}
