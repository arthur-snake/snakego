package game

import (
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/arthur-snake/snakego/engine/maptool"
	"github.com/arthur-snake/snakego/engine/servtool"
	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/proto"
)

const (
	minPlayerLength = 3
)

type StdServer struct {
	mutex sync.Mutex
	field [][]domain.Cell
	ids   *servtool.IDs

	cfg  Config
	auto *servtool.AutoServer

	freeID domain.ObjectID
	foodID domain.ObjectID
}

//nolint:dupl
func NewStdServer(cfg Config) (*StdServer, *servtool.AutoServer) {
	ids, free, food, _ := servtool.NewBasicIDs()

	field := maptool.CreateMap(cfg.Size, domain.Cell{ID: free.ID})
	state := servtool.NewState(cfg.Size, servtool.StateUpdate{
		NewMap: field,
		NewIDs: ids.All(),
	})

	srv := &StdServer{
		cfg:    cfg,
		field:  field,
		ids:    ids,
		freeID: free.ID,
		foodID: food.ID,
	}

	auto := servtool.NewAutoServer(srv, state)
	srv.auto = auto

	return srv, auto
}

// Player interface

func (s *StdServer) Join(base *servtool.PlayerBase, message proto.JoinMessage) {
	const (
		maxNickLength = 15
		defaultNick   = "snake"
	)

	s.mutex.Lock()
	defer s.mutex.Unlock()

	if base.InGame {
		// already joined
		return
	}

	nick := strings.TrimSpace(message.Nick)
	if len(nick) > maxNickLength {
		nick = nick[:maxNickLength]
	}

	if nick == "" {
		nick = defaultNick
	}

	var loc domain.Pair
	distr := maptool.DistributeMap(s.cfg.Size, s.field)
	if len(distr[s.freeID]) > 0 {
		loc = distr[s.freeID][0].Location
	} else if len(distr[s.foodID]) > 0 {
		loc = distr[s.foodID][0].Location
	} else {
		// no place
		return
	}

	// proceed the join
	base.InGame = true
	base.Nick = nick
	base.Color = domain.GenerateColor()
	base.ObjectID = s.ids.Add(proto.UpdateID{
		Type:  domain.PlayerCell,
		Color: base.Color,
		Nick:  base.Nick,
	}).ID
	base.Cells = []domain.Pair{loc}
	base.Stock = minPlayerLength - 1
	base.Controller = servtool.NewController()

	s.field[loc.X][loc.Y] = domain.Cell{
		ID: base.ObjectID,
	}
}

func (s *StdServer) Leave(base *servtool.PlayerBase, message proto.LeaveMessage) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !base.InGame {
		// not in game
		return
	}

	distr := maptool.DistributeMap(s.cfg.Size, s.field)
	for _, ptr := range distr[base.ObjectID] {
		*ptr.Cell = domain.Cell{ID: s.freeID}
	}

	s.ids.Remove(base.ObjectID)
	base.ObjectID = ""

	base.InGame = false
	base.Cells = []domain.Pair{}
}

func (s *StdServer) Turn(base *servtool.PlayerBase, message proto.TurnMessage) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !base.InGame {
		// not in game
		return
	}

	base.Controller.Turn(message.Direction)
}

func (s *StdServer) Chat(base *servtool.PlayerBase, message proto.ChatMessage) {
	// TODO:
}

// Server methods

func (s *StdServer) Run() {
	ticker := time.NewTicker(s.cfg.TickTime)
	for range ticker.C {
		s.tick()
	}
}

func (s *StdServer) tick() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	players := s.playersByLength()

	for _, p := range players {
		s.moveHead(p)
	}

	for _, p := range players {
		s.moveTail(p)
	}

	s.placeFood()

	s.auto.MakeUpdate(servtool.StateUpdate{
		NewMap: s.field,
		NewIDs: s.ids.All(),
	})
}

func (s *StdServer) moveHead(p *servtool.PlayerBase) {
	dir := p.Controller.Move().Dir

	head := p.Cells[p.Length()-1]
	nxt := s.cfg.Size.Move(head, dir)
	cell := &s.field[nxt.X][nxt.Y]

	if cell.ID == s.foodID {
		p.Stock += cell.Food
	} else if cell.ID != s.freeID {
		// cannot join
		return
	}

	*cell = domain.Cell{
		ID: p.ObjectID,
	}
	p.Cells = append(p.Cells, nxt)
}

func (s *StdServer) moveTail(p *servtool.PlayerBase) {
	if p.Stock > 0 {
		p.Stock--
		return
	}

	if p.Length() <= minPlayerLength {
		return
	}

	tail := p.Cells[0]
	p.Cells = p.Cells[1:]

	s.field[tail.X][tail.Y] = domain.Cell{
		ID: s.freeID,
	}
}

func (s *StdServer) playersByLength() []*servtool.PlayerBase {
	all := s.auto.Players()

	var joined []*servtool.PlayerBase
	for _, player := range all {
		if player.InGame {
			joined = append(joined, player)
		}
	}

	sort.Slice(joined, func(i, j int) bool {
		return joined[i].Length() < joined[j].Length()
	})

	return joined
}

func (s *StdServer) placeFood() {
	distr := maptool.DistributeMap(s.cfg.Size, s.field)
	for i := 0; i < len(distr[s.freeID]); i++ {
		if len(distr[s.foodID])+i+1 > s.cfg.FoodCells {
			break
		}

		ptr := distr[s.freeID][i]
		*ptr.Cell = domain.Cell{
			ID:   s.foodID,
			Food: rand.Intn(9) + 1, //nolint:gomnd,gosec
		}
	}
}
