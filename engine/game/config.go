package game

import (
	"github.com/arthur-snake/snakego/pkg/domain"
	"time"
)

type Config struct {
	Size      domain.FieldSize
	TickTime  time.Duration
	FoodCells int

	// TODO:
	//  - name
	//  - fillers []string
}

var DefaultGame = Config{
	Size:      domain.FieldSize{SizeX: 30, SizeY: 15},
	TickTime:  time.Millisecond * 200,
	FoodCells: 4,
}
