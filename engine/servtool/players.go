package servtool

import (
	"sync"

	"github.com/arthur-snake/snakego/pkg/proto"
	"github.com/google/uuid"
)

type Players struct {
	players      map[uuid.UUID]*PlayerBase
	playersMutex sync.RWMutex
}

func NewPlayers() *Players {
	return &Players{
		players: make(map[uuid.UUID]*PlayerBase),
	}
}

func (p *Players) Add(player proto.Player) *PlayerBase {
	base := NewPlayerBase(player)

	p.playersMutex.Lock()
	defer p.playersMutex.Unlock()
	p.players[player.UID()] = base

	return base
}

func (p *Players) Get(player proto.Player) *PlayerBase {
	p.playersMutex.RLock()
	defer p.playersMutex.RUnlock()
	return p.players[player.UID()]
}

func (p *Players) Remove(player proto.Player) {
	p.playersMutex.Lock()
	defer p.playersMutex.Unlock()
	delete(p.players, player.UID())
}

func (p *Players) All() []*PlayerBase {
	p.playersMutex.RLock()
	defer p.playersMutex.RUnlock()

	players := make([]*PlayerBase, 0, len(p.players))
	for _, p := range p.players {
		players = append(players, p)
	}

	return players
}
