package test

import (
	"testing"
	"time"

	"github.com/arthur-snake/snakego/engine/game"
	"github.com/arthur-snake/snakego/pkg/domain"
)

func TestName(t *testing.T) {
	def := game.Config{
		Size:      domain.FieldSize{SizeX: 75, SizeY: 40},
		TickTime:  80,
		FoodCells: 3,
	}

	srv1, _ := game.NewStdServer(def)
	go srv1.Run()

	srv2, _ := game.NewTickerServer(def)
	go srv2.Run()

	srv3, _ := game.NewStdServer(def)
	go srv3.Run()

	srv4, _ := game.NewStdServer(game.Config{
		Size:      domain.FieldSize{SizeX: 75, SizeY: 40},
		TickTime:  100,
		FoodCells: 2,
	})
	go srv4.Run()

	srv5, _ := game.NewStdServer(game.Config{
		Size:      domain.FieldSize{SizeX: 20, SizeY: 10},
		TickTime:  80,
		FoodCells: 1,
	})
	go srv5.Run()

	time.Sleep(time.Second * 15)
}
