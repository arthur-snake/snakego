package servtool

import (
	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/proto"
)

type PlayerBase struct {
	Player proto.Player

	InGame     bool
	Nick       string
	Color      domain.Color
	ObjectID   domain.ObjectID
	Cells      []domain.Pair
	Stock      int
	Controller *Controller
}

func NewPlayerBase(p proto.Player) *PlayerBase {
	return &PlayerBase{
		Player: p,
		InGame: false,
		Nick:   "snake",
	}
}

func (p *PlayerBase) Length() int {
	return len(p.Cells)
}
