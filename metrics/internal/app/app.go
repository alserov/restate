package app

import (
	"context"
	"github.com/alserov/restate/metrics/internal/async"
	"github.com/alserov/restate/metrics/internal/config"
	"github.com/alserov/restate/metrics/internal/log"
	"github.com/alserov/restate/metrics/internal/workers"
	_ "github.com/joho/godotenv/autoload"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os/signal"
	"syscall"
)

const (
	systemWorkersAmount   = 5
	businessWorkersAmount = 3
)

func MustStart(cfg *config.Config) {
	lg := log.NewLogger(cfg.Env, log.KindZap)
	reg := prometheus.NewRegistry()

	// workers init
	var collectors []prometheus.Collector

	systemWorker := workers.NewWorker(
		workers.System,
		async.NewConsumer(async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.Metrics),
		&collectors,
	)

	businessWorker := workers.NewWorker(
		workers.Business,
		async.NewConsumer(async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.Metrics),
		&collectors,
	)

	reg.MustRegister(collectors...)

	// endpoint for prometheus
	m := http.NewServeMux()
	m.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = log.WithLogger(ctx, lg)

	run(func() {
		systemWorker.Run(ctx, systemWorkersAmount)
		businessWorker.Run(ctx, businessWorkersAmount)

		lg.Info("server is running", nil)

		if err := http.ListenAndServe(cfg.Addr, m); err != nil {
			panic("failed to serve: " + err.Error())
		}
	})

	lg.Info("shutdown server", nil)
}

func run(fn func()) {
	go fn()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
}
