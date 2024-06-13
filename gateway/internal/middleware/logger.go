package middleware

import (
	"github.com/alserov/restate/gateway/internal/log"
	"github.com/labstack/echo/v4"
)

func WithLogger(lg log.Logger) func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(string(log.ContextLogger), lg)
			return handlerFunc(c)
		}
	}
}
