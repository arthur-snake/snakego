package main

import (
	"github.com/arthur-snake/snakego/pkg/game"
	"github.com/arthur-snake/snakego/pkg/ws"
	"net/http"

	"github.com/arthur-snake/snakego/pkg/conf"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

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

	server := game.NewServer()

	h := ws.NewHandler(server)
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.Handle)

	log.WithField("bind", cfg.ServerBind).Info("starting server")
	http.ListenAndServe(cfg.ServerBind, mux)
}
