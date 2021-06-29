package ws

import (
	"github.com/arthur-snake/snakego/pkg/game"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Handler struct {
	server *game.Server
}

func NewHandler(server *game.Server) *Handler {
	return &Handler{
		server: server,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithError(err).Error("failed to upgrade ws")
		return
	}
	defer conn.Close()

	session := game.NewSession(conn, h.server)
	session.Start()
}
