package main

import (
	"embed"
	"io/fs"
	"net/http"
	"time"

	"github.com/arthur-snake/snakego/pkg/domain"

	"github.com/arthur-snake/snakego/engine/game"
	"github.com/arthur-snake/snakego/pkg/handlers"

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

	servers := lookup.NewMany()

	def := game.Config{
		Size:      domain.FieldSize{SizeX: cfg.SizeX, SizeY: cfg.SizeY},
		TickTime:  time.Millisecond * time.Duration(cfg.TickTime),
		FoodCells: cfg.FoodCount,
	}

	srv1, auto1 := game.NewStdServer(def)
	go srv1.Run()
	servers.Add("std", auto1)
	servers.Add("", auto1)

	srv2, auto2 := game.NewTickerServer(def)
	go srv2.Run()
	servers.Add("tick", auto2)

	srv3, auto3 := game.NewStdServer(game.Config{
		Size:      domain.FieldSize{SizeX: 75, SizeY: 40},
		TickTime:  80,
		FoodCells: 3,
	})
	go srv3.Run()
	servers.Add("faster", auto3)

	srv4, auto4 := game.NewStdServer(game.Config{
		Size:      domain.FieldSize{SizeX: 75, SizeY: 40},
		TickTime:  100,
		FoodCells: 2,
	})
	go srv4.Run()
	servers.Add("slower", auto4)

	srv5, auto5 := game.NewStdServer(game.Config{
		Size:      domain.FieldSize{SizeX: 20, SizeY: 10},
		TickTime:  80,
		FoodCells: 1,
	})
	go srv5.Run()
	servers.Add("small", auto5)

	wsHandler := ws.NewHandler(servers)

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.FS(staticFiles)))
	mux.HandleFunc("/ws", wsHandler.Handle)
	mux.HandleFunc("/servers", handlers.ShowServers(servers))

	log.WithField("bind", cfg.ServerBind).Info("starting server")
	err = http.ListenAndServe(cfg.ServerBind, mux)
	log.WithError(err).Error("http server finished")
}
