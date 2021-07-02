package test

import (
	"testing"
	"time"

	"github.com/arthur-snake/snakego/engine/game"
	"github.com/arthur-snake/snakego/pkg/domain"
)

var defaultConfig = game.Config{
	Size:      domain.FieldSize{SizeX: 75, SizeY: 40},
	TickTime:  time.Millisecond * 80,
	FoodCells: 3,
}

func TestManyServers(t *testing.T) {
	srv1, _ := game.NewStdServer(defaultConfig)
	go srv1.Run()

	srv2, _ := game.NewTickerServer(defaultConfig)
	go srv2.Run()

	srv3, _ := game.NewStdServer(defaultConfig)
	go srv3.Run()

	srv4, _ := game.NewStdServer(game.Config{
		Size:      domain.FieldSize{SizeX: 75, SizeY: 40},
		TickTime:  time.Millisecond * 100,
		FoodCells: 2,
	})
	go srv4.Run()

	srv5, _ := game.NewStdServer(game.Config{
		Size:      domain.FieldSize{SizeX: 20, SizeY: 10},
		TickTime:  time.Millisecond * 80,
		FoodCells: 1,
	})
	go srv5.Run()

	time.Sleep(time.Second * 15)
}

func BenchmarkSingleTick(b *testing.B) {
	b.ReportAllocs()

	srv, _ := game.NewStdServer(defaultConfig)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		srv.Tick()
	}
}
