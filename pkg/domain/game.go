package domain

import "time"

type GameConfig struct {
	Name      string
	Size      FieldSize
	TickTime  time.Duration
	FoodCells int
	Fillers   []string // TODO
}

var DefaultGame = GameConfig{
	Name:      "main",
	Size:      FieldSize{SizeX: 30, SizeY: 15},
	TickTime:  time.Millisecond * 200,
	FoodCells: 4,
	Fillers:   nil,
}
