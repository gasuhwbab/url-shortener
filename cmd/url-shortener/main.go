package main

import (
	"os"

	"github.com/gasuhwbab/url-shortener/internal/config"
	"github.com/gasuhwbab/url-shortener/internal/logger"
	"github.com/gasuhwbab/url-shortener/internal/storage/sqlite"
)

func main() {
	// config
	cfg := config.MustLoad()
	// logger
	log := logger.SetupLogger(cfg.Env)
	// storage
	storage, err := sqlite.New(cfg.Storage_path)
	if err != nil {
		log.Error("failed to init storage", logger.ErrorAttr(err))
		os.Exit(1)
	}

	_ = storage
	// server
}
