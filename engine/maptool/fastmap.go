package maptool

import (
	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/proto"
)

type FastMap struct {
	field [][]domain.Cell
	ids   map[domain.ObjectID]*RandomSet

	fastUpdate *RandomSet
}

func NewFastMap(size domain.FieldSize, zero domain.Cell) *FastMap {
	f := &FastMap{
		field:      CreateMap(size, zero),
		ids:        map[domain.ObjectID]*RandomSet{},
		fastUpdate: NewRandomSet(),
	}

	for x := 0; x < size.SizeX; x++ {
		for y := 0; y < size.SizeY; y++ {
			f.getSet(zero.ID).Add(f.at(domain.Pair{X: x, Y: y}))
			f.fastUpdate.Add(f.at(domain.Pair{X: x, Y: y}))
		}
	}

	return f
}

func (f *FastMap) getSet(id domain.ObjectID) *RandomSet {
	set, ok := f.ids[id]
	if !ok {
		set = NewRandomSet()
		f.ids[id] = set
	}

	return set
}

func (f *FastMap) at(p domain.Pair) domain.CellWithLocation {
	return domain.CellWithLocation{
		Location: p,
		Cell:     &f.field[p.X][p.Y],
	}
}

func (f *FastMap) Get(p domain.Pair) domain.Cell {
	return f.field[p.X][p.Y]
}

func (f *FastMap) Set(p domain.Pair, cell domain.Cell) {
	val := f.at(p)
	f.getSet(val.Cell.ID).Remove(val)
	*val.Cell = cell
	f.getSet(val.Cell.ID).Add(val)

	f.fastUpdate.Add(val)
}

func (f *FastMap) Field() [][]domain.Cell {
	return f.field
}

func (f *FastMap) AnyRandom(ids ...domain.ObjectID) (domain.Pair, bool) {
	for _, id := range ids {
		cell, ok := f.getSet(id).Random()
		if !ok {
			continue
		}

		return cell.Location, true
	}

	return domain.Pair{}, false
}

func (f *FastMap) Count(id domain.ObjectID) int {
	return f.getSet(id).Len()
}

func (f *FastMap) PurgeFastUpdate() []proto.UpdateCell {
	cells := f.fastUpdate.PurgeAll()

	var updates []proto.UpdateCell
	for _, cell := range cells {
		updates = append(updates, proto.UpdateCell{
			Location: cell.Location,
			Cell:     *cell.Cell,
		})
	}

	return updates
}
