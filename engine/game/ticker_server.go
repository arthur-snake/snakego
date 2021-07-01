package game

import (
	"github.com/arthur-snake/snakego/engine/maptool"
	"github.com/arthur-snake/snakego/engine/servtool"
	"github.com/arthur-snake/snakego/pkg/proto"
	"sync"
	"time"

	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/draws"
)

type TickerServer struct {
	mutex sync.Mutex

	cfg       Config
	field     [][]domain.ObjectID
	slabQueue []draws.Slab

	ids     *servtool.IDs
	freeID  domain.ObjectID
	blockID domain.ObjectID

	auto *servtool.AutoServer
}

func NewTickerServer(cfg Config) (*TickerServer, *servtool.AutoServer) {
	ids := servtool.NewIDs()
	free := ids.Add(proto.UpdateID{
		Type:  domain.FreeCell,
		Color: domain.ClearColor,
	})
	block := ids.Add(proto.UpdateID{
		Type:  domain.BlockCell,
		Color: domain.BlockColor,
	})

	field := maptool.CreateMap(cfg.Size, free.ID)
	state := servtool.NewState(cfg.Size, servtool.StateUpdate{
		NewMap: field,
		NewIDs: ids.All(),
	})

	srv := &TickerServer{
		cfg:     cfg,
		field:   field,
		ids:     ids,
		freeID:  free.ID,
		blockID: block.ID,
	}

	auto := servtool.NewAutoServer(srv, state)
	srv.auto = auto

	return srv, auto
}

/// Player interface

func (s *TickerServer) Join(base *servtool.PlayerBase, message proto.JoinMessage) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	slabs := draws.Text(s.cfg.Size.SizeY, message.Nick)
	s.slabQueue = append(s.slabQueue, slabs...)
}

func (s *TickerServer) Leave(base *servtool.PlayerBase, message proto.LeaveMessage) {

}

func (s *TickerServer) Turn(base *servtool.PlayerBase, message proto.TurnMessage) {

}

func (s *TickerServer) Chat(base *servtool.PlayerBase, message proto.ChatMessage) {

}

/// Server methods

func (s *TickerServer) Run() {
	ticker := time.NewTicker(s.cfg.TickTime)
	for range ticker.C {
		s.tick()
	}
}

func (s *TickerServer) tick() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	maptool.ShiftDir(s.cfg.Size, s.field, domain.Left.Dir, s.freeID)

	if len(s.slabQueue) > 0 {
		cur := s.slabQueue[0]
		s.slabQueue = s.slabQueue[1:]

		for _, y := range cur.Filled {
			s.field[s.cfg.Size.SizeX-1][y] = s.blockID
		}
	}

	s.auto.MakeUpdate(servtool.StateUpdate{
		NewMap: s.field,
		NewIDs: s.ids.All(),
	})
}
