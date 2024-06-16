package app

import (
	"context"
	"github.com/alserov/restate/meetings/internal/config"
	"github.com/alserov/restate/meetings/internal/db/posgtres"
	"github.com/alserov/restate/meetings/internal/log"
	"github.com/alserov/restate/meetings/internal/metrics"
	"github.com/alserov/restate/meetings/internal/server/grpc"
	"github.com/alserov/restate/meetings/internal/service"
	"net"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	lg := log.NewLogger(cfg.Env, log.KindZap)

	db, closeConn := posgtres.MustConnect(cfg.DB.Dsn())
	defer closeConn()

	metr := metrics.NewMetrics(cfg.Broker.Addr)
	repo := posgtres.NewRepository(db)
	srvc := service.NewService(repo)
	srvr := grpc.RegisterHandler(srvc, metr, lg)

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
