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
