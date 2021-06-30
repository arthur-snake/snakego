package ws

import (
	"encoding/json"
	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/proto"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Player struct {
	conn      *websocket.Conn
	server    proto.Server
	writeLock sync.Mutex
}

func NewPlayer(conn *websocket.Conn, server proto.Server) *Player {
	return &Player{
		conn:   conn,
		server: server,
	}
}

func (p *Player) Init(message proto.InitMessage) {
	p.send(convertInitMessage(message))
}

func (p *Player) Update(message proto.UpdateMessage) {
	p.send(convertUpdateMessage(message))
}

func (p *Player) send(msg interface{}) {
	err := p.sendError(msg)
	if err != nil {
		log.WithError(err).Error("failed to send message")
	}
}

func (p *Player) sendError(msg interface{}) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	p.writeLock.Lock()
	defer p.writeLock.Unlock()

	err = p.conn.WriteMessage(websocket.TextMessage, b)
	if err != nil {
		return err
	}

	return nil
}

func (p *Player) ExecuteSync() {
	// init session
	p.server.Connect(p)
	defer p.server.Disconnect(p)

	for {
		var playerMessage map[string]string
		err := p.conn.ReadJSON(&playerMessage)
		if err != nil {
			// end session on error
			log.WithError(err).Error("failed to read player message")
			return
		}

		p.handleMessage(playerMessage)
	}
}

func (p *Player) handleMessage(m map[string]string) {
	defer func() {
		if r := recover(); r != nil {
			log.WithField("r", r).Error("recovered from panic in player handler")
		}
	}()

	log.WithField("msg", m).Info("got message from player")

	switch m["act"] {
	case "join":
		p.server.Join(p, proto.JoinMessage{
			Nick: m["nick"],
		})

	case "leave":
		p.server.Leave(p, proto.LeaveMessage{})

	case "turn":
		dir, ok := domain.ParseDirection(m["dir"])
		if !ok {
			log.WithField("msg", m).Error("unknown direction")
			return
		}

		p.server.Turn(p, proto.TurnMessage{
			Direction: dir,
		})

	case "chat":
		p.server.Chat(p, proto.ChatMessage{
			Message: m["msg"],
		})
	}
}
