package main

import (
	"fmt"

	"github.com/gasuhwbab/url-shortener/internal/config"
)

func main() {
	// config
	cfg := config.MustLoad()
	fmt.Println(cfg)
	// logger

	// storage

	// server
}
