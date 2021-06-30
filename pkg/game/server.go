package game

import (
	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/draws"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

type types struct {
	free  domain.IDInfo
	block domain.IDInfo
}

var basicTypes = types{
	free: domain.IDInfo{
		ID:    "0",
		Type:  domain.FreeCell,
		Color: domain.ClearColor,
	},
	block: domain.IDInfo{
		ID:    "1",
		Type:  domain.BlockCell,
		Color: domain.BlockColor,
	},
}

type Server struct {
	cfg       domain.GameConfig
	field     [][]domain.CellUpdate
	slabQueue []draws.Slab

	connected   *Subscribers
	types       types
	globalMutex sync.Mutex
}

func NewServer(cfg domain.GameConfig) *Server {
	field := make([][]domain.CellUpdate, cfg.Size.SizeX)
	for x := 0; x < cfg.Size.SizeX; x++ {
		field[x] = make([]domain.CellUpdate, cfg.Size.SizeY)
		for y := 0; y < cfg.Size.SizeY; y++ {
			field[x][y] = domain.CellUpdate{
				X:  x,
				Y:  y,
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
				log.Infof("filled cell %v %v", s.cfg.Size.SizeX-1, y)
			}
		}

		msg := s.buildInitMessage()
		msg.Act = "upd"
		s.connected.Broadcast(msg)
	})
}

func (s *Server) GlobalLock(f func()) {
	s.globalMutex.Lock()
	defer s.globalMutex.Unlock()

	f()
}

func (s *Server) ConnectSession(session *Session) {
	s.GlobalLock(func() {
		s.connected.Add(session)
		session.SendMessage(s.buildInitMessage())
	})
}

func (s *Server) DisconnectSession(session *Session) {
	s.connected.Remove(session)
}

func (s *Server) buildInitMessage() domain.InitMessage {
	var cells []domain.CellUpdate
	for x := 0; x < s.cfg.Size.SizeX; x++ {
		for y := 0; y < s.cfg.Size.SizeY; y++ {
			cells = append(cells, s.field[x][y])
		}
	}

	return domain.InitMessage{
		Act:        "init",
		MapUpdates: domain.MarshalMapUpdates(cells),
		IDUpdates:  []domain.IDInfo{s.types.free, s.types.block},
		SizeX:      s.cfg.Size.SizeX,
		SizeY:      s.cfg.Size.SizeY,
	}
}

func (s *Server) Join(text string) {
	s.GlobalLock(func() {
		slabs := draws.Text(s.cfg.Size.SizeY, text)
		s.slabQueue = append(s.slabQueue, slabs...)
	})
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
