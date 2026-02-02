package main

import (
	"github.com/gasuhwbab/url-shortener/internal/config"
	"github.com/gasuhwbab/url-shortener/internal/logger"
)

func main() {
	// config
	cfg := config.MustLoad()
	// logger
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting app")
	// storage

	// server
}
