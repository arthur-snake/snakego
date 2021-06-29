package game

import "sync"

type Subscriber interface{}

type Subscribers struct {
	list []Subscriber
	lock sync.Mutex
}

func NewSubscribers() *Subscribers {
	return &Subscribers{}
}

func (s *Subscribers) Add(sub Subscriber) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.list = append(s.list, sub)
}

func (s *Subscribers) Remove(sub Subscriber) {
	s.lock.Lock()
	defer s.lock.Unlock()

	for i := 0; i < len(s.list); i++ {
		if s.list[i] == sub {
			s.list[i] = s.list[len(s.list)-1]
			s.list = s.list[:len(s.list)-1]
			return
		}
	}
}

func (s *Subscribers) GetAll() []Subscriber {
	s.lock.Lock()
	defer s.lock.Unlock()

	var all []Subscriber
	all = append(all, s.list...)
	return all
}
