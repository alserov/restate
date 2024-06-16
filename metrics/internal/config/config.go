package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env  string
	Addr string

	Broker Kafka
}

type Kafka struct {
	Addr string
}

func MustLoad() *Config {
	var cfg Config

	cfg.Addr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	cfg.Env = os.Getenv("ENV")
	cfg.Broker.Addr = os.Getenv("KAFKA_ADDR")

	return &cfg
}
