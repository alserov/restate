package middleware

import (
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/alserov/restate/gateway/internal/metrics"
	"github.com/alserov/restate/gateway/internal/middleware/wrappers"
	"github.com/labstack/echo/v4"
	"time"
)

func WithRequestObserver(metr metrics.Metrics) func(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(fn echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			defer func() {
				dur := time.Since(start)

				if err := metr.ObserveRequest(c.Request().Context(), c.Response().Status, dur, ""); err != nil {
					wrappers.ExtractLogger(c.Request().Context()).Warn("failed to observe request", log.WithData("warn", err.Error()))
				}
			}()

			err := fn(c)
			return err
		}
	}
}
