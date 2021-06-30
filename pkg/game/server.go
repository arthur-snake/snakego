package game

import (
	"github.com/arthur-snake/snakego/pkg/proto"
	"sync"
	"time"

	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/draws"
)

type types struct {
	free  proto.UpdateID
	block proto.UpdateID
}

var basicTypes = types{
	free: proto.UpdateID{
		ID:    "0",
		Type:  domain.FreeCell,
		Color: domain.ClearColor,
	},
	block: proto.UpdateID{
		ID:    "1",
		Type:  domain.BlockCell,
		Color: domain.BlockColor,
	},
}

type Server struct {
	cfg       domain.GameConfig
	field     [][]domain.Cell
	slabQueue []draws.Slab

	connected   *Subscribers
	types       types
	globalMutex sync.Mutex
}

func NewServer(cfg domain.GameConfig) *Server {
	field := make([][]domain.Cell, cfg.Size.SizeX)
	for x := 0; x < cfg.Size.SizeX; x++ {
		field[x] = make([]domain.Cell, cfg.Size.SizeY)
		for y := 0; y < cfg.Size.SizeY; y++ {
			field[x][y] = domain.Cell{
				ID: basicTypes.free.ID,
			}
		}
	}

	return &Server{
		cfg:       cfg,
		field:     field,
		connected: NewSubscribers(),
		types:     basicTypes,
	}
}

func (s *Server) Connect(player proto.Player) {
	s.GlobalLock(func() {
		s.connected.Add(player)

		upd := s.buildUpdate()
		init := proto.InitMessage{
			Update: upd,
			Size:   s.cfg.Size,
		}

		player.Init(init)
	})
}

func (s *Server) Join(player proto.Player, message proto.JoinMessage) {
	s.GlobalLock(func() {
		slabs := draws.Text(s.cfg.Size.SizeY, message.Nick)
		s.slabQueue = append(s.slabQueue, slabs...)
	})
}

func (s *Server) Leave(player proto.Player, message proto.LeaveMessage) {
	panic("implement me")
}

func (s *Server) Turn(player proto.Player, message proto.TurnMessage) {
	panic("implement me")
}

func (s *Server) Chat(player proto.Player, message proto.ChatMessage) {
	panic("implement me")
}

func (s *Server) Disconnect(player proto.Player) {
	s.connected.Remove(player)
}

func (s *Server) Run() {
	ticker := time.NewTicker(s.cfg.TickTime)
	for range ticker.C {
		s.Tick()
	}
}

func (s *Server) Tick() {
	s.GlobalLock(func() {
		s.ShiftAll(domain.Directions[1])

		if len(s.slabQueue) > 0 {
			cur := s.slabQueue[0]
			s.slabQueue = s.slabQueue[1:]

			for _, y := range cur.Filled {
				s.field[s.cfg.Size.SizeX-1][y].ID = s.types.block.ID
			}
		}

		upd := s.buildUpdate()
		s.connected.BroadcastUpdate(upd)
	})
}

func (s *Server) GlobalLock(f func()) {
	s.globalMutex.Lock()
	defer s.globalMutex.Unlock()

	f()
}

func (s *Server) buildUpdate() proto.UpdateMessage {
	var cells []proto.UpdateCell
	for x := 0; x < s.cfg.Size.SizeX; x++ {
		for y := 0; y < s.cfg.Size.SizeY; y++ {
			cells = append(cells, proto.UpdateCell{
				Location: domain.Location{
					X: x,
					Y: y,
				},
				ID:   s.field[x][y].ID,
				Food: 0,
			})
		}
	}

	return proto.UpdateMessage{
		IDUpdates: []proto.UpdateID{
			s.types.free, s.types.block,
		},
		CellUpdates: cells,
		ChatUpdates: nil,
	}
}

func (s *Server) ShiftAll(dir domain.Direction) {
	iterateX := func(callback func(x int)) {
		if dir.DeltaX <= 0 {
			for x := 0; x < s.cfg.Size.SizeX; x++ {
				callback(x)
			}
		} else {
			for x := 0; x < s.cfg.Size.SizeX; x++ {
				callback(s.cfg.Size.SizeX - 1 - x)
			}
		}
	}

	iterateY := func(callback func(y int)) {
		if dir.DeltaY <= 0 {
			for y := 0; y < s.cfg.Size.SizeY; y++ {
				callback(y)
			}
		} else {
			for y := 0; y < s.cfg.Size.SizeY; y++ {
				callback(s.cfg.Size.SizeY - 1 - y)
			}
		}
	}

	iterateX(func(x int) {
		iterateY(func(y int) {
			loc := domain.Location{X: x, Y: y}
			loc = loc.Add(dir.Negate())

			if !s.cfg.Size.IsInside(loc) {
				s.field[x][y].ID = s.types.free.ID
			} else {
				s.field[x][y].ID = s.field[loc.X][loc.Y].ID
			}
		})
	})
}
