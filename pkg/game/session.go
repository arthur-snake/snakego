package game

import (
	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Session struct {
	conn      *websocket.Conn
	server    *Server
	writeLock sync.Mutex
}

func NewSession(conn *websocket.Conn, server *Server) *Session {
	return &Session{
		conn:   conn,
		server: server,
	}
}

func (s *Session) SendMessage(msg domain.Message) {
	s.writeLock.Lock()
	defer s.writeLock.Unlock()

	err := s.conn.WriteMessage(websocket.TextMessage, msg.Marshal())
	if err != nil {
		log.WithError(err).Error("failed to write message")
	}
}

func (s *Session) Start() {
	// init session
	s.server.ConnectSession(s)
	defer s.server.DisconnectSession(s)

	for {
		var playerMessage map[string]string
		err := s.conn.ReadJSON(&playerMessage)
		if err != nil {
			// end session
			log.WithError(err).Error("failed to read player message")
			return
		}

		log.WithField("msg", playerMessage).Info("got message from player")

		if playerMessage["act"] == "join" {
			s.server.Join(playerMessage["nick"])
		}
	}
}
