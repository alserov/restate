package middleware

import (
	"github.com/alserov/restate/gateway/internal/wrappers"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func WithIdempotencyKey(fn echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(string(wrappers.ContextIdempotencyKey), uuid.NewString())
		return fn(c)
	}
}
