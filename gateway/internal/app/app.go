package app

import (
	"context"
	"errors"
	estate "github.com/alserov/restate/estate/pkg/grpc"
	"github.com/alserov/restate/gateway/internal/async"
	"github.com/alserov/restate/gateway/internal/config"
	"github.com/alserov/restate/gateway/internal/controller"
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/alserov/restate/gateway/internal/metrics"
	"github.com/alserov/restate/gateway/internal/services"
	grpcDial "github.com/alserov/restate/gateway/internal/services/grpc"
	meetings "github.com/alserov/restate/meetings/pkg/grpc"
	"github.com/labstack/echo/v4"
	"net/http"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

func MustStart(cfg *config.Config) {
	lg := log.NewLogger(cfg.Env, log.KindZap)

	// app
	app := echo.New()
	lg.Info("initialized server", nil)

	// metrics
	metr := metrics.NewMetrics(async.NewProducer(async.Kafka, cfg.Broker.Addr, cfg.Broker.Topics.Metrics))
	lg.Info("initialized metrics", nil)

	// services
	estateGRPCClient := services.Dial[estate.EstateServiceClient](cfg.Services.Estate, services.GRPCClient, grpcDial.NewEstateClient)
	meetingsGRPCClient := services.Dial[meetings.MeetingsServiceClient](cfg.Services.Meetings, services.GRPCClient, grpcDial.NewMeetingsClient)
	lg.Info("dialed services", nil)

	// routes
	ctrl := controller.NewController(app, metr, lg, &controller.Clients{
		Estate:   estateGRPCClient,
		Meetings: meetingsGRPCClient,
	})
	ctrl.SetupRoutes()
	lg.Info("set routes", nil)

	// server start
	run(func() {
		if err := app.Start(cfg.Addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("failed to start server: " + err.Error())
		}
	})

	lg.Info("shutdown server", nil)
	if err := app.Shutdown(context.Background()); err != nil {
		lg.Error("failed to shutdown server", log.WithData("error", err.Error()))
	}
}

func run(fn func()) {
	go fn()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
}
