package proto

import "github.com/arthur-snake/snakego/pkg/domain"

type Server interface {
	Connect(Player)

	Join(Player, JoinMessage)
	Leave(Player, LeaveMessage)
	Turn(Player, TurnMessage)
	Chat(Player, ChatMessage)

	Disconnect(Player)
}

type JoinMessage struct {
	Nick string
}

type LeaveMessage struct{}

type TurnMessage struct {
	Direction domain.Direction
}

type ChatMessage struct {
	Message string
}
