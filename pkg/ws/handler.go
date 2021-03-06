package ws

import (
	"net/http"

	"github.com/arthur-snake/snakego/pkg/proto"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// TODO: enable reasonable origin check
		return true
	},
}

type serverLookup interface {
	Lookup(name string) proto.Server
}

type Handler struct {
	lookup serverLookup
}

func NewHandler(lookup serverLookup) *Handler {
	return &Handler{
		lookup: lookup,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	server := h.lookup.Lookup(name)
	if server == nil {
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.WithError(err).Error("failed to upgrade ws")
		return
	}
	defer conn.Close()

	player := NewPlayer(conn, server)
	player.ExecuteSync()
}
