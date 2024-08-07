package config

import (
	"fmt"
	"os"
)

type Config struct {
	Env  string
	Addr string
	DB   Postgres

	Broker Kafka
}

type Kafka struct {
	Addr   string
	Topics struct {
		Metrics string
	}
}

type Postgres struct {
	User     string
	Password string
	DB       string
	Port     string
	Host     string
}

func (p *Postgres) Dsn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.User, p.Password, p.Host, p.Port, p.DB)
}

func MustLoad() *Config {
	var cfg Config

	cfg.Addr = fmt.Sprintf(":%s", os.Getenv("PORT"))
	cfg.Env = os.Getenv("ENV")

	// DB
	cfg.DB = Postgres{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		Host:     os.Getenv("DB_HOST"),
	}
	cfg.Broker.Addr = os.Getenv("KAFKA_ADDR")

	// Broker
	cfg.Broker.Addr = os.Getenv("KAFKA_ADDR")
	cfg.Broker.Topics.Metrics = os.Getenv("TOPIC_METRICS")

	return &cfg
}
