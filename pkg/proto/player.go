package proto

import (
	"github.com/arthur-snake/snakego/pkg/domain"
)

type Player interface {
	Init(InitMessage)
	Update(UpdateMessage)
}

// UpdateMessage contains all updates.
type UpdateMessage struct {
	IDUpdates   []UpdateID
	CellUpdates []UpdateCell
	ChatUpdates []UpdateChat
}

// InitMessage contains all info to initialize player.
type InitMessage struct {
	Update UpdateMessage
	Size   domain.FieldSize
}

// UpdateCell contains update for ID and Food at specified Location.
type UpdateCell struct {
	Location domain.Location

	ID   domain.ObjectID
	Food int
}

type UpdateID struct {
	ID domain.ObjectID

	Type  domain.CellType
	Color domain.Color
	Nick  string
}

type UpdateChat struct {
	ID      domain.ObjectID
	Message string
}
