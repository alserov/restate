package middleware

import (
	"github.com/alserov/restate/gateway/internal/metrics"
	"github.com/labstack/echo/v4"
	"time"
)

func WithRequestObserver(metr metrics.Metrics) func(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(fn echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			defer func() {
				dur := time.Since(start)

				metr.ObserveRequest(c.Request().Context(), c.Response().Status, dur, "")
			}()

			err := fn(c)
			return err
		}
	}
}
