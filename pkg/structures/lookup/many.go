package lookup

import (
	"github.com/arthur-snake/snakego/pkg/proto"
)

type Many struct {
	servers map[string]proto.Server
}

func NewMany() *Many {
	return &Many{
		servers: map[string]proto.Server{},
	}
}

func (m *Many) Lookup(name string) proto.Server {
	return m.servers[name]
}

func (m *Many) Add(name string, server proto.Server) {
	m.servers[name] = server
}

func (m *Many) All() map[string]proto.Server {
	return m.servers
}
