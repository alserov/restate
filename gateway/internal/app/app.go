package app

import (
	"context"
	"fmt"
	"github.com/alserov/restate/gateway/internal/config"
	"github.com/alserov/restate/gateway/internal/controller"
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/labstack/echo/v4"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {
	lg := log.NewLogger(cfg.Env, log.KindZap)

	// TODO: implement dial to services

	app := echo.New()

	ctrl := controller.NewController(app, lg)
	ctrl.SetupRoutes()

	run(func() {
		if err := app.Start(fmt.Sprintf(":%s", cfg.Addr)); err != nil {
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
