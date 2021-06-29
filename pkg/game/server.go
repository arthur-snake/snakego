package game

import (
	"github.com/arthur-snake/snakego/pkg/domain"
	"sync"
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
	rows    int
	columns int
	field   [][]domain.CellUpdate

	connected   *Subscribers
	types       types
	globalMutex sync.Mutex
}

func NewServer() *Server {
	rows := 10
	columns := 20

	field := make([][]domain.CellUpdate, columns)
	for x := 0; x < columns; x++ {
		field[x] = make([]domain.CellUpdate, rows)
		for y := 0; y < rows; y++ {
			field[x][y] = domain.CellUpdate{
				X:  x,
				Y:  y,
				ID: basicTypes.free.ID,
			}
		}
	}

	return &Server{
		rows:      rows,
		columns:   columns,
		field:     field,
		connected: NewSubscribers(),
		types:     basicTypes,
	}
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
	for x := 0; x < s.columns; x++ {
		for y := 0; y < s.rows; y++ {
			cells = append(cells, s.field[x][y])
		}
	}

	return domain.InitMessage{
		Act:        "init",
		MapUpdates: domain.MarshalMapUpdates(cells),
		IDUpdates:  []domain.IDInfo{s.types.free, s.types.block},
		Rows:       s.rows,
		Columns:    s.columns,
	}
}
