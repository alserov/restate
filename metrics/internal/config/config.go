package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env  string
	Addr string
}

func MustLoad() *Config {
	var cfg Config

	cfg.Addr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	cfg.Env = os.Getenv("ENV")

	return &cfg
}
