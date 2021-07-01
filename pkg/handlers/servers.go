package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arthur-snake/snakego/pkg/structures/lookup"
)

type ServerInfo struct {
	Name string
}

func ShowServers(servers *lookup.Many) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		all := servers.All()
		var infos []ServerInfo
		for name := range all {
			infos = append(infos, ServerInfo{Name: name})
		}

		_ = json.NewEncoder(w).Encode(infos)
	}
}
