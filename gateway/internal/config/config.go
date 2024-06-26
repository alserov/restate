package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env  string
	Addr string

	Services Services

	Broker Kafka
}

type Services struct {
	Estate   string
	Meetings string
}

type Kafka struct {
	Addr   string
	Topics struct {
		Metrics string
	}
}

func MustLoad() *Config {
	var cfg Config

	cfg.Addr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	cfg.Env = os.Getenv("ENV")

	// Broker
	cfg.Broker.Addr = os.Getenv("KAFKA_ADDR")

	// Services
	cfg.Services.Meetings = os.Getenv("MEETINGS_ADDR")
	cfg.Services.Estate = os.Getenv("ESTATE_ADDR")

	return &cfg
}
