package middleware

import (
	"context"
	"github.com/alserov/restate/gateway/internal/utils"
	"github.com/labstack/echo/v4"
	"time"
)

func WithRateLimiter(lim int64) echo.MiddlewareFunc {
	limiter := utils.NewLimiter(context.Background(), lim, time.Millisecond*10)
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !limiter.Allow(c.Request().Context()) {
				return utils.NewError("request limiter", utils.TooManyRequests)
			}

			return next(c)
		}
	}
}
