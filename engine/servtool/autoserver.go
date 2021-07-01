package servtool

import (
	"github.com/arthur-snake/snakego/pkg/proto"
	"github.com/arthur-snake/snakego/pkg/structures/subs"
	log "github.com/sirupsen/logrus"
	"sync"
)

type AutoImpl interface {
	Join(*PlayerBase, proto.JoinMessage)
	Leave(*PlayerBase, proto.LeaveMessage)
	Turn(*PlayerBase, proto.TurnMessage)
	Chat(*PlayerBase, proto.ChatMessage)
}

type AutoServer struct {
	impl AutoImpl

	connected *subs.Subscribers
	players   *Players
	state     *State

	tickMutex sync.Mutex
}

func NewAutoServer(impl AutoImpl, state *State) *AutoServer {
	return &AutoServer{
		impl:      impl,
		connected: subs.NewSubscribers(),
		players:   NewPlayers(),
		state:     state,
	}
}

/// Server zone

func (s *AutoServer) TickLock(f func()) {
	s.tickMutex.Lock()
	defer s.tickMutex.Unlock()
	f()
}

func (s *AutoServer) MakeUpdate(upd StateUpdate) {
	s.TickLock(func() {
		s.connected.BroadcastUpdate(s.state.Update(upd))
	})
}

/// Client zone

func (s *AutoServer) Connect(player proto.Player) {
	log.WithField("uid", player.UID()).Info("user connected")

	s.players.Add(player)

	s.TickLock(func() {
		s.connected.Add(player)
		player.Init(s.state.Init())
	})
}

func (s *AutoServer) Disconnect(player proto.Player) {
	log.WithField("uid", player.UID()).Info("user disconnected")

	s.connected.Remove(player)

	// leave, if user is joined now
	s.Leave(player, proto.LeaveMessage{})

	s.players.Remove(player)
}

func (s *AutoServer) Join(player proto.Player, message proto.JoinMessage) {
	base := s.players.Get(player)

	log.
		WithField("uid", player.UID()).
		WithField("nick", message.Nick).
		Info("user join")

	s.impl.Join(base, message)
}

func (s *AutoServer) Leave(player proto.Player, message proto.LeaveMessage) {
	base := s.players.Get(player)

	log.
		WithField("uid", player.UID()).
		WithField("nick", base.Nick).
		Info("user leave")

	s.impl.Leave(base, message)
}

func (s *AutoServer) Turn(player proto.Player, message proto.TurnMessage) {
	base := s.players.Get(player)

	log.
		WithField("uid", player.UID()).
		WithField("nick", base.Nick).
		WithField("turn", message.Direction.Name).
		Debug("user turn")

	s.impl.Turn(base, message)
}

func (s *AutoServer) Chat(player proto.Player, message proto.ChatMessage) {
	base := s.players.Get(player)

	log.
		WithField("uid", player.UID()).
		WithField("nick", base.Nick).
		WithField("message", message.Message).
		Info("user chat")

	s.impl.Chat(base, message)
}
