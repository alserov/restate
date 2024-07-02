package app

import (
	"context"
	"github.com/alserov/restate/metrics/internal/async"
	"github.com/alserov/restate/metrics/internal/config"
	"github.com/alserov/restate/metrics/internal/log"
	"github.com/alserov/restate/metrics/internal/workers"
	"github.com/alserov/restate/metrics/pkg/models"
	_ "github.com/joho/godotenv/autoload"
	"os/signal"
	"syscall"
)

const (
	systemWorkers = 5
)

func MustStart(cfg *config.Config) {
	lg := log.NewLogger(cfg.Env, log.KindZap)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	run(func() {
		go workers.NewWorker(
			workers.System,
			async.NewConsumer(async.Kafka, cfg.Broker.Addr, models.TopicMetrics),
		).Run(log.WithLogger(ctx, lg), systemWorkers)

		lg.Info("server is running", nil)
	})

	lg.Info("shutdown server", nil)
}

func run(fn func()) {
	go fn()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
}
