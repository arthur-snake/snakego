package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/arthur-snake/snakego/engine/game"
	"github.com/arthur-snake/snakego/pkg/structures/lookup"

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
		err2 := http.ListenAndServe(cfg.PrometheusBind, mux)
		if err2 != nil && err2 != http.ErrServerClosed {
			log.WithError(err).Fatal("prometheus server error")
		}
	}()

	staticFiles, err := fs.Sub(content, "static")
	if err != nil {
		log.WithError(err).Fatal("failed to found static files")
	}

	realServer, auto := game.NewStdServer(game.DefaultGame)
	go realServer.Run()

	servers := lookup.NewSingle(auto)
	wsHandler := ws.NewHandler(servers)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(staticFiles)))
	mux.HandleFunc("/ws", wsHandler.Handle)

	log.WithField("bind", cfg.ServerBind).Info("starting server")
	err = http.ListenAndServe(cfg.ServerBind, mux)
	log.WithError(err).Error("http server finished")
}
