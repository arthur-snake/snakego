package servtool

import (
	"github.com/arthur-snake/snakego/engine/maptool"
	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/proto"
)

type StateUpdate struct {
	NewMap       [][]domain.ObjectID
	NewIDs       []proto.UpdateID
	ChatMessages []proto.UpdateChat
}

type State struct {
	size      domain.FieldSize
	latestMap [][]domain.ObjectID
	latestIDs map[domain.ObjectID]proto.UpdateID
}

func NewState(size domain.FieldSize, upd StateUpdate) *State {
	field := maptool.CreateMap(size, "")
	for x := 0; x < size.SizeX; x++ {
		for y := 0; y < size.SizeY; y++ {
			field[x][y] = upd.NewMap[x][y]
		}
	}

	return &State{
		size:      size,
		latestMap: field,
		latestIDs: convertIDsToMap(upd.NewIDs),
	}
}

func (s *State) Init() proto.InitMessage {
	var cells []proto.UpdateCell
	for x := 0; x < s.size.SizeX; x++ {
		for y := 0; y < s.size.SizeY; y++ {
			cells = append(cells, proto.UpdateCell{
				Location: domain.Pair{X: x, Y: y},
				ID:       s.latestMap[x][y],
			})
		}
	}

	return proto.InitMessage{
		Update: &proto.UpdateMessage{
			IDUpdates:   convertIDsFromMap(s.latestIDs),
			CellUpdates: cells,
		},
		Size: &s.size,
	}
}

func (s *State) Update(upd StateUpdate) proto.UpdateMessage {
	var cells []proto.UpdateCell
	for x := 0; x < s.size.SizeX; x++ {
		for y := 0; y < s.size.SizeY; y++ {
			before := s.latestMap[x][y]
			after := upd.NewMap[x][y]

			if after != before {
				s.latestMap[x][y] = after
				cells = append(cells, proto.UpdateCell{
					Location: domain.Pair{X: x, Y: y},
					ID:       after,
				})
			}
		}
	}

	var ids []proto.UpdateID
	for _, newID := range upd.NewIDs {
		if s.latestIDs[newID.ID] != newID {
			ids = append(ids, newID)
		}
	}
	s.latestIDs = convertIDsToMap(upd.NewIDs)

	var chatMessages []proto.UpdateChat
	chatMessages = append(chatMessages, upd.ChatMessages...)

	return proto.UpdateMessage{
		IDUpdates:   ids,
		CellUpdates: cells,
		ChatUpdates: chatMessages,
	}
}

func convertIDsToMap(ids []proto.UpdateID) map[domain.ObjectID]proto.UpdateID {
	res := make(map[domain.ObjectID]proto.UpdateID)
	for _, val := range ids {
		res[val.ID] = val
	}
	return res
}

func convertIDsFromMap(m map[domain.ObjectID]proto.UpdateID) []proto.UpdateID {
	var res []proto.UpdateID
	for _, upd := range m {
		res = append(res, upd)
	}
	return res
}
