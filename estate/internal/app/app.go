package app

import (
	"context"
	"github.com/alserov/restate/estate/internal/async"
	"github.com/alserov/restate/estate/internal/cache/redis"
	"github.com/alserov/restate/estate/internal/config"
	"github.com/alserov/restate/estate/internal/db/posgtres"
	"github.com/alserov/restate/estate/internal/log"
	"github.com/alserov/restate/estate/internal/metrics"
	"github.com/alserov/restate/estate/internal/server/grpc"
	"github.com/alserov/restate/estate/internal/service"
	_ "github.com/joho/godotenv/autoload"
	"net"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	lg := log.NewLogger(cfg.Env, log.KindZap)

	// external dependencies
	db := posgtres.MustConnect(cfg.DB.Dsn())
	defer func() {
		_ = db.Close()
	}()

	c := redis.MustConnect(cfg.Cache.Addr)
	defer func() {
		_ = c.Close()
	}()

	// initializing instances
	cch := redis.NewCache(c)
	metr := metrics.NewMetrics(async.NewProducer(async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.Metrics))
	repo := posgtres.NewRepository(db)
	srvc := service.NewService(repo)
	srvr := grpc.RegisterHandler(srvc, cch, metr, lg)

	// running server
	run(func() {
		l, err := net.Listen("tcp", cfg.Addr)
		if err != nil {
			panic("failed to listen tcp: " + err.Error())
		}
		defer func() {
			_ = l.Close()
		}()

		lg.Info("starting server", log.WithData("port", cfg.Addr))

		if err = srvr.Serve(l); err != nil {
			if err != nil {
				panic("failed to serve: " + err.Error())
			}
		}
	})

	lg.Info("shutdown server", nil)
	srvr.GracefulStop()
}

func run(fn func()) {
	go fn()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
}
