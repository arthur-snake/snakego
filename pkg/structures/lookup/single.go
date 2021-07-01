package lookup

import "github.com/arthur-snake/snakego/pkg/proto"

type Single struct {
	server proto.Server
}

func NewSingle(server proto.Server) Single {
	return Single{
		server: server,
	}
}

func (s Single) Lookup(name string) proto.Server {
	return s.server
}
