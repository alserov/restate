package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env  string
	Addr string
	DB   Postgres
}

type Postgres struct {
	User     string
	Password string
	DB       string
	Port     int
	Host     string
}

func (p *Postgres) Dsn() string {
	return fmt.Sprintf("dsn")
}

func MustLoad() *Config {
	var cfg Config

	cfg.Addr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	cfg.Env = os.Getenv("ENV")

	return &cfg
}
