package main

import (
	"embed"
	game2 "github.com/arthur-snake/snakego/engine/game"
	lookup2 "github.com/arthur-snake/snakego/pkg/structures/lookup"
	"io/fs"
	"net/http"

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

	tickerServer, auto := game2.NewTickerServer(game2.DefaultGame)
	go tickerServer.Run()

	servers := lookup2.NewSingle(auto)
	wsHandler := ws.NewHandler(servers)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(staticFiles)))
	mux.HandleFunc("/ws", wsHandler.Handle)

	log.WithField("bind", cfg.ServerBind).Info("starting server")
	http.ListenAndServe(cfg.ServerBind, mux)
}
