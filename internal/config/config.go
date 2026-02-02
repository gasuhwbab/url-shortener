package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env          string `yaml:"env" env-default:"local"`
	Storage_path string `yaml:"storage_path" env-required:"true"`
	HTTPServer   `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"10s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"100s"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	confPath := os.Getenv("CONFIG_PATH")
	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		panic("ERROR: config doesn't exist")
	}
	var cfg Config
	if err := cleanenv.ReadConfig(confPath, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
