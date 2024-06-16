package app

import (
	"context"
	"github.com/alserov/restate/metrics/internal/config"
	"github.com/alserov/restate/metrics/internal/log"
	"github.com/alserov/restate/metrics/internal/workers"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

func MustStart(cfg *config.Config) {
	lg := log.NewLogger(cfg.Env, log.KindZap)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	run(func() {
		lg.Info("starting server", nil)
		go workers.NewWorker(workers.System, cfg.Broker.Addr).Run(log.WithLogger(ctx, lg))
	})

	lg.Info("shutdown server", nil)
}

func run(fn func()) {
	go fn()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
}
