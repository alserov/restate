package app

import (
	"context"
	"github.com/alserov/restate/metrics/internal/config"
	"github.com/alserov/restate/metrics/internal/log"
	"github.com/alserov/restate/metrics/internal/workers"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	lg := log.NewLogger(cfg.Env, log.KindZap)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	run(func() {
		go workers.NewWorker(workers.System).Run(log.WithLogger(ctx, lg))
	})

	lg.Info("shutdown server", nil)
}

func run(fn func()) {
	go fn()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
}
