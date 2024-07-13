package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
	"time"
)

func WithRateLimiter(lim uint) echo.MiddlewareFunc {
	vals := make(map[string]uint)
	mu := sync.RWMutex{}

	go func() {
		for range time.Tick(time.Minute) {
			mu.Lock()
			clear(vals)
			mu.Unlock()
		}
	}()

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			mu.Lock()
			val, ok := vals[c.Request().RemoteAddr]
			mu.Unlock()

			if ok && val > lim {
				c.JSON(http.StatusTooManyRequests, nil)
				return nil
			} else {
				mu.Lock()
				vals[c.Request().RemoteAddr]++
				mu.Unlock()
			}

			return next(c)
		}
	}
}
