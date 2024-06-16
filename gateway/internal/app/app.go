package app

import (
	"context"
	"errors"
	"fmt"
	estate "github.com/alserov/restate/estate/pkg/grpc"
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
)

func MustStart(cfg *config.Config) {
	lg := log.NewLogger(cfg.Env, log.KindZap)

	// app
	app := echo.New()
	lg.Info("initialized server", nil)

	// metrics
	metr := metrics.NewMetrics()
	lg.Info("initialized metrics", nil)

	// services
	estateGRPCClient := services.Dial[estate.EstateServiceClient]("", services.GRPCClient, grpcDial.NewEstateClient)
	meetingsGRPCClient := services.Dial[meetings.MeetingsServiceClient]("", services.GRPCClient, grpcDial.NewMeetingsClient)
	lg.Info("dialed services", nil)

	// routes
	ctrl := controller.NewController(app, lg, metr, &controller.Clients{
		Estate:   estateGRPCClient,
		Meetings: meetingsGRPCClient,
	})
	ctrl.SetupRoutes()
	lg.Info("set routes", nil)

	// server start
	run(func() {
		if err := app.Start(fmt.Sprintf("%s", cfg.Addr)); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic("failed to start server: " + err.Error())
		}
	})

	lg.Info("shutdown server", nil)
	app.Close()
}

func run(fn func()) {
	go fn()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	<-ctx.Done()
}
