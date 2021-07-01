package subs

import (
	"sync"

	"github.com/arthur-snake/snakego/pkg/proto"
)

type Subscribers struct {
	list []proto.Player
	lock sync.Mutex
}

func NewSubscribers() *Subscribers {
	return &Subscribers{}
}

func (s *Subscribers) Add(p proto.Player) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.list = append(s.list, p)
}

func (s *Subscribers) Remove(p proto.Player) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for i := 0; i < len(s.list); i++ {
		if s.list[i] == p {
			s.list[i] = s.list[len(s.list)-1]
			s.list = s.list[:len(s.list)-1]
			return
		}
	}
}

func (s *Subscribers) GetAll() []proto.Player {
	s.lock.Lock()
	defer s.lock.Unlock()

	var all []proto.Player
	all = append(all, s.list...)
	return all
}

func (s *Subscribers) Broadcast(f func(proto.Player)) {
	all := s.GetAll()
	for _, it := range all {
		f(it)
	}
}

func (s *Subscribers) BroadcastUpdate(upd proto.UpdateMessage) {
	s.Broadcast(func(player proto.Player) {
		player.Update(upd)
	})
}
