package game

import (
	"time"

	"github.com/arthur-snake/snakego/pkg/domain"
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
	Size:      domain.FieldSize{SizeX: 60, SizeY: 30},
	TickTime:  time.Millisecond * 100,
	FoodCells: 4,
}
