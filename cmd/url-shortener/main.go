package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gasuhwbab/url-shortener/internal/config"
	"github.com/gasuhwbab/url-shortener/internal/logger"
	"github.com/gasuhwbab/url-shortener/internal/server/handlers/url/save"
	"github.com/gasuhwbab/url-shortener/internal/storage/sqlite"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Post("/", save.New(log, storage))

	// server
	log.Info("starting server", slog.String("address", cfg.Address))
	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server", logger.ErrorAttr(err))
	}
}
