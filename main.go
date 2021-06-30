package main

import (
	"embed"
	"github.com/arthur-snake/snakego/pkg/lookup"
	"io/fs"
	"net/http"

	"github.com/arthur-snake/snakego/pkg/domain"
	"github.com/arthur-snake/snakego/pkg/game"
	"github.com/arthur-snake/snakego/pkg/ws"

	"github.com/arthur-snake/snakego/pkg/conf"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

//go:embed static/*
var content embed.FS

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetLevel(log.DebugLevel)

	cfg, err := conf.ParseEnv()
	if err != nil {
		log.WithError(err).Fatal("failed to parse config from env")
	}

	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		err := http.ListenAndServe(cfg.PrometheusBind, mux)
		if err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("prometheus server error")
		}
	}()

	staticFiles, err := fs.Sub(content, "static")
	if err != nil {
		log.WithError(err).Fatal("failed to found static files")
	}

	server := game.NewServer(domain.DefaultGame)
	go server.Run()

	servers := lookup.NewSingle(server)
	wsHandler := ws.NewHandler(servers)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(staticFiles)))
	mux.HandleFunc("/ws", wsHandler.Handle)

	log.WithField("bind", cfg.ServerBind).Info("starting server")
	http.ListenAndServe(cfg.ServerBind, mux)
}
