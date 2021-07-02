package maptool

import (
	"math/rand"

	"github.com/arthur-snake/snakego/pkg/domain"
)

type RandomSet struct {
	pos map[domain.Pair]int
	all []domain.CellWithLocation
}

func NewRandomSet() *RandomSet {
	return &RandomSet{
		pos: map[domain.Pair]int{},
	}
}

func (s *RandomSet) Add(cell domain.CellWithLocation) {
	key := cell.Location
	_, ok := s.pos[key]
	if ok {
		// already here
		return
	}

	s.pos[key] = len(s.all)
	s.all = append(s.all, cell)
}

func (s *RandomSet) Remove(cell domain.CellWithLocation) {
	key := cell.Location
	pos, ok := s.pos[key]
	if !ok {
		// not there
		return
	}

	s.swap(pos, len(s.pos)-1)
	delete(s.pos, key)
	s.all = s.all[:len(s.all)-1]
}

func (s *RandomSet) swap(i, j int) {
	keyI, keyJ := s.all[i].Location, s.all[j].Location

	s.pos[keyI], s.pos[keyJ] = j, i
	s.all[i], s.all[j] = s.all[j], s.all[i]
}

func (s *RandomSet) Len() int {
	return len(s.all)
}

func (s *RandomSet) Random() (domain.CellWithLocation, bool) {
	if len(s.all) == 0 {
		return domain.CellWithLocation{}, false
	}

	i := rand.Intn(len(s.all)) //nolint:gosec
	return s.all[i], true
}

func (s *RandomSet) PurgeAll() []domain.CellWithLocation {
	res := s.all

	s.pos = map[domain.Pair]int{}
	s.all = nil

	return res
}
