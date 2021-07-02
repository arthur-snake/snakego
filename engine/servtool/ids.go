package servtool

import (
	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/proto"
)

type IDs struct {
	ids map[domain.ObjectID]proto.UpdateID

	updates []proto.UpdateID
}

func NewIDs() *IDs {
	return &IDs{
		ids: make(map[domain.ObjectID]proto.UpdateID),
	}
}

func NewBasicIDs() (ids *IDs, free, block, food proto.UpdateID) {
	ids = NewIDs()
	free = ids.Add(proto.UpdateID{
		Type:  domain.FreeCell,
		Color: domain.ClearColor,
	})
	food = ids.Add(proto.UpdateID{
		Type:  domain.FoodCell,
		Color: domain.ClearColor,
	})
	block = ids.Add(proto.UpdateID{
		Type:  domain.BlockCell,
		Color: domain.BlockColor,
	})

	return ids, free, food, block
}

func (p *IDs) Remove(id domain.ObjectID) {
	delete(p.ids, id)
}

func (p *IDs) put(id proto.UpdateID) {
	p.ids[id.ID] = id
	p.updates = append(p.updates, id)
}

func (p *IDs) Add(upd proto.UpdateID) proto.UpdateID {
	var id domain.ObjectID
	for {
		id += domain.ObjectID(generateChar())
		if _, has := p.ids[id]; !has {
			break
		}
	}

	upd.ID = id
	p.put(upd)
	return upd
}

func (p *IDs) All() []proto.UpdateID {
	var ids []proto.UpdateID
	for _, obj := range p.ids {
		ids = append(ids, obj)
	}
	return ids
}

func (p *IDs) PurgeFastUpdate() []proto.UpdateID {
	res := p.updates
	p.updates = nil

	return res
}
