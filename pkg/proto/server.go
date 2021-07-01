package proto

import "github.com/arthur-snake/snakego/pkg/domain"

// Server is an interface for player to use.
// All calls must be sequential, i.e. parallel calls
// are not allowed.
type Server interface {
	// Connect must be called once and must be the first call.
	Connect(Player)

	Join(Player, JoinMessage)
	Leave(Player, LeaveMessage)
	Turn(Player, TurnMessage)
	Chat(Player, ChatMessage)

	// Disconnect must be called once and must be the last call.
	Disconnect(Player)
}

type JoinMessage struct {
	Nick string
}

type LeaveMessage struct{}

type TurnMessage struct {
	Direction domain.BaseDirection
}

type ChatMessage struct {
	Message string
}
